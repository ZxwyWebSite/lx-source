package kw

import (
	"errors"
	"io"
	"lx-source/src/env"
	"lx-source/src/sources"
	"lx-source/src/sources/custom/utils"
	"net/http"
	"strconv"
	"strings"
	"sync"

	"github.com/ZxwyWebSite/ztool"
	"github.com/ZxwyWebSite/ztool/x/bytesconv"
	"github.com/ZxwyWebSite/ztool/zcypt"
)

var (
	kw_pool *sync.Pool

	Url func(string, string) (string, string)

	parsemod bool
	convtype string
	desource string
	// desParse func([]byte, *playInfo) error
	// desParse func(any) ztool.Net_ResHandlerFunc
)

func init() {
	env.Inits.Add(func() {
		loger := env.Loger.NewGroup(`KwInit`)
		switch env.Config.Custom.Kw_Mode {
		case `0`, `bdapi`:
			loger.Debug(`use bdapi`)
			if ztool.Chk_IsNilStr(
				env.Config.Custom.Kw_Bd_Uid,
				env.Config.Custom.Kw_Bd_Token,
				env.Config.Custom.Kw_Bd_DevId,
			) {
				loger.Fatal(`使用bdapi且验证参数为空`)
			}
			// bdheader[`token`] = env.Config.Custom.Kw_Bd_Token
			bdheader[`uid`] = env.Config.Custom.Kw_Bd_Uid
			bdheader[`devId`] = env.Config.Custom.Kw_Bd_DevId
			kw_pool = &sync.Pool{New: func() any { return new(kwApi_Song) }}
			Url = bdapi
		case `1`, `kwdes`:
			Url = kwdes
			switch env.Config.Custom.Kw_Des_Type {
			case `0`, `text`:
				loger.Debug(`use kwdes text`)
				convtype = `convert_url2`
				// desParse = txtParse
			case `1`, `json`:
				loger.Debug(`use kwdes json`)
				convtype = `convert_url_with_sign`
				// desParse = ztool.Net_ResToStruct
				parsemod = true
			case `2`, `anti`:
				loger.Debug(`use kwdes anti`)
				Url = manti
			default:
				loger.Fatal(`未定义的返回格式，请检查配置 [Custom].Kw_Des_Type`)
			}
			desheader[`User-Agent`] = env.Config.Custom.Kw_Des_Header
			kw_pool = &sync.Pool{New: func() any { return new(playInfo) }}
			if env.Config.Custom.Kw_Des_Source != `` {
				desource = env.Config.Custom.Kw_Des_Source
			} else {
				dec, _ := zcypt.HexDecode([]byte{0x36, 0x62, 0x37, 0x37, 0x37, 0x30, 0x36, 0x63, 0x36, 0x31, 0x37, 0x39, 0x36, 0x35, 0x37, 0x32, 0x36, 0x38, 0x36, 0x34, 0x35, 0x66, 0x36, 0x31, 0x37, 0x32, 0x35, 0x66, 0x33, 0x35, 0x32, 0x65, 0x33, 0x31, 0x32, 0x65, 0x33, 0x30, 0x32, 0x65, 0x33, 0x30, 0x35, 0x66, 0x34, 0x32, 0x35, 0x66, 0x36, 0x61, 0x36, 0x39, 0x36, 0x31, 0x36, 0x62, 0x36, 0x66, 0x36, 0x65, 0x36, 0x37, 0x35, 0x66, 0x37, 0x36, 0x36, 0x38, 0x32, 0x65, 0x36, 0x31, 0x37, 0x30, 0x36, 0x62})
				desource = bytesconv.BytesToString(dec)
			}
		default:
			loger.Fatal(`未定义的接口模式，请检查配置 [Custom].Kw_Mode`)
		}
		loger.Free()
	})
}

func bdapi(songMid, quality string) (ourl, msg string) {
	loger := env.Loger.NewGroup(`Kw`)
	defer loger.Free()
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
		`&uid=`, env.Config.Custom.Kw_Bd_Uid,
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
		msg = ztool.Str_FastConcat(strconv.Itoa(resp.Code), `: `, resp.Msg)
		return
	}
	ourl = utils.DelQuery(resp.Data.URL) //strings.Split(resp.Data.URL, `?`)[0]
	return
}

func kwdes(songMid, quality string) (ourl, msg string) {
	loger := env.Loger.NewGroup(`Kw`)
	defer loger.Free()
	infoFile, ok := fileInfo[quality]
	if !ok {
		msg = sources.E_QNotSupport
		return
	}
	target_url := ztool.Str_FastConcat(
		`https://mobi.kuwo.cn/mobi.s?f=kuwo&q=`,
		base64_encrypt(ztool.Str_FastConcat(
			`corp=kuwo&p2p=1&sig=0&notrace=0&priority=bitrate&network=WIFI&mode=down`,
			`&source=`, desource,
			`&type=`, convtype,
			`&br=`, infoFile.H, infoFile.E,
			`&format=`, infoFile.E,
			`&rid=`, songMid,
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
		resp.Data.URL = utils.DelQuery(resp.Data.URL)
		loger.Debug(`Resp: %+v`, resp)
		if resp.Code != http.StatusOK {
			msg = ztool.Str_FastConcat(`failed: `, resp.Msg)
			loger.Debug(msg)
			return
		}
		realQuality := strconv.Itoa(resp.Data.Bitrate)
		if realQuality != infoFile.H[:len(infoFile.H)-1] /*&& resp.Data.Bitrate != 1*/ {
			msg = sources.E_QNotMatch
			if !env.Config.Source.ForceFallback {
				return
			}
		}
		ourl = resp.Data.URL //resp.Data.URL[:strings.Index(resp.Data.URL, `?`)]
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
				realQuality := infoData[`bitrate`]
				if realQuality != infoFile.H[:len(infoFile.H)-1] {
					msg = sources.E_QNotMatch
					if !env.Config.Source.ForceFallback {
						return
					}
				}
				ourl = utils.DelQuery(infoData[`url`]) //infoData[`url`][:strings.Index(infoData[`url`], `?`)]
				return
			},
		},
	)
	return
}

// 一种替代方案，仅在试听可用时获取cdn前缀，缺点是无法获取不能试听的歌曲
func manti(songMid, quality string) (ourl, msg string) {
	loger := env.Loger.NewGroup(`Kw`)
	defer loger.Free()
	infoFile, ok := fileInfo[quality]
	if !ok {
		msg = sources.E_QNotSupport
		return
	}
	// 获取CDN地址
	// var out_a string
	/*var out_c struct {
		// Timestamp int `json:"timestamp"`
		Songs []struct {
			ID                int    `json:"id"`
			Duration          int    `json:"duration"`
			URL               string `json:"url"`
			HTTPS             string `json:"https"`
			CarURL            string `json:"car_url"`
			CarURLHTTPS       string `json:"car_url_https"`
			Format            string `json:"format"`
			Br                int    `json:"br"`
			OverseasCopyright string `json:"overseas_copyright"`
			Start             int    `json:"start"`
			End               int    `json:"end"`
			Group             string `json:"group"`
		} `json:"songs"`
		IP         string `json:"ip"`
		Country    string `json:"country"`
		Region     string `json:"region"`
		Locationid int    `json:"locationid"`
		Code       int    `json:"code"`
		Result     string `json:"result"`
	}
	err := ztool.Net_Request(
		http.MethodGet,
		`https://musicpay30.kuwo.cn/audi.tion?op=query&ids=`+songMid,
		nil,
		[]ztool.Net_ReqHandlerFunc{ztool.Net_ReqAddHeaders()},
		[]ztool.Net_ResHandlerFunc{ztool.Net_ResToStruct(&out_c)},
	)
	if err != nil {
		loger.Error(`Request: %s`, err)
		msg = sources.ErrHttpReq
		return
	}*/
	// 获取音频路径
	if quality != sources.Q_128k {
		ourl, msg = manti(songMid, sources.Q_128k)
		if msg != `` {
			return
		}
		if i := strings.LastIndexByte(ourl, '/'); i != -1 {
			if ourl[i+1:] == `2272659253.mp3` {
				msg = sources.ErrNoLink
				return
			}
		}
		target_url := ztool.Str_FastConcat(
			`https://mobi.kuwo.cn/mobi.s?f=web&type=convert_url`,
			`&br=`, infoFile.H, infoFile.E,
			`&format=`, infoFile.E,
			`&rid=`, songMid,
		)
		var out_u, realQuality string
		err := ztool.Net_Request(
			http.MethodGet, target_url, nil,
			[]ztool.Net_ReqHandlerFunc{ztool.Net_ReqAddHeaders()},
			[]ztool.Net_ResHandlerFunc{func(res *http.Response) (err error) {
				var data []byte
				data, err = io.ReadAll(res.Body)
				if err != nil {
					return
				}
				if res.StatusCode != http.StatusOK {
					return errors.New(`failed: ` + res.Status)
				}
				infoData := mkMap(data)
				loger.Debug(`uData: %+v`, infoData)
				realQuality = infoData[`bitrate`]
				out_u = utils.DelQuery(infoData[`url`])
				return
			}},
		)
		if err != nil {
			loger.Error(`Request: %s`, err)
			msg = sources.ErrHttpReq
			return
		}
		if realQuality != infoFile.H[:len(infoFile.H)-1] {
			msg = sources.E_QNotMatch
			if !env.Config.Source.ForceFallback {
				return
			}
		}
		if ourl != `` && out_u != `` {
			ourl = ourl[:24] + out_u
		}
		return
	}
	resp := kw_pool.Get().(*playInfo)
	defer kw_pool.Put(resp)

	err := ztool.Net_Request(
		http.MethodGet, `https://mobi.kuwo.cn/mobi.s?f=web&type=convert_url_with_sign&br=128kmp3&format=mp3&rid=`+songMid, nil,
		[]ztool.Net_ReqHandlerFunc{ztool.Net_ReqAddHeader(desheader)},
		[]ztool.Net_ResHandlerFunc{ztool.Net_ResToStruct(&resp)},
	)
	if err != nil {
		loger.Error(`Request: %s`, err)
		msg = sources.ErrHttpReq
		return
	}
	loger.Debug(`tData: %+v`, resp)
	if resp.Code != http.StatusOK {
		msg = ztool.Str_FastConcat(`failed: `, resp.Msg)
		loger.Debug(msg)
		return
	}
	realQuality := strconv.Itoa(resp.Data.Bitrate)
	if realQuality != infoFile.H[:len(infoFile.H)-1] && resp.Data.Bitrate != 1 {
		msg = sources.E_QNotMatch
		if !env.Config.Source.ForceFallback {
			return
		}
	}
	ourl = utils.DelQuery(resp.Data.URL)
	return
}
