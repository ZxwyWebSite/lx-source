package server

import (
	"lx-source/src/env"
	"lx-source/src/middleware/dynlink"
	"lx-source/src/sources"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

var (
	accnum int64
	reqnum int64
	secnum int64
)

// 载入路由
func InitRouter() *gin.Engine {
	r := gin.Default()
	qmap := loadQMap()
	// Gzip压缩
	if env.Config.Main.Gzip {
		r.Use(gzip.Gzip(gzip.DefaultCompression, gzip.WithExcludedPaths([]string{"/file/"})))
	}
	// Cors跨域
	if env.Config.Main.Cors {
		r.Use(cors.Default())
	}
	startime := time.Now().Unix()
	// 源信息
	r.GET(`/`, func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			`version`: env.Version,         // 服务端程序版本
			`name`:    `lx-music-source`,   // 名称
			`msg`:     `Hello~::^-^::~v1~`, // Api大版本
			// `developer`: []string{`Zxwy`},    // 开发者列表, 可在保留原作者的基础上添加你自己的名字?
			// 仓库地址
			// `github`: `https://github.com/ZxwyWebSite/lx-source`,
			// 可用平台
			`source`: gin.H{
				sources.S_wy: qmap[sources.I_wy],
				sources.S_mg: qmap[sources.I_mg],
				sources.S_kw: qmap[sources.I_kw],
				sources.S_kg: qmap[sources.I_kg],
				sources.S_tx: qmap[sources.I_tx],
				sources.S_lx: qmap[sources.I_lx],
			},
			// 自定义源脚本更新
			`script`: env.Config.Script.Update, //env.Config.Script,
			// 数据统计
			`summary`: gin.H{
				`StartAt`: startime, // 启动时间
				`Accessn`: accnum,   // 访问次数
				`Request`: reqnum,   // 解析次数
				`Success`: secnum,   // 成功次数
			},
			// 验证方式
			`auth`: gin.H{
				`apikey`: env.Config.Auth.ApiKey_Enable,
			},
		})
	})
	// 静态文件
	loadPublic(r)
	// r.StaticFile(`/favicon.ico`, `public/icon.ico`)
	// r.StaticFile(`/lx-custom-source.js`, `public/lx-custom-source.js`)
	// 解析接口
	loadMusic(r)
	// r.GET(`/link/:s/:id/:q`, auth.InitHandler(linkHandler)...)
	dynlink.LoadHandler(r)
	// 动态链?
	// r.GET(`/file/:t/:x/:f`, dynlink.FileHandler())
	// if cache, ok := caches.UseCache.(*localcache.Cache); ok {
	// 	r.Static(`/file`, cache.Path)
	// }
	// if env.Config.Cache.Mode == `local` {
	// 	r.Static(`/file`, env.Config.Cache.Local_Path)
	// }
	// 功能接口
	// api := r.Group(`/api`)
	// {
	// 	api.GET(`/:s/:m/:q`) // {source}/{method}/{query}
	// }
	// 软件接口
	// app := r.Group(`/app`)
	// {
	// 	loadLxMusic(app.Group(`/lxmusic`))
	// 	loadMusicFree(app.Group(`/musicfree`))
	// }
	// 数据接口
	// r.GET(`/file/:t/:hq/:n`, func(c *gin.Context) {
	// 	c.String(http.StatusOK, time.Now().Format(`20060102150405`))
	// })
	// 暂不对文件接口进行验证 脚本返回链接无法附加请求头 只可在Get添加Query
	// g := r.Group(``)
	// {
	// 	g.Use(authHandler)
	// 	g.GET(`/link/:s/:id/:q`, linkHandler)
	// 	g.Static(`/file`, LocalCachePath)
	// }
	return r
}

// 数据返回格式
const (
	cacheHIT  = `Cache HIT`   // 缓存已命中
	cacheMISS = `Cache MISS`  // 缓存未命中
	cacheSet  = `Cache Seted` // 缓存已设置
	cacheFAIL = `Cache FAIL`  // 缓存未成功

	memHIT = `Memory HIT`    // 内存已命中
	memRej = `Memory Reject` // 内存已拒绝
)

// 外链解析
// func linkHandler(c *gin.Context) {
// 	resp.Wrap(c, func() *resp.Resp {
// 		// 获取传入参数 检查合法性
// 		arr := util.ParaArr(c, `s`, `id`, `q`)
// 		s, id, q := arr[0], arr[1], arr[2]
// 		// parms := util.ParaMap(c)
// 		// getParam := func(p string) string { return strings.TrimSuffix(strings.TrimPrefix(c.Param(p), `/`), `/`) } //strings.Trim(c.Param(p), `/`)
// 		// s := parms[`s`]   //c.Param(`s`)   //getParam(`s`)   // source 平台 wy, mg, kw
// 		// id := parms[`id`] //c.Param(`id`) //getParam(`id`) // sid 音乐ID wy: songmid, mg: copyrightId
// 		// q := parms[`q`]   //c.Param(`q`)   //getParam(`q`)   // quality 音质 128k / 320k / flac / flac24bit
// 		env.Loger.NewGroup(`LinkQuery`).Debug(`s: %v, id: %v, q: %v`, s, id, q).Free()
// 		if ztool.Chk_IsNilStr(s, q, id) {
// 			return &resp.Resp{Code: 6, Msg: `参数不全`} // http.StatusBadRequest
// 		}
// 		cquery := caches.NewQuery(s, id, q)
// 		cquery.Request = c.Request
// 		// fmt.Printf("%+v\n", cquery)
// 		defer cquery.Free()
// 		// _, ok := sources.UseSource.Verify(cquery) // 获取请求音质 同时检测是否支持(如kw源没有flac24bit) qualitys[q][s]rquery
// 		// if !ok {
// 		// 	return &resp.Resp{Code: 6, Msg: `不支持的平台或音质`}
// 		// }

// 		// 查询内存
// 		if clink, ok := env.Cache.Get(cquery.Query()); ok {
// 			if str, ok := clink.(string); ok {
// 				env.Loger.NewGroup(`MemCache`).Debug(`MemHIT [%q]=>[%q]`, cquery.Query(), str).Free()
// 				if str == `` {
// 					return &resp.Resp{Code: 2, Msg: memRej} // 拒绝请求，当前一段时间内解析出错 `MemCache Reject`
// 				}
// 				return &resp.Resp{Msg: memHIT, Data: str} // `MemCache HIT`
// 			}
// 		}
// 		// 查询缓存
// 		var cstat bool
// 		if caches.UseCache != nil {
// 			cstat = caches.UseCache.Stat()
// 		}
// 		sc := env.Loger.NewGroup(`StatCache`)
// 		defer sc.Free()
// 		if cstat {
// 			sc.Debug(`Method: Get, Query: %v`, cquery.Query())
// 			if link := caches.UseCache.Get(cquery); link != `` {
// 				env.Cache.Set(cquery.Query(), link, 3600)
// 				return &resp.Resp{Msg: cacheHIT, Data: link}
// 			}
// 		} else {
// 			sc.Debug(`Disabled`)
// 		}
// 		atomic.AddInt64(&reqnum, 1)
// 		// 解析歌曲外链
// 		outlink, emsg := sources.UseSource.GetLink(cquery)
// 		if emsg != `` {
// 			if emsg == sources.Err_Verify { // Verify Failed: 不支持的平台或音质
// 				return &resp.Resp{Code: 6, Msg: ztool.Str_FastConcat(emsg, `: 不支持的平台或音质`)}
// 			}
// 			env.Cache.Set(cquery.Query(), outlink, 600) // 发生错误的10分钟内禁止再次查询
// 			return &resp.Resp{Code: 2, Msg: emsg, Data: outlink}
// 		}
// 		atomic.AddInt64(&secnum, 1)
// 		// 缓存并获取直链 !(s == `kg` || (s == `tx` && !tx_en)) => (s != `kg` && (s != `tx` || tx_en))
// 		if outlink != `` && cstat && cquery.Source != sources.S_kg && (cquery.Source != sources.S_tx || env.Config.Custom.Tx_Enable) {
// 			sc.Debug(`Method: Set, Link: %v`, outlink)
// 			if link := caches.UseCache.Set(cquery, outlink); link != `` {
// 				env.Cache.Set(cquery.Query(), link, 3600)
// 				return &resp.Resp{Msg: cacheSet, Data: link}
// 			}
// 		}
// 		// 无法获取直链 直接返回原链接
// 		env.Cache.Set(cquery.Query(), outlink, 1200)
// 		return &resp.Resp{Msg: cacheMISS, Data: outlink}
// 	})
// }
