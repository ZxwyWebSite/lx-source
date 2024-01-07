package router

import (
	"lx-source/src/caches"
	"lx-source/src/env"
	"lx-source/src/middleware/auth"
	"lx-source/src/middleware/dynlink"
	"lx-source/src/middleware/loadpublic"
	"lx-source/src/middleware/resp"
	"lx-source/src/middleware/util"
	"lx-source/src/sources"
	"net/http"

	"github.com/ZxwyWebSite/ztool"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

var (
	// 默认音质
	defQuality = []string{`128k`, `320k`, `flac`, `flac24bit`}
	// 试听音质
	tstQuality = []string{`128k`}
	// 标准音质
	stdQuality = []string{`128k`, `320k`, `flac`}
)

// 自动生成支持的音质表
func loadQMap() [][]string {
	m := make([][]string, 6)
	// 0.wy
	if env.Config.Custom.Wy_Enable {
		m[0] = defQuality
	}
	// 1.mg
	m[1] = defQuality
	// 2.kw
	m[2] = stdQuality
	// 3.kg
	m[3] = tstQuality
	// 4.tx
	if env.Config.Custom.Tx_Enable {
		m[4] = stdQuality
	} else {
		m[4] = tstQuality
	}
	// 5.lx
	// m[sources.S_lx] = defQuality
	return m
}

// 载入路由
func InitRouter() *gin.Engine {
	r := gin.Default()
	qmap := loadQMap()
	// Gzip压缩
	if env.Config.Main.Gzip {
		r.Use(gzip.Gzip(gzip.DefaultCompression, gzip.WithExcludedPaths([]string{"/file/"})))
	}
	// 源信息
	r.GET(`/`, func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			`version`:   env.Version,         // 服务端程序版本
			`name`:      `lx-music-source`,   // 名称
			`msg`:       `Hello~::^-^::~v1~`, // Api大版本
			`developer`: []string{`Zxwy`},    // 开发者列表, 可在保留原作者的基础上添加你自己的名字?
			// 仓库地址
			`github`: `https://github.com/ZxwyWebSite/lx-source`,
			// 可用平台
			`source`: gin.H{
				sources.S_wy: qmap[0], //true,
				sources.S_mg: qmap[1], //true,
				sources.S_kw: qmap[2], //true,
				sources.S_kg: qmap[3], //[]string{`128k`, `320k`}, // 测试结构2, 启用时返回音质列表, 禁用为false
				sources.S_tx: qmap[4], //gin.H{ // "测试结构 不代表最终方式"
				// 	`enable`:   false,
				// 	`qualitys`: []string{`128k`, `320k`, `flac`, `flac24bit`},
				// },
				sources.S_lx: qmap[5],
			},
			// 自定义源脚本更新
			`script`: env.Config.Script,
		})
	})
	// 静态文件
	loadpublic.LoadPublic(r)
	// r.StaticFile(`/favicon.ico`, `public/icon.ico`)
	// r.StaticFile(`/lx-custom-source.js`, `public/lx-custom-source.js`)
	// 解析接口
	r.GET(`/link/:s/:id/:q`, auth.InitHandler(linkHandler)...)
	dynlink.LoadHandler(r)
	// r.GET(`/file/:t/:x/:f`, dynlink.FileHandler())
	// if cache, ok := caches.UseCache.(*localcache.Cache); ok {
	// 	r.Static(`/file`, cache.Path)
	// }
	// if env.Config.Cache.Mode == `local` {
	// 	r.Static(`/file`, env.Config.Cache.Local_Path)
	// }
	// 软件接口
	// api := r.Group(`/api`)
	// {
	// 	api.GET(`/lx`, lxHandler) // 洛雪音乐
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

	memHIT = `Memory HIT`    // 内存已命中
	memRej = `Memory Reject` // 内存已拒绝
)

// 外链解析
func linkHandler(c *gin.Context) {
	resp.Wrap(c, func() *resp.Resp {
		// 获取传入参数 检查合法性
		parms := util.ParaMap(c)
		// getParam := func(p string) string { return strings.TrimSuffix(strings.TrimPrefix(c.Param(p), `/`), `/`) } //strings.Trim(c.Param(p), `/`)
		s := parms[`s`]   //c.Param(`s`)   //getParam(`s`)   // source 平台 wy, mg, kw
		id := parms[`id`] //c.Param(`id`) //getParam(`id`) // sid 音乐ID wy: songmid, mg: copyrightId
		q := parms[`q`]   //c.Param(`q`)   //getParam(`q`)   // quality 音质 128k / 320k / flac / flac24bit
		env.Loger.NewGroup(`LinkQuery`).Debug(`s: %v, id: %v, q: %v`, s, id, q)
		if ztool.Chk_IsNilStr(s, q, id) {
			return &resp.Resp{Code: 6, Msg: `参数不全`} // http.StatusBadRequest
		}
		cquery := caches.NewQuery(s, id, q)
		// fmt.Printf("%+v\n", cquery)
		defer cquery.Free()
		// _, ok := sources.UseSource.Verify(cquery) // 获取请求音质 同时检测是否支持(如kw源没有flac24bit) qualitys[q][s]rquery
		// if !ok {
		// 	return &resp.Resp{Code: 6, Msg: `不支持的平台或音质`}
		// }

		// 查询内存
		clink, ok := env.Cache.Get(cquery.Query())
		if ok {
			if str, ok := clink.(string); ok {
				env.Loger.NewGroup(`MemCache`).Debug(`MemHIT [%q]=>[%q]`, cquery.Query(), str)
				if str == `` {
					return &resp.Resp{Code: 2, Msg: memRej} // 拒绝请求，当前一段时间内解析出错 `MemCache Reject`
				}
				return &resp.Resp{Msg: memHIT, Data: str} // `MemCache HIT`
			}
		}
		// 查询缓存
		var cstat bool
		if caches.UseCache != nil {
			cstat = caches.UseCache.Stat()
		}
		sc := env.Loger.NewGroup(`StatCache`)
		if cstat {
			sc.Debug(`Method: Get, Query: %v`, cquery.Query())
			if link := caches.UseCache.Get(cquery); link != `` {
				env.Cache.Set(cquery.Query(), link, 3600)
				return &resp.Resp{Msg: cacheHIT, Data: link}
			}
		} else {
			sc.Debug(`Disabled`)
		}
		// 解析歌曲外链
		outlink, emsg := sources.UseSource.GetLink(cquery)
		if emsg != `` {
			if emsg == sources.Err_Verify { // Verify Failed: 不支持的平台或音质
				return &resp.Resp{Code: 6, Msg: ztool.Str_FastConcat(emsg, `: 不支持的平台或音质`)}
			}
			env.Cache.Set(cquery.Query(), ``, 600) // 发生错误的10分钟内禁止再次查询
			return &resp.Resp{Code: 2, Msg: emsg}
		}
		// 缓存并获取直链 !(s == `kg` || (s == `tx` && !tx_en)) => (s != `kg` && (s != `tx` || tx_en))
		if outlink != `` && cstat && cquery.Source != sources.S_kg && (cquery.Source != sources.S_tx || env.Config.Custom.Tx_Enable) {
			sc.Debug(`Method: Set, Link: %v`, outlink)
			if link := caches.UseCache.Set(cquery, outlink); link != `` {
				env.Cache.Set(cquery.Query(), link, 3600)
				return &resp.Resp{Msg: cacheSet, Data: link}
			}
		}
		// 无法获取直链 直接返回原链接
		env.Cache.Set(cquery.Query(), outlink, 1200)
		return &resp.Resp{Msg: cacheMISS, Data: outlink}
	})
}
