// 全局验证
package auth

import (
	"lx-source/src/env"
	"lx-source/src/middleware/resp"
	"sync/atomic"
	"time"

	"github.com/gin-gonic/gin"
)

type (
	RateLimit struct {
		Tim int64  // 创建时间 (注：原子操作64位数据需放在结构体第一位或保证8字节对齐，否则不兼容32位平台 https://pkg.go.dev/sync/atomic#pkg-note-BUG)
		Num uint32 // 请求次数
	}
)

func InitHandler(h gin.HandlerFunc) (out []gin.HandlerFunc) {
	loger := env.Loger.NewGroup(`AuthHandler`)
	// ApiKey 请求头验证
	if env.Config.Auth.ApiKey_Enable {
		loger.Debug(`ApiKeyAuth Enabled`)
		out = append(out, func(c *gin.Context) {
			resp.Wrap(c, func() *resp.Resp {
				if auth := c.Request.Header.Get(`X-LxM-Auth`); auth != env.Config.Auth.ApiKey_Value {
					loger.Debug(`验证失败: %q`, auth)
					return &resp.Resp{Code: 3, Msg: `验证Key失败, 请联系网站管理员`}
				}
				return nil
			})
		})
	}
	// RateLimit 速率限制
	/*
	 逻辑：
	  记录访问者ip，到内存(缓存)中查找"块"，没有则新建，
	  检测"块"是否过期，否则新建，
	   // 判断ip是否在白名单内，是则直接放行
	  判断请求数+1是否大于限制，True: 429 请求过快，请稍后重试
	   // 判断是否超出容忍限度，是则封禁ip (暂未实现)
	   // 超过容忍限度每次请求增加一个Block的时间
	  继续执行后续Handler
	*/
	if env.Config.Auth.RateLimit_Enable {
		loger.Debug(`RateLimit Enabled`)
		loger.Info(`已启用速率限制，当前配置 %v/%v`, env.Config.Auth.RateLimit_Single, env.Config.Auth.RateLimit_Block)
		newRateLimit := func() *RateLimit { return &RateLimit{Tim: time.Now().Unix(), Num: 1} }
		block_int64 := int64(env.Config.Auth.RateLimit_Block)
		block_mem := int(env.Config.Auth.RateLimit_Block * env.Config.Auth.RateLimit_BanNum)
		bannum := env.Config.Auth.RateLimit_Single + env.Config.Auth.RateLimit_BanNum
		bantim := int64(env.Config.Auth.RateLimit_BanTim)
		out = append(out, func(c *gin.Context) {
			resp.Wrap(c, func() *resp.Resp {
				rip := c.RemoteIP()
				if rip == `` {
					rip = `0.0.0.0`
				}
				cip, ok := env.Cache.Get(rip)
				loger.Debug(`GetMemRip: %v`, rip)
				if ok {
					if oip, ok := cip.(*RateLimit); ok {
						loger.Debug(`GetMemOut: %+v`, oip)
						if oip.Tim+block_int64 > time.Now().Unix() {
							oi := atomic.AddUint32(&oip.Num, 1)
							if oi > env.Config.Auth.RateLimit_Single {
								if oi > bannum {
									atomic.AddInt64(&oip.Tim, bantim)
								}
								return &resp.Resp{Code: 5, Msg: `请求过快，请稍后重试`}
							}
							return nil
						}
					}
				}
				val := newRateLimit()
				if err := env.Cache.Set(rip, val, block_mem); err != nil {
					loger.Error(`写入内存: %s`, err)
					return &resp.Resp{Code: 4, Msg: `速率限制内部异常，请联系网站管理员`}
				}
				loger.Debug(`SetMemVal: %+v`, val)
				return nil
			})
		})
	}
	return append(out, h)
}

// 请求验证
// func AuthHandler(c *gin.Context) {
// 	loger := env.Loger.NewGroup(`AuthHandler`)
// 	resp.Wrap(c, func() *resp.Resp {
// 		// ApiKey
// 		if env.Config.Auth.ApiKey_Enable {
// 			if auth := c.Request.Header.Get(`X-LxM-Auth`); auth != env.Config.Auth.ApiKey_Value {
// 				loger.Debug(`验证失败: %q`, auth)
// 				return &resp.Resp{Code: 3, Msg: `验证Key失败, 请联系网站管理员`}
// 			}
// 		}
// 		return nil
// 	})
// 	// c.Next()
// }
