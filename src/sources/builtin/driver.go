// 内置解析源
package builtin

import (
	"lx-source/src/caches"
	"lx-source/src/env"
	"lx-source/src/sources"
	"lx-source/src/sources/custom/kw"
	"lx-source/src/sources/custom/tx"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/ZxwyWebSite/ztool"
)

type Source struct{}

// 预检 (兼容旧接口)
func (s *Source) Verify(c *caches.Query) (rquery string, ok bool) {
	rquery, ok = qualitys[c.Quality][c.Source]
	return
}

var (
	// 并发对象池 (用户限制在Router处实现)
	wy_pool = &sync.Pool{New: func() any { return new(WyApi_Song) }}
	mg_pool = &sync.Pool{New: func() any { return new(MgApi_Song) }}
	kw_pool = &sync.Pool{New: func() any { return new(KwApi_Song) }}
	kg_pool = &sync.Pool{New: func() any { return new(KgApi_Song) }}
	// tx_pool = &sync.Pool{New: func() any { return new(res_tx) }}
)

const (
	errHttpReq = `无法连接解析接口`
	errNoLink  = `无法获取试听链接`
	errDisable = `该音乐源已被禁用`
)

// 查询
func (s *Source) GetLink(c *caches.Query) (outlink string, msg string) {
	rquery, ok := s.Verify(c)
	if !ok /*&& c.Source != `tx`*/ {
		msg = sources.Err_Verify //`Verify Failed`
		return
	}
	// var outlink string
	jx := env.Loger.NewGroup(`Sources`) //sources.Loger.AppGroup(`builtin`) //env.Loger.NewGroup(`JieXiApis`)
	switch c.Source {
	case s_wy:
		if !env.Config.Custom.Wy_Enable {
			msg = errDisable
			return
		}
		resp := wy_pool.Get().(*WyApi_Song)
		defer wy_pool.Put(resp)

		// url := ztool.Str_FastConcat(`http://`, api_wy, `?id=`, c.MusicID, `&level=`, rquery, `&noCookie=true`)
		url := ztool.Str_FastConcat(`https://`, api_wy, `&id=`, c.MusicID, `&level=`, rquery, `&encodeType=`, c.Extname)
		// jx.Debug(`Wy, Url: %v`, url)
		// wy源增加后端重试 默认3次
		for i := 0; true; i++ {
			// _, err := ztool.Net_HttpReq(http.MethodGet, url, nil, header_wy, &resp)
			_, err := ztool.Net_HttpReq(http.MethodGet, url, nil, nil, &resp)
			if err != nil {
				jx.Error(`HttpReq, Err: %s, ReTry: %v`, err, i)
				if i > 3 {
					jx.Error(`Wy, HttpReq: %s`, err)
					msg = errHttpReq //err.Error()
					return
				}
				time.Sleep(time.Second)
				continue
			}
			break
		}
		jx.Debug(`Wy, Resp: %+v`, resp)
		if len(resp.Data) == 0 {
			msg = `No Data：Api接口忙，请稍后重试`
			return
		}
		var data = resp.Data[0]
		if data.Code != 200 || data.FreeTrialInfo != nil {
			// jx.Error("发生错误, 返回数据:\n%#v", resp)
			msg = `触发风控或专辑单独收费`
			return
		}
		if data.Level != rquery {
			msg = ztool.Str_FastConcat(`实际音质不匹配: `, rquery, ` <= `, data.Level) // 实际音质不匹配: exhigh <= standard
			return
		}
		// jx.Info(`WyLink, RealQuality: %v`, data.Level)
		outlink = data.URL
	case s_mg:
		resp := mg_pool.Get().(*MgApi_Song)
		defer mg_pool.Put(resp)

		url := ztool.Str_FastConcat(`https://`, api_mg, `?copyrightId=`, c.MusicID, `&type=`, rquery)
		// jx.Debug(`Mg, Url: %v`, url)
		_, err := ztool.Net_HttpReq(http.MethodGet, url, nil, header_mg, &resp)
		if err != nil {
			jx.Error(`Mg, HttpReq: %s`, err)
			msg = errHttpReq //err.Error()
			return
		}
		jx.Debug(`Mg, Resp: %+v`, resp)
		if link := resp.Data.PlayURL; link != `` {
			outlink = `https:` + link
		} // else {
		// 	jx.Debug(`Mg, Err: %#v`, resp)
		// }
	case s_kw:
		if !env.Config.Custom.Kw_Enable {
			msg = errDisable
			return
		}
		ourl, emsg := kw.Url(c.MusicID, c.Quality)
		if emsg != `` {
			msg = emsg
			return
		}
		outlink = ourl
	// case s_kw:
	// 	resp := kw_pool.Get().(*KwApi_Song)
	// 	defer kw_pool.Put(resp)

	// 	url := ztool.Str_FastConcat(`https://`, api_kw, `/`, c.MusicID, `?isMv=0&format=`, c.Extname, `&br=`, rquery, c.Extname, `&level=`)
	// 	// jx.Debug(`Kw, Url: %s`, url)
	// 	_, err := ztool.Net_HttpReq(http.MethodGet, url, nil, header_kw, &resp)
	// 	if err != nil {
	// 		jx.Error(`Kw, HttpReq: %s`, err)
	// 		msg = errHttpReq //err.Error()
	// 		return
	// 	}
	// 	jx.Debug(`Kw, Resp: %+v`, resp)
	// 	if resp.Code != 200 || resp.Data.AudioInfo.Bitrate == `1` {
	// 		// jx.Debug(`Kw, Err: %#v`, resp)
	// 		msg = ztool.Str_FastConcat(`failed: `, resp.Msg)
	// 		return
	// 	}
	// 	outlink = strings.Split(resp.Data.URL, `?`)[0]
	case s_kg:
		resp := kg_pool.Get().(*KgApi_Song)
		defer kg_pool.Put(resp)

		// sep := strings.Split(c.MusicID, `-`) // 分割 Hash-Album 如 6DC276334F56E22BE2A0E8254D332B45-13097991
		// alb := func() string {
		// 	if len(sep) >= 2 {
		// 		return sep[1]
		// 	}
		// 	return ``
		// }()
		sep := c.Split()
		url := ztool.Str_FastConcat(api_kg, `&hash=`, sep[0], `&album_id=`, sep[1], `&_=`, strconv.FormatInt(time.Now().UnixMilli(), 10))
		// jx.Debug(`Kg, Url: %s`, url)
		_, err := ztool.Net_HttpReq(http.MethodGet, url, nil, nil, &resp)
		if err != nil {
			jx.Error(`Kg, HttpReq: %s`, err)
			msg = errHttpReq //err.Error()
			return
		}
		jx.Debug(`Kw, Resp: %+v`, resp)
		if resp.ErrCode != 0 {
			msg = ztool.Str_FastConcat(`Error: `, strconv.Itoa(resp.ErrCode))
			return
		}
		var data KgApi_Data
		err = ztool.Val_MapToStruct(resp.Data, &data)
		if err != nil {
			msg = err.Error()
			return
		}
		if data.PlayBackupURL == `` {
			if data.PlayURL == `` {
				msg = errNoLink
				return
			}
			outlink = data.PlayURL
		}
		outlink = data.PlayBackupURL
	case s_tx:
		sep := c.Split()
		ourl, emsg := tx.Url(sep[0], c.Quality)
		if emsg != `` {
			msg = emsg
			return
		}
		outlink = ourl
	// case `otx`:
	// 	resp := tx_pool.Get().(*res_tx)
	// 	defer tx_pool.Put(resp)

	// 	sep := c.Split()
	// 	url := ztool.Str_FastConcat(api_tx,
	// 		`{"comm":{"ct":24,"cv":0,"format":"json","uin":"10086"},"req":{"method":"GetCdnDispatch","module":"CDN.SrfCdnDispatchServer","param":{"calltype":0,"guid":"1535153710","userip":""}},"req_0":{"method":"CgiGetVkey","module":"vkey.GetVkeyServer","param":{`,
	// 		func(s string) string {
	// 			if s == `` {
	// 				return ``
	// 			}
	// 			return ztool.Str_FastConcat(`"filename":["`, rquery, s, `.`, c.Extname, `"],`)
	// 		}(sep[1]),
	// 		`"guid":"1535153710","loginflag":1,"platform":"20","songmid":["`, sep[0], `"],"songtype":[0],"uin":"10086"}}}`,
	// 	)
	// 	// jx.Debug(`Tx, Url: %s`, url)
	// 	out, err := ztool.Net_HttpReq(http.MethodGet, url, nil, header_tx, &resp)
	// 	if err != nil {
	// 		jx.Error(`Tx, HttpReq: %s`, err)
	// 		msg = errHttpReq
	// 		return
	// 	}
	// 	jx.Debug(`Tx, Resp: %s`, out)
	// 	if resp.Code != 0 {
	// 		msg = ztool.Str_FastConcat(`Error: `, strconv.Itoa(resp.Code))
	// 		return
	// 	}
	// 	if resp.Req0.Data.Midurlinfo[0].Purl == `` {
	// 		msg = errNoLink
	// 		return
	// 	}
	// 	outlink = ztool.Str_FastConcat(`https://dl.stream.qqmusic.qq.com/`, resp.Req0.Data.Midurlinfo[0].Purl)
	default:
		msg = `不支持的平台`
		return
	}
	return
}
