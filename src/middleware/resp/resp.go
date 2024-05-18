// 返回值处理
package resp

import (
	"lx-source/src/env"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 统一输出
/*
返回码对应表 (参考Python版)：
 0: http.StatusOK,                  // [200] 成功
 1: http.StatusForbidden,           // [403] IP被封禁
 2: http.StatusServiceUnavailable,  // [503] 获取失败
 3: http.StatusUnauthorized,        // [401] 验证失败
 4: http.StatusInternalServerError, // [500] 服务器内部错误
 5: http.StatusTooManyRequests,     // [429] 请求过于频繁
 6: http.StatusBadRequest,          // [400] 参数错误
*/
type Resp struct {
	Code int    `json:"code"`          // 状态码 为兼容内置源设置 暂无实际作用 (1.0.2后已兼容Python版定义)
	Msg  string `json:"msg"`           // 提示or报错信息
	Data any    `json:"data"`          // 音乐URL
	Ext  any    `json:"ext,omitempty"` // 其它信息
}

// 获取失败默认音频
// var ErrMp3 = `https://r2eu.zxwy.link/gh/lx-source/static/error.mp3`

// 返回请求
/*
 注：Code不为0时调用c.Abort()终止Handler
*/
func (o *Resp) Execute(c *gin.Context) {
	// StatusCode转换 (小分支switch快, 大分支map快)
	var status int
	switch o.Code {
	case 0:
		status = http.StatusOK
	case 1:
		status = http.StatusForbidden
	case 2:
		status = http.StatusServiceUnavailable
		if o.Data == nil || o.Data == `` {
			o.Data = env.Config.Main.ErrMp3 //ErrMp3
		}
	case 3:
		status = http.StatusUnauthorized
	case 4:
		status = http.StatusInternalServerError
	case 5:
		status = http.StatusTooManyRequests
	case 6:
		status = http.StatusBadRequest
	default:
		status = http.StatusOK
	}
	if o.Code != 0 {
		// if o.Code == 2 /*&& o.Data == ``*/ {
		// 	o.Data = ErrMp3
		// }
		c.Abort()
	}
	c.JSON(status, o)
}

// 包装请求并自动处理
/*
 注：返回nil以继续执行下一个Handler
*/
func Wrap(c *gin.Context, f func() *Resp) {
	if r := f(); r != nil {
		r.Execute(c)
	}
}

// func Wrap2(c *gin.Context, p []string, f func([]string) *Resp) {
// 	if r := f(util.ParaArr(c)); r != nil {
// 		r.Execute(c)
// 	}
// }
