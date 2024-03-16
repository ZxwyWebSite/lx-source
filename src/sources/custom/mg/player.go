package mg

import (
	"lx-source/src/env"
	"lx-source/src/sources"
	"lx-source/src/sources/custom/utils"
	"lx-source/src/sources/example"
	"net/http"
	"sync"

	"github.com/ZxwyWebSite/ztool"
)

var (
	Url func(string, string) (string, string)

	mg_pool *sync.Pool
)

func init() {
	env.Inits.Add(func() {
		loger := env.Loger.NewGroup(`MgInit`)
		switch env.Config.Custom.Mg_Mode {
		case `0`, `builtin`:
			loger.Debug(`use builtin`)
			mg_pool = &sync.Pool{New: func() any { return new(mgApi_Song) }}
			Url = builtin
		case `1`, `custom`:
			loger.Debug(`use custom`)
			if ztool.Chk_IsNilStr(
				// env.Config.Custom.Mg_Usr_VerId,
				// env.Config.Custom.Mg_Usr_Token,
				env.Config.Custom.Mg_Usr_OSVer,
				env.Config.Custom.Mg_Usr_ReqUA,
			) {
				loger.Fatal(`使用自定义账号且用户参数为空`)
			}
			mg_pool = &sync.Pool{New: func() any { return new(playInfo) }}
			Url = mcustom
		case `2`, `malbum`:
			loger.Debug(`use malbum`)
			Url = malbum
		default:
			loger.Fatal(`未定义的接口模式，请检查配置 [Custom].Mg_Mode`)
		}
		loger.Free()
	})
}

func builtin(songMid, quality string) (ourl, msg string) {
	loger := env.Loger.NewGroup(`Mg`)
	defer loger.Free()
	rquality, ok := qualitys[quality]
	if !ok {
		msg = sources.E_QNotSupport
		return
	}
	resp := mg_pool.Get().(*mgApi_Song)
	defer mg_pool.Put(resp)
	url := ztool.Str_FastConcat(`https://`, example.Api_mg, `?copyrightId=`, songMid, `&type=`, rquality)
	err := ztool.Net_Request(
		http.MethodGet, url, nil,
		[]ztool.Net_ReqHandlerFunc{ztool.Net_ReqAddHeaders(example.Header_mg)},
		[]ztool.Net_ResHandlerFunc{ztool.Net_ResToStruct(&resp)},
	)
	if err != nil {
		loger.Error(`HttpReq: %s`, err)
		msg = sources.ErrHttpReq
		return
	}
	loger.Debug(`Resp: %+v`, resp)
	// ourl, msg = resp.Url(rquality)
	if resp.Data.PlayURL != `` {
		ourl = `https:` + utils.DelQuery(resp.Data.PlayURL)
	} else {
		msg = ztool.Str_FastConcat(resp.Code, `: `, resp.Msg)
	}
	return
}

func mcustom(songMid, quality string) (ourl, msg string) {
	loger := env.Loger.NewGroup(`Mg`)
	defer loger.Free()
	rquality, ok := qualityMap[quality]
	if !ok {
		msg = sources.E_QNotSupport
		return
	}
	url := ztool.Str_FastConcat(
		`https://app.c.nf.migu.cn/MIGUM2.0/strategy/listen-url/v2.4?toneFlag=`, rquality,
		`&songId=`, songMid,
		`&resourceType=2`,
	)
	resp := mg_pool.Get().(*playInfo)
	defer mg_pool.Put(resp)
	err := ztool.Net_Request(
		http.MethodGet, url, nil,
		[]ztool.Net_ReqHandlerFunc{ztool.Net_ReqAddHeaders(map[string]string{
			`User-Agent`: env.Config.Custom.Mg_Usr_ReqUA,
			`aversionid`: env.Config.Custom.Mg_Usr_VerId,
			`token`:      env.Config.Custom.Mg_Usr_Token,
			`channel`:    `0146832`,
			`language`:   `Chinese`,
			`ua`:         `Android_migu`,
			`mode`:       `android`,
			`os`:         `Android ` + env.Config.Custom.Mg_Usr_OSVer,
		})},
		[]ztool.Net_ResHandlerFunc{ztool.Net_ResToStruct(&resp)},
	)
	if err != nil {
		loger.Error(`Request: %s`, err)
		msg = sources.ErrHttpReq
		return
	}
	loger.Debug(`Resp: %+v`, resp)
	if resp.Code != `000000` {
		msg = resp.Info
		return
	}
	if resp.Data.URL == `` {
		msg = `No Data: 无返回链接`
		return
	}
	if resp.Data.AudioFormatType != rquality {
		msg = ztool.Str_FastConcat(`实际音质不匹配: `, rquality, ` <= `, resp.Data.AudioFormatType)
		if !env.Config.Source.ForceFallback {
			return
		}
	}
	ourl = utils.DelQuery(resp.Data.URL)
	return
}

func malbum(songMid, quality string) (ourl, msg string) {
	loger := env.Loger.NewGroup(`Mg`)
	defer loger.Free()
	rquality, ok := qualityMap[quality]
	if !ok {
		msg = sources.E_QNotSupport
		return
	}
	minfo, err := getMusicInfo(songMid)
	if err != nil {
		msg = err.Error()
		return
	}
	loger.Debug(`mInfo: %+v`, minfo)
	var hasQuality bool
	for _, v := range minfo.Resource[0].NewRateFormats {
		if hasQuality = v.FormatType == rquality; hasQuality {
			break
		}
	}
	if !hasQuality {
		msg = sources.E_QNotMatch
		return
	}
	ainfo, err := getAlbumInfo(minfo.Resource[0].AlbumID)
	if err != nil {
		msg = err.Error()
		return
	}
	loger.Debug(`aInfo: %+v`, ainfo)
	for _, v := range ainfo.Resource[0].SongItems {
		if v.CopyrightID == songMid {
			for _, w := range v.NewRateFormats {
				if w.FormatType == rquality {
					if rquality == `PQ` || rquality == `HQ` {
						ourl = w.URL
					} else {
						ourl = w.AndroidURL
					}
					break
				}
			}
			break
		}
	}
	if ourl == `` {
		msg = sources.E_NoLink
	} else {
		ourl = `https://freetyst.nf.migu.cn` + ourl[24:]
	}
	return
}
