// 全局验证
package auth

import (
	"lx-source/src/env"
	"lx-source/src/middleware/resp"

	"github.com/gin-gonic/gin"
)

// 请求验证
func AuthHandler(c *gin.Context) {
	loger := env.Loger.NewGroup(`AuthHandler`)
	resp.Wrap(c, func() *resp.Resp {
		// ApiKey
		if env.Config.Auth.ApiKey_Enable {
			if auth := c.Request.Header.Get(`X-LxM-Auth`); auth != env.Config.Auth.ApiKey_Value {
				loger.Debug(`验证失败: %q`, auth)
				return &resp.Resp{Code: 3, Msg: `验证Key失败, 请联系网站管理员`}
			}
		}
		return nil
	})
	// c.Next()
}
