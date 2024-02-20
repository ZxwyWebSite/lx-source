package wy

import (
	"lx-source/src/env"
	"lx-source/src/sources"
	"lx-source/src/sources/custom/utils"
	wm "lx-source/src/sources/custom/wy/modules"
	"lx-source/src/sources/example"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/ZxwyWebSite/ztool"
	"github.com/ZxwyWebSite/ztool/x/cookie"
)

var (
	wy_pool = &sync.Pool{New: func() any { return new(wm.PlayInfo) }}
	// wv_pool *sync.Pool

	Url func(string, string) (string, string)
)

func init() {
	env.Inits.Add(func() {
		loger := env.Loger.NewGroup(`WyInit`)
		switch env.Config.Custom.Wy_Mode {
		case `0`, `builtin`:
			loger.Debug(`use builtin`)
			// if env.Config.Source.MusicIdVerify {
			// 	wv_pool = &sync.Pool{New: func() any { return new(verifyInfo) }}
			// }
			Url = builtin
		case `1`, `163api`:
			if env.Config.Custom.Wy_Api_Cookie == `` {
				loger.Fatal(`使用163api且Cookie参数为空`)
			}
			switch env.Config.Custom.Wy_Api_Type {
			case `0`, `native`:
				loger.Debug(`use 163api module`)
				Url = nmModule
			case `1`, `remote`:
				loger.Debug(`use 163api custom`)
				if env.Config.Custom.Wy_Api_Address == `` {
					loger.Fatal(`自定义接口地址为空`)
				}
				if env.Config.Custom.Wy_Api_Address[len(env.Config.Custom.Wy_Api_Address)-1] != '/' {
					env.Config.Custom.Wy_Api_Address += "/" // 补全尾部斜杠
				}
				loger.Info(`使用自定义接口: %v`, env.Config.Custom.Wy_Api_Address)
				Url = nmCustom
			default:
				loger.Fatal(`未定义的调用方式，请检查配置 [Custom].Wy_Api_Type`)
			}
		default:
			loger.Fatal(`未定义的接口模式，请检查配置 [Custom].Wy_Mode`)
		}
		loger.Free()
	})
}

func builtin(songMid, quality string) (ourl, msg string) {
	loger := env.Loger.NewGroup(`Wy`)
	defer loger.Free()
	rquality, ok := qualityMap[quality]
	if !ok {
		msg = sources.E_QNotSupport
		return
	}
	resp := wy_pool.Get().(*wm.PlayInfo)
	defer wy_pool.Put(resp)
	url := ztool.Str_FastConcat(
		`https://`, example.Api_wy, `&id=`, songMid, `&level=`, rquality,
		`&timestamp=`, strconv.FormatInt(time.Now().UnixMilli(), 10),
	)
	err := ztool.Net_Request(
		http.MethodGet, url, nil,
		[]ztool.Net_ReqHandlerFunc{ztool.Net_ReqAddHeaders(example.Header_wy)},
		[]ztool.Net_ResHandlerFunc{ztool.Net_ResToStruct(&resp)},
	)
	if err != nil {
		loger.Error(`HttpReq: %s`, err)
		msg = sources.ErrHttpReq
		return
	}
	loger.Debug(`Resp: %+v`, resp)
	if len(resp.Data) == 0 {
		msg = `No Data：Api接口忙，请稍后重试`
		return
	}
	var data = resp.Data[0]
	if data.Code != 200 || data.FreeTrialInfo != nil {
		msg = `触发风控或专辑单独收费: ` + strconv.Itoa(data.Code)
		return
	}
	if data.Level != rquality {
		msg = ztool.Str_FastConcat(`实际音质不匹配: `, rquality, ` <= `, data.Level)
		if !env.Config.Source.ForceFallback {
			return
		}
	}
	ourl = data.URL
	return
}

func nmModule(songMid, quality string) (ourl, msg string) {
	loger := env.Loger.NewGroup(`Wy`)
	defer loger.Free()
	rquality, ok := qualityMap[quality]
	if !ok {
		msg = sources.E_QNotSupport
		return
	}
	cookies := cookie.Parse(env.Config.Custom.Wy_Api_Cookie)
	answer, err := wm.SongUrlV1(wm.ReqQuery{
		Cookie: cookie.ToMap(cookies),
		Ids:    songMid,
		// Br:     rquality,
		Level: rquality,
	})
	body := wy_pool.Get().(*wm.PlayInfo)
	defer wy_pool.Put(body)
	if err == nil {
		err = ztool.Val_MapToStruct(answer.Body, &body)
	}
	if err != nil {
		loger.Error(`SongUrl: %s`, err)
		msg = sources.ErrHttpReq
		return
	}
	loger.Debug(`Resp: %+v`, body)
	if len(body.Data) == 0 {
		msg = `No Data：无返回数据`
		return
	}
	data := body.Data[0]
	if data.Code != 200 {
		msg = `触发风控或专辑单独收费: ` + strconv.Itoa(data.Code)
		return
	}
	if data.Level != rquality {
		msg = ztool.Str_FastConcat(`实际音质不匹配: `, rquality, ` <= `, data.Level)
		if !env.Config.Source.ForceFallback {
			return
		}
	}
	// br := strconv.Itoa(data.Br) // 注：由于flac返回br值不固定，暂无法进行比较
	// if br != rquality && !ztool.Chk_IsMatch(br, sources.Q_flac, sources.Q_fl24) {
	// 	msg = sources.E_QNotMatch
	// 	return
	// }
	ourl = utils.DelQuery(data.URL)
	return
}

func nmCustom(songMid, quality string) (ourl, msg string) {
	loger := env.Loger.NewGroup(`Wy`)
	defer loger.Free()
	rquality, ok := qualityMap[quality]
	if !ok {
		msg = sources.E_QNotSupport
		return
	}
	body := wy_pool.Get().(*wm.PlayInfo)
	defer wy_pool.Put(body)
	err := ztool.Net_Request(
		http.MethodGet,
		ztool.Str_FastConcat(
			env.Config.Custom.Wy_Api_Address, `song/url/v1`, `?id=`, songMid, `&level=`, rquality,
			// `&timestamp=`, strconv.FormatInt(time.Now().UnixMilli(), 10),
		), nil,
		[]ztool.Net_ReqHandlerFunc{ztool.Net_ReqAddHeaders(map[string]string{
			`Cookie`: env.Config.Custom.Wy_Api_Cookie,
		})},
		[]ztool.Net_ResHandlerFunc{ztool.Net_ResToStruct(&body)},
	)
	if err != nil {
		loger.Error(`SongUrl: %s`, err)
		msg = sources.ErrHttpReq
		return
	}
	loger.Debug(`Resp: %+v`, body)
	if len(body.Data) == 0 {
		msg = `No Data：无返回数据`
		return
	}
	data := body.Data[0]
	if data.Code != 200 {
		msg = `触发风控或专辑单独收费: ` + strconv.Itoa(data.Code)
		return
	}
	if data.Level != rquality {
		msg = ztool.Str_FastConcat(`实际音质不匹配: `, rquality, ` <= `, data.Level)
		if !env.Config.Source.ForceFallback {
			return
		}
	}
	ourl = utils.DelQuery(data.URL)
	return
}
