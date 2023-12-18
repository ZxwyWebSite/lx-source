// 返回值处理
package resp

import (
	_ "embed"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 统一输出
/*
返回码对应表：
 0: http.StatusOK,                  // [200] 成功
 1: http.StatusForbidden,           // [403] IP被封禁
 2: http.StatusServiceUnavailable,  // [503] 获取失败
 3: http.StatusUnauthorized,        // [401] 验证失败
 4: http.StatusInternalServerError, // [500] 服务器内部错误
 5: http.StatusTooManyRequests,     // [429] 请求过于频繁
 6: http.StatusBadRequest,          // [400] 参数错误
*/
type Resp struct {
	Code int    `json:"code"` // 状态码 为兼容内置源设置 暂无实际作用 (1.0.2后已兼容Python版定义)
	Msg  string `json:"msg"`  // 提示or报错信息
	Data string `json:"data"` // 音乐URL
	Ext  string `json:"ext"`  // 其它信息
}

// 返回码对应列表 (参考Python版)
var statusMap = map[int]int{
	0: http.StatusOK,                  // 成功
	1: http.StatusForbidden,           // IP被封禁
	2: http.StatusServiceUnavailable,  // 获取失败
	3: http.StatusUnauthorized,        // 验证失败
	4: http.StatusInternalServerError, // 服务器内部错误
	5: http.StatusTooManyRequests,     // 请求过于频繁
	6: http.StatusBadRequest,          // 参数错误
}

//go:embed error.base64
var errormp3 string

// 返回请求
/*
 注：Code不为0时调用c.Abort()终止Handler
*/
func (o *Resp) Execute(c *gin.Context) {
	status, ok := statusMap[o.Code]
	if !ok {
		status = http.StatusOK
	}
	if o.Code != 0 {
		if o.Code == 2 /*&& o.Data == ``*/ {
			o.Data = errormp3
		}
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
