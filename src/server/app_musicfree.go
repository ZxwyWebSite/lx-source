package server

import (
	"lx-source/src/middleware/util"
	"lx-source/src/sources"
	"net/http"

	"github.com/ZxwyWebSite/ztool"
	"github.com/gin-gonic/gin"
)

func loadMusicFree(mf *gin.RouterGroup) {
	// 插件订阅
	mf.GET(`/subscribe`, func(c *gin.Context) {
		slist := []string{sources.S_wy, sources.S_mg, sources.S_kw, sources.S_kg, sources.S_tx, sources.S_lx}
		type plugins struct {
			Name    string `json:"name"`
			Url     string `json:"url"`
			Version string `json:"version"`
		}
		length := len(slist)
		plgs := make([]plugins, length)
		url := ztool.Str_FastConcat(util.GetPath(c.Request, `app/`), `public/musicfree/`)
		for i := 0; i < length; i++ {
			name := `lxs-` + slist[i]
			plgs[i] = plugins{
				Name:    name,
				Url:     ztool.Str_FastConcat(url, name, `.js`),
				Version: `0.0.0`,
			}
		}
		c.JSON(http.StatusOK, gin.H{`plugins`: plgs})
	})
}
