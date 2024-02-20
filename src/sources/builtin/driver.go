// 内置解析源
package builtin

import (
	"lx-source/src/caches"
	"lx-source/src/env"
	"lx-source/src/sources"
	"lx-source/src/sources/custom/kg"
	"lx-source/src/sources/custom/kw"
	"lx-source/src/sources/custom/mg"
	"lx-source/src/sources/custom/tx"
	"lx-source/src/sources/custom/wy"
	wm "lx-source/src/sources/custom/wy/modules"
	"lx-source/src/sources/example"
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
	wy_pool *sync.Pool
	mg_pool = &sync.Pool{New: func() any { return new(MgApi_Song) }}
	// kw_pool = &sync.Pool{New: func() any { return new(KwApi_Song) }}
	// kg_pool = &sync.Pool{New: func() any { return new(KgApi_Song) }}
	// tx_pool = &sync.Pool{New: func() any { return new(res_tx) }}
	wv_pool *sync.Pool
)

func init() {
	env.Inits.Add(func() {
		if env.Config.Source.Enable_Wy {
			wy_pool = &sync.Pool{New: func() any { return new(wm.PlayInfo) }}
			if env.Config.Source.MusicIdVerify {
				wv_pool = &sync.Pool{New: func() any { return new(wm.VerifyInfo) }}
			}
		}
	})
}

// 查询
func (s *Source) GetLink(c *caches.Query) (outlink string, msg string) {
	rquery, ok := s.Verify(c)
	if !ok /*&& c.Source != `tx`*/ {
		msg = sources.Err_Verify //`Verify Failed`
		return
	}
	// var outlink string
	jx := env.Loger.NewGroup(`Sources`) //sources.Loger.AppGroup(`builtin`) //env.Loger.NewGroup(`JieXiApis`)
	defer jx.Free()
	switch c.Source {
	case sources.S_wy:
		if !env.Config.Source.Enable_Wy {
			msg = sources.ErrDisable
			return
		}
		if wy.Url != nil {
			outlink, msg = wy.Url(c.MusicID, c.Quality)
			break
		}
		// 可用性验证
		if env.Config.Source.MusicIdVerify {
			vef := wv_pool.Get().(*wm.VerifyInfo)
			defer wv_pool.Put(vef)
			vurl := ztool.Str_FastConcat(`https://`, example.Vef_wy, `&id=`, c.MusicID)
			_, err := ztool.Net_HttpReq(http.MethodGet, vurl, nil, example.Header_wy, &vef)
			if err != nil {
				jx.Error(`Wy, VefReq: %s`, err)
				msg = sources.ErrHttpReq
				return
			}
			jx.Debug(`Wy, Vef: %+v`, vef)
			if vef.Code != 200 || !vef.Success {
				msg = ztool.Str_FastConcat(`暂不可用：`, vef.Message)
				return
			}
		}
		// 获取外链
		resp := wy_pool.Get().(*wm.PlayInfo)
		defer wy_pool.Put(resp)
		// 分流逻辑 (暂无其它节点)
		// urls := [...]string{
		// 	ztool.Str_FastConcat(`http://`, api_wy, `?id=`, c.MusicID, `&level=`, rquery, `&noCookie=true`),
		// 	ztool.Str_FastConcat(`https://`, api_wy, `&id=`, c.MusicID, `&level=`, rquery, `&encodeType=`, c.Extname),
		// }
		// url := urls[rand.Intn(len(urls))]
		url := ztool.Str_FastConcat(
			`https://`, example.Api_wy, `&id=`, c.MusicID, `&level=`, rquery,
			`&timestamp=`, strconv.FormatInt(time.Now().UnixMilli(), 10),
		)
		// jx.Debug(`Wy, Url: %v`, url)
		// wy源增加后端重试 默认3次
		for i := 0; true; i++ {
			_, err := ztool.Net_HttpReq(http.MethodGet, url, nil, example.Header_wy, &resp)
			if err != nil {
				jx.Error(`HttpReq, Err: %s, ReTry: %v`, err, i)
				if i > 3 {
					jx.Error(`Wy, HttpReq: %s`, err)
					msg = sources.ErrHttpReq
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
			msg = `触发风控或专辑单独收费: ` + strconv.Itoa(data.Code)
			return
		}
		if data.Level != rquery {
			msg = ztool.Str_FastConcat(`实际音质不匹配: `, rquery, ` <= `, data.Level) // 实际音质不匹配: exhigh <= standard
			if !env.Config.Source.ForceFallback {
				return
			}
		}
		// jx.Info(`WyLink, RealQuality: %v`, data.Level)
		outlink = data.URL
	case sources.S_mg:
		if !env.Config.Source.Enable_Mg {
			msg = sources.ErrDisable
			return
		}
		if len(c.MusicID) != 11 {
			msg = sources.E_VefMusicId
			return
		}
		if mg.Url != nil {
			outlink, msg = mg.Url(c.MusicID, c.Quality)
			break
		}
		resp := mg_pool.Get().(*MgApi_Song)
		defer mg_pool.Put(resp)

		url := ztool.Str_FastConcat(`https://`, example.Api_mg, `?copyrightId=`, c.MusicID, `&type=`, rquery)
		// jx.Debug(`Mg, Url: %v`, url)
		_, err := ztool.Net_HttpReq(http.MethodGet, url, nil, example.Header_mg, &resp)
		if err != nil {
			jx.Error(`Mg, HttpReq: %s`, err)
			msg = sources.ErrHttpReq
			return
		}
		jx.Debug(`Mg, Resp: %+v`, resp)
		if link := resp.Data.PlayURL; link != `` {
			outlink = `https:` + link
		} else {
			msg = ztool.Str_FastConcat(resp.Code, `: `, resp.Msg)
		}
	case sources.S_kw:
		if !env.Config.Source.Enable_Kw {
			msg = sources.ErrDisable
			return
		}
		outlink, msg = kw.Url(c.MusicID, c.Quality)
		// if emsg != `` {
		// 	msg = emsg
		// 	return
		// }
		// outlink = ourl
	case sources.S_kg:
		if !env.Config.Source.Enable_Kg {
			msg = sources.ErrDisable
			return
		}
		if len(c.MusicID) != 32 {
			msg = sources.E_VefMusicId
			return
		}
		outlink, msg = kg.Url(c.MusicID, c.Quality)
		// if emsg != `` {
		// 	msg = emsg
		// 	return
		// }
		// outlink = ourl
	// case sources.S_kg:
	// 	if !env.Config.Custom.Kg_Enable {
	// 		msg = sources.ErrDisable
	// 		return
	// 	}
	// 	resp := kg_pool.Get().(*KgApi_Song)
	// 	defer kg_pool.Put(resp)

	// 	// sep := strings.Split(c.MusicID, `-`) // 分割 Hash-Album 如 6DC276334F56E22BE2A0E8254D332B45-13097991
	// 	// alb := func() string {
	// 	// 	if len(sep) >= 2 {
	// 	// 		return sep[1]
	// 	// 	}
	// 	// 	return ``
	// 	// }()
	// 	sep := c.Split()
	// 	url := ztool.Str_FastConcat(api_kg, `&hash=`, sep[0], `&album_id=`, sep[1], `&_=`, strconv.FormatInt(time.Now().UnixMilli(), 10))
	// 	// jx.Debug(`Kg, Url: %s`, url)
	// 	_, err := ztool.Net_HttpReq(http.MethodGet, url, nil, nil, &resp)
	// 	if err != nil {
	// 		jx.Error(`Kg, HttpReq: %s`, err)
	// 		msg = sources.ErrHttpReq
	// 		return
	// 	}
	// 	jx.Debug(`Kg, Resp: %+v`, resp)
	// 	if resp.ErrCode != 0 {
	// 		msg = ztool.Str_FastConcat(`Error: `, strconv.Itoa(resp.ErrCode))
	// 		return
	// 	}
	// 	var data KgApi_Data
	// 	err = ztool.Val_MapToStruct(resp.Data, &data)
	// 	if err != nil {
	// 		msg = err.Error()
	// 		return
	// 	}
	// 	if data.PlayBackupURL == `` {
	// 		if data.PlayURL == `` {
	// 			msg = sources.ErrNoLink
	// 			return
	// 		}
	// 		outlink = data.PlayURL
	// 	}
	// 	outlink = data.PlayBackupURL
	case sources.S_tx:
		if !env.Config.Source.Enable_Tx {
			msg = sources.ErrDisable
			return
		}
		if len(c.MusicID) != 14 {
			msg = sources.E_VefMusicId
			return
		}
		outlink, msg = tx.Url(c.MusicID, c.Quality)
		// if emsg != `` {
		// 	msg = emsg
		// 	return
		// }
		// outlink = ourl
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
