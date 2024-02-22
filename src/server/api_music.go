package server

import (
	"lx-source/src/caches"
	"lx-source/src/env"
	"lx-source/src/middleware/auth"
	"lx-source/src/middleware/resp"
	"lx-source/src/middleware/util"
	"lx-source/src/sources"
	"lx-source/src/sources/custom"
	"strings"
	"sync/atomic"

	"github.com/ZxwyWebSite/ztool"
	"github.com/gin-gonic/gin"
)

// type (
// 	Context struct {
// 		Parms map[string]string
// 	}
// 	Source interface {
// 		Link(c *Context) (string, error)
// 		Lyric(c *Context)
// 	}
// )

func loadMusic(api gin.IRouter) {
	// /{method}/{source}/{musicId}/{?quality}
	api.GET(`/:m/:s/:id/*q`, auth.InitHandler(musicHandler)...)
}

func musicHandler(c *gin.Context) {
	resp.Wrap(c, func() *resp.Resp {
		// 获取请求参数 (测试用Array会不会提升性能)
		arr := util.ParaArr(c, `s`, `m`, `id`, `q`)
		ps, pm, pid, pq := arr[0], arr[1], arr[2], strings.TrimPrefix(arr[3], `/`)
		out := &resp.Resp{Code: 0} // 默认Code:6 (参数错误)
		loger := env.Loger.NewGroup(`MusicHandler`)
		defer loger.Free()
		loger.Debug(`s:'%v', m:'%v', id:'%v', q:'%v'`, ps, pm, pid, pq)
		// 定位音乐源
		var source custom.Source
		var active bool // 是否激活(自定义账号)
		switch ps {
		case sources.S_wy:
			active = env.Config.Custom.Wy_Enable
			source = custom.WySource
		case sources.S_mg:
			active = env.Config.Custom.Mg_Enable
			source = custom.MgSource
		case sources.S_kw:
			active = env.Config.Custom.Kw_Enable
			source = custom.KwSource
		case sources.S_kg:
			active = env.Config.Custom.Kg_Enable
			source = custom.KgSource
		case sources.S_tx:
			active = env.Config.Custom.Tx_Enable
			source = custom.TxSource
		case sources.S_lx:
			source = custom.LxSource
		default:
			out.Code = 6
			out.Msg = ztool.Str_FastConcat(`无效源参数:'`, ps, `'`)
			return out
		}
		if source == nil {
			out.Code = 6
			out.Msg = sources.ErrDisable
			return out
		}
		if !source.Vef(&pid) {
			out.Code = 6
			out.Msg = sources.E_VefMusicId
			return out
		}
		// 查询内存缓存
		cquery := strings.Join([]string{pm, ps, pid, pq}, `/`)
		loger.Debug(`MemoGet: %v`, cquery)
		if cdata, ok := env.Cache.Get(cquery); ok {
			loger.Debug(`MemoHIT: %q`, cdata)
			if cdata == `` {
				out.Code = 2
				out.Msg = memRej
			} else {
				out.Msg = memHIT
				out.Data = cdata
			}
			return out
		}
		// 定位源方法
		switch pm {
		case `url`, `link`:
			// 查询文件缓存
			var cstat bool
			if caches.UseCache != nil {
				cstat = caches.UseCache.Stat()
			}
			uquery := caches.NewQuery(ps, pid, pq)
			defer uquery.Free()
			if cstat {
				loger.Debug(`FileGet: %v`, uquery.Query())
				if olink := caches.UseCache.Get(uquery); olink != `` {
					env.Cache.Set(cquery, olink, sources.C_lx)
					out.Msg = cacheHIT
					out.Data = olink
					return out
				}
			}
			// 解析歌曲外链
			atomic.AddInt64(&reqnum, 1)
			out.Data, out.Msg = source.Url(pid, pq)
			if out.Data != `` {
				// 缓存并获取直链
				atomic.AddInt64(&secnum, 1)
				if out.Msg == `` {
					if cstat && active {
						loger.Debug(`FileSet: %v`, out.Data)
						if link := caches.UseCache.Set(uquery, out.Data.(string)); link != `` {
							env.Cache.Set(cquery, link, sources.C_lx)
							out.Msg = cacheSet
							out.Data = link
							return out
						}
						out.Msg = cacheFAIL
					} else {
						out.Msg = cacheMISS
					}
				}
				// 无法获取直链 直接返回原链接
				env.Cache.Set(cquery, out.Data, source.Exp()-300)
				return out
			}
		case `lrc`, `lyric`:
			out.Data, out.Msg = source.Lrc(pid)
		case `pic`, `cover`:
			out.Data, out.Msg = source.Pic(pid)
		default:
			out.Code = 6
			out.Msg = ztool.Str_FastConcat(`无效源方法:'`, pm, `'`)
			return out
		}
		if out.Msg != `` {
			out.Code = 2
			env.Cache.Set(cquery, out.Data, 600)
		}
		return out
	})
}
