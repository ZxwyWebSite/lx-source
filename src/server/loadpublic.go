// 静态资源
package server

import (
	"bytes"
	"embed"
	"fmt"
	"io"
	"io/fs"
	"lx-source/src/env"
	"net/http"
	"path/filepath"

	"github.com/ZxwyWebSite/ztool"
	"github.com/ZxwyWebSite/ztool/x/bytesconv"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
)

//go:embed public
var publicEM embed.FS // 打包默认Public目录 src/server/public

// 载入Public目录并设置路由
func loadPublic(r *gin.Engine) {
	pf := env.Loger.NewGroup(`PublicFS`)
	dir := ztool.Str_FastConcat(env.RunPath, `/data/public`)
	publicFS, err := fs.Sub(publicEM, `public`)
	var httpFS http.FileSystem = http.FS(publicFS)
	if err != nil {
		pf.Fatal(`内置Public目录载入错误: %s, 请尝试重新编译`, err)
	}
	if !ztool.Fbj_IsExists(dir) {
		pf.Info(`不存在Public目录, 释放默认静态文件`)
		walk := func(relPath string, d fs.DirEntry, err error) error {
			if err != nil {
				return fmt.Errorf(`无法获取[%q]的信息: %s`, relPath, err)
			}
			if !d.IsDir() {
				out, err := ztool.Fbj_CreatFile(filepath.Join(dir, relPath))
				if err != nil {
					return fmt.Errorf(`无法创建文件[%q]: %s`, relPath, err)
				}
				defer out.Close()
				pf.Debug(`导出 [%q]...`, relPath)
				obj, err := publicFS.Open(relPath)
				if err != nil {
					return fmt.Errorf(`无法打开文件[%q]: %s`, relPath, err)
				}
				if _, err := io.Copy(out, obj); err != nil {
					return fmt.Errorf(`无法写入文件[%q]: %s`, relPath, err)
				}
			}
			return nil
		}
		if err := fs.WalkDir(publicFS, `.`, walk); err != nil {
			pf.Error(`无法释放静态文件: %s`, err)
			// pf.Warn(`正在使用内置Public目录, 将无法自定义静态文件`)
			// httpFS = http.FS(publicFS)
		} else {
			pf.Info(`全部静态资源导出完成, 祝你使用愉快 ^_^`)
		}
	}
	pf.Free()
	// 使用本地public目录
	// httpFS = gin.Dir(dir, false)
	// r.GET(`/:file`, func(c *gin.Context) {
	// 	file := c.Param(`file`)
	// 	switch file {
	// 	case `favicon.ico`:
	// 		c.FileFromFS(`icon.ico`, httpFS)
	// 	// case `lx-custom-source.js`:
	// 	// 	c.FileFromFS(`lx-custom-source.js`, http.FS(publicFS))
	// 	default:
	// 		c.FileFromFS(file, httpFS)
	// 	}
	// })
	// 自动填写源脚本参数
	if env.Config.Script.Auto > 0 {
		file, _ := publicFS.Open(`lx-custom-source.js`)
		data, _ := io.ReadAll(file)
		file.Close()
		data = bytes.Replace(data,
			bytesconv.StringToBytes(`http://127.0.0.1:1011/`),
			bytesconv.StringToBytes(env.Config.Cache.Local_Bind), 1,
		)
		if env.Config.Auth.ApiKey_Enable && env.Config.Script.Auto >= 2 {
			data = bytes.Replace(data,
				bytesconv.StringToBytes(`apipass = ''`),
				bytesconv.StringToBytes(ztool.Str_FastConcat(
					`apipass = '`, env.Config.Auth.ApiKey_Value, `'`,
				)), 1,
			)
		}
		r.GET(`/lx-custom-source.js`, func(c *gin.Context) {
			c.Render(http.StatusOK, render.Data{
				ContentType: `text/javascript; charset=utf-8`,
				Data:        data,
			})
		})
	} else {
		r.StaticFileFS(`/lx-custom-source.js`, `lx-custom-source.js`, httpFS)
	}
	r.StaticFileFS(`/favicon.ico`, `lx-icon.ico`, httpFS)
	r.StaticFS(`/public`, httpFS)
}
