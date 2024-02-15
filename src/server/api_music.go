package server

import (
	"lx-source/src/caches"
	"lx-source/src/env"
	"lx-source/src/middleware/resp"
	"lx-source/src/middleware/util"
	"lx-source/src/sources"
	"lx-source/src/sources/custom"
	"strings"

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

func loadMusic(api *gin.RouterGroup) {
	// /{method}/{source}/{musicId}/{?quality}
	api.GET(`/:m/:s/:id/*q`, musicHandler)

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
		// 查询内存缓存
		cquery := strings.Join([]string{pm, ps, pid, pq}, `/`)
		loger.Debug(cquery)
		if clink, ok := env.Cache.Get(cquery); ok {
			if cstr, ok := clink.(string); ok {
				loger.Debug(`MemHIT [%q]=>[%q]`, cquery, cstr)
				if cstr == `` {
					out.Code = 2
					out.Msg = `Memory Reject`
				} else {
					out.Msg = `Memory HIT`
					out.Data = cstr
				}
				return out
			}
		}
		// 定位音乐源
		var source custom.Source
		switch ps {
		case sources.S_wy:
			source = custom.WySource
		case sources.S_mg:
			source = custom.MgSource
		case sources.S_kw:
			source = custom.KwSource
		case sources.S_kg:
			source = custom.KgSource
		case sources.S_tx:
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
		// 定位源方法
		if !source.Vef(pid) {
			out.Code = 6
			out.Msg = sources.E_VefMusicId
			return out
		}
		switch pm {
		case `url`, `link`:
			// 查询文件缓存
			if caches.UseCache.Stat() {
				uquery := caches.NewQuery(ps, pid, pq)
				defer uquery.Free()
				if olink := caches.UseCache.Get(uquery); olink != `` {
					env.Cache.Set(cquery, olink, 3600)
					out.Msg = `Cache HIT`
					out.Data = olink
					return out
				}
			}
			out.Msg = `No Link`
			// out.Data, out.Msg = source.Url(pid, pq)
		case `lrc`, `lyric`:
			out.Data, out.Msg = source.Lrc(pid)
		case `pic`, `cover`:
			out.Data, out.Msg = source.Pic(pid)
		default:
			out.Code = 6
			out.Msg = ztool.Str_FastConcat(`无效源方法:'`, pm, `'`)
			// return
		}
		// 缓存并获取直链
		if out.Data != `` {
		}
		return out
	})
}
