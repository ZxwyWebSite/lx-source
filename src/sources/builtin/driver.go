// 内置解析源
package builtin

import (
	"lx-source/src/caches"
	"lx-source/src/env"
	"lx-source/src/sources"
	"net/http"
	"strings"
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
	wy_pool = &sync.Pool{New: func() any { return new(FyApi_Song) }}
	mg_pool = &sync.Pool{New: func() any { return new(MgApi_Song) }}
	kw_pool = &sync.Pool{New: func() any { return new(KwApi_Song) }}
)

// 查询
func (s *Source) GetLink(c *caches.Query) (outlink string, msg string) {
	rquery, ok := s.Verify(c)
	if !ok {
		msg = sources.Err_Verify //`Verify Failed`
		return
	}
	// var outlink string
	jx := env.Loger.NewGroup(`Sources`) //sources.Loger.AppGroup(`builtin`) //env.Loger.NewGroup(`JieXiApis`)
	switch c.Source {
	case s_wy:
		resp := wy_pool.Get().(*FyApi_Song)
		defer wy_pool.Put(resp)

		url := ztool.Str_FastConcat(`http://`, api_wy, `?id=`, c.MusicID, `&level=`, rquery, `&noCookie=true`)
		// jx.Debug(`Wy, Url: %v`, url)
		// wy源增加后端重试 默认3次
		for i := 0; true; i++ {
			_, err := ztool.Net_HttpReq(http.MethodGet, url, nil, header_wy, &resp)
			if err != nil {
				jx.Error(`HttpReq, Err: %s, ReTry: %v`, err, i)
				if i > 3 {
					msg = err.Error()
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
		if data.FreeTrialInfo != nil {
			// jx.Error("发生错误, 返回数据:\n%#v", resp)
			msg = `触发风控或专辑单独收费`
			return
		}
		if data.Level != rquery {
			msg = `实际音质不匹配`
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
			msg = err.Error()
			return
		}
		jx.Debug(`Mg, Resp: %+v`, resp)
		if link := resp.Data.PlayURL; link != `` {
			outlink = `https:` + link
		} // else {
		// 	jx.Debug(`Mg, Err: %#v`, resp)
		// }
	case s_kw:
		resp := kw_pool.Get().(*KwApi_Song)
		defer kw_pool.Put(resp)

		url := ztool.Str_FastConcat(`https://`, api_kw, `/`, c.MusicID, `?isMv=0&format=`, c.Extname, `&br=`, rquery, c.Extname, `&level=`)
		// jx.Debug(`Kw, Url: %s`, url)
		_, err := ztool.Net_HttpReq(http.MethodGet, url, nil, header_kw, &resp)
		if err != nil {
			msg = err.Error()
			return
		}
		jx.Debug(`Kw, Resp: %+v`, resp)
		if resp.Code != 200 || resp.Data.AudioInfo.Bitrate == `1` {
			// jx.Debug(`Kw, Err: %#v`, resp)
			msg = ztool.Str_FastConcat(`failed: `, resp.Msg)
			return
		}
		outlink = strings.Split(resp.Data.URL, `?`)[0]
	default:
		msg = `不支持的平台`
		return
	}
	return
}
