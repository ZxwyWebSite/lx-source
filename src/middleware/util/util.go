package util

import (
	"lx-source/src/env"
	"net/http"

	"github.com/ZxwyWebSite/ztool"
	"github.com/gin-gonic/gin"
)

// 将路由参数转为Map
/*
 `:s/:id/:q` -> {
	`s`:  `source`,
	`id`: `musicId`,
	`q`:  `quality`,
 }
*/
func ParaMap(c *gin.Context) map[string]string {
	parmlen := len(c.Params)
	parms := make(map[string]string, parmlen)
	for i := 0; i < parmlen; i++ {
		parms[c.Params[i].Key] = c.Params[i].Value
	}
	return parms
}

// 将路由参数转为Array
/*
 ParaArr(c, `id`, `s`, `xxx`) => [
	`musicId`,
	`source`,
	``,
 ]
*/
func ParaArr(c *gin.Context, s ...string) []string {
	parmlen := len(c.Params)
	parslen := len(s)
	out := make([]string, parslen)
	for im := 0; im < parmlen; im++ {
		obj := c.Params[im]
		for is := 0; is < parslen; is++ {
			if s[is] == obj.Key {
				out[is] = obj.Value
			}
		}
	}
	return out
}

var pathCache string

func init() {
	env.Inits.Add(func() {
		if !env.Config.Main.NgProxy {
			pathCache = `/`
		}
	})
}

// 动态获取相对路径 <Ctx, 特征> <路径>
/*
 HOST = `192.168.10.22:1011`
 URI  = `/path/to/lxs/link/wy/2049512697/flac`
 sub  = `link/`
 -> http://192.168.10.22:1011/path/to/lxs/
*/
func GetPath(c *http.Request, sub string) string {
	// 从缓存读取相对路径 `/path/to/lxs/` or `/`
	if pathCache == `` {
		pathCache = ztool.Str_Before(c.RequestURI, sub)
	}
	return ztool.Str_FastConcat(`http://`, c.Host, pathCache)
}
