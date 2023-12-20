package dynlink

import (
	"lx-source/src/caches"
	"lx-source/src/caches/localcache"
	"lx-source/src/env"

	// "lx-source/src/middleware/util"
	// "net/http"

	// "github.com/ZxwyWebSite/ztool"
	"github.com/gin-gonic/gin"
)

type DynLink struct {
	Mode uint8
	Link string
}

func LoadHandler(r *gin.Engine) {
	loger := env.Loger.NewGroup(`DynLink`)
	if cache, ok := caches.UseCache.(*localcache.Cache); ok {
		loger.Debug(`UseStatic`)
		r.Static(`/file`, cache.Path)
	}
	// 动态链暂未完成...
}

// func FileHandler() gin.HandlerFunc {
// 	loger := env.Loger.NewGroup(`DynLink`)
// 	// 为了兼容原静态链，必须设置3个参数
// 	// file/:{time.unix}/:{md5(cquery)}/:{fname}  1703006183//77792f3434343730363834382f3332306b.mp3
// 	// file/:date/:second/:fname 20231219/1703006183/77792f3434343730363834382f3332306b.mp3
// 	env.Cache.Set(`20211008/hello/test.mp3`, DynLink{Link: `/www/wwwroot/lx-source/data/cache/wy/3203127/320k.mp3`}, 0)

// 	if env.Config.Cache.LinkMode == `dynamic` || env.Config.Cache.LinkMode == `2` /*|| true*/ {
// 		loger.Debug(`UseDynamic`)
// 		return func(c *gin.Context) {
// 			parms := util.ParaMap(c)
// 			t, x, f := parms[`t`], parms[`x`], parms[`f`]
// 			if clink, ok := env.Cache.Get(ztool.Str_FastConcat(t, `/`, x, `/`, f)); ok {
// 				if dyn, ok := clink.(DynLink); ok {
// 					if dyn.Mode == 0 {
// 						c.File(ztool.Str_FastConcat(dyn.Link))
// 						return
// 					}
// 					c.Redirect(http.StatusFound, dyn.Link)
// 					return
// 				}
// 			}
// 			c.AbortWithStatus(http.StatusNotFound)
// 		}
// 	}
// 	if cache, ok := caches.UseCache.(*localcache.Cache); ok {
// 		loger.Debug(`UseStatic`)
// 		return func(c *gin.Context) {
// 			parms := util.ParaMap(c)
// 			t, x, f := parms[`t`], parms[`x`], parms[`f`]
// 			c.File(ztool.Str_FastConcat(cache.Path, `/`, t, `/`, x, `/`, f))
// 		}
// 	}
// 	return func(c *gin.Context) {
// 		c.AbortWithStatus(http.StatusNotFound)
// 	}
// }
