package kw

import (
	"io"
	"lx-source/src/env"
	"lx-source/src/sources"
	"net/http"
	"strconv"
	"strings"
	"sync"

	"github.com/ZxwyWebSite/ztool"
)

var (
	kw_pool *sync.Pool

	Url func(string, string) (string, string)

	parsemod bool
	convtype string
	// desParse func([]byte, *playInfo) error
	// desParse func(any) ztool.Net_ResHandlerFunc
)

func init() {
	env.Inits.Add(func() {
		loger := env.Loger.NewGroup(`KwInit`)
		switch env.Config.Custom.Kw_Mode {
		case `0`, `bdapi`:
			loger.Debug(`Use bdapi`)
			if ztool.Chk_IsNilStr(
				env.Config.Custom.Kw_Bd_Uid,
				env.Config.Custom.Kw_Bd_Token,
				env.Config.Custom.Kw_Bd_DevId,
			) {
				loger.Fatal(`使用bdapi且验证参数为空`)
			}
			bdheader[`uid`] = env.Config.Custom.Kw_Bd_Uid
			bdheader[`devId`] = env.Config.Custom.Kw_Bd_DevId
			kw_pool = &sync.Pool{New: func() any { return new(kwApi_Song) }}
			Url = bdapi
		case `1`, `kwdes`:
			switch env.Config.Custom.Kw_Des_Type {
			case `0`, `text`:
				loger.Debug(`Use kwdes_text`)
				convtype = `convert_url2`
				// desParse = txtParse
			case `1`, `json`:
				loger.Debug(`Use kwdes_json`)
				convtype = `convert_url_with_sign`
				// desParse = ztool.Net_ResToStruct
				parsemod = true
			default:
				loger.Fatal(`未定义的返回格式，请检查配置 [Custom].Kw_Des_Type`)
			}
			desheader[`User-Agent`] = env.Config.Custom.Kw_Des_Header
			kw_pool = &sync.Pool{New: func() any { return new(playInfo) }}
			Url = kwdes
		default:
			loger.Fatal(`未定义的接口模式，请检查配置 [Custom].Kw_Mode`)
		}
		loger = nil
	})
}

func bdapi(songMid, quality string) (ourl, msg string) {
	loger := env.Loger.NewGroup(`Kw`)
	info, ok := fileInfo[quality]
	if !ok {
		msg = sources.E_QNotSupport
		return
	}
	resp := kw_pool.Get().(*kwApi_Song)
	defer kw_pool.Put(resp)

	url := ztool.Str_FastConcat(
		`https://bd-api.kuwo.cn/api/service/music/downloadInfo/`, songMid,
		`?isMv=0&format=`, info.E,
		`&br=`, info.H, info.E, //`&level=`,
		`&uin=`, env.Config.Custom.Kw_Bd_Uid,
		`&token=`, env.Config.Custom.Kw_Bd_Token,
	)
	// jx.Debug(`Kw, Url: %s`, url)
	_, err := ztool.Net_HttpReq(http.MethodGet, url, nil, bdheader, &resp)
	if err != nil {
		loger.Error(`HttpReq: %s`, err)
		msg = sources.ErrHttpReq
		return
	}
	loger.Debug(`Resp: %+v`, resp)
	if resp.Code != 200 || resp.Data.AudioInfo.Bitrate == `1` {
		// jx.Debug(`Kw, Err: %#v`, resp)
		msg = ztool.Str_FastConcat(`failed: `, resp.Msg)
		return
	}
	ourl = strings.Split(resp.Data.URL, `?`)[0]
	return
}

func kwdes(songMid, quality string) (ourl, msg string) {
	loger := env.Loger.NewGroup(`Kw`)
	infoFile, ok := fileInfo[quality]
	if !ok {
		msg = sources.E_QNotSupport
		return
	}
	target_url := ztool.Str_FastConcat(
		`https://mobi.kuwo.cn/mobi.s?f=kuwo&q=`,
		Base64_encrypt(ztool.Str_FastConcat(
			`user=0&android_id=0&prod=kwplayer_ar_8.5.5.0&corp=kuwo&newver=3&vipver=8.5.5.0&source=kwplayer_ar_8.5.5.0_apk_keluze.apk&p2p=1&notrace=0`,
			`&type=`, convtype,
			`&br=`, infoFile.H, infoFile.E,
			`&format=`, infoFile.E,
			`&rid=`, songMid,
			`&priority=bitrate&loginUid=0&network=WIFI&loginSid=0&mode=down`,
		)),
	)
	if parsemod {
		resp := kw_pool.Get().(*playInfo)
		defer kw_pool.Put(resp)

		err := ztool.Net_Request(http.MethodGet, target_url, nil,
			[]ztool.Net_ReqHandlerFunc{ztool.Net_ReqAddHeader(desheader)},
			[]ztool.Net_ResHandlerFunc{ztool.Net_ResToStruct(&resp)},
		)
		if err != nil {
			loger.Error(`Request: %s`, err)
			msg = sources.ErrHttpReq
			return
		}
		loger.Debug(`Resp: %+v`, resp)
		if resp.Code != http.StatusOK {
			msg = ztool.Str_FastConcat(`failed: `, resp.Msg)
			loger.Debug(msg)
			return
		}
		realQuality := strconv.Itoa(resp.Data.Bitrate)
		if qualityMapReverse[realQuality] != quality {
			msg = sources.E_QNotMatch
			return
		}
		ourl = resp.Data.URL[:strings.Index(resp.Data.URL, `?`)]
		return
	}
	ztool.Net_Request(http.MethodGet, target_url, nil,
		[]ztool.Net_ReqHandlerFunc{
			ztool.Net_ReqAddHeader(desheader),
		},
		[]ztool.Net_ResHandlerFunc{
			func(res *http.Response) (err error) {
				data, err := io.ReadAll(res.Body)
				if err != nil {
					msg = err.Error()
					return
				}
				if res.StatusCode != http.StatusOK {
					msg = ztool.Str_FastConcat(`failed: `, res.Status)
					loger.Debug(msg)
					return
				}
				infoData := mkMap(data)
				loger.Debug(`infoData: %+v`, infoData)
				realQuality := qualityMapReverse[infoData[`bitrate`]]
				if realQuality != quality {
					msg = sources.E_QNotMatch
					return
				}
				ourl = infoData[`url`][:strings.Index(infoData[`url`], `?`)]
				return
			},
		},
	)
	return
}
