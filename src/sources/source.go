package sources

import (
	"lx-source/src/caches"
)

// var Loger = env.Loger.NewGroup(`Sources`) // JieXiApis
const (
	Err_Verify = `Verify Failed`
)

// 源查询接口
/*
 Origin:
 首先调用Verify验证源是否支持
 再尝试查询缓存
 无缓存则解析链接

 参考Python版:
 不验证当前源是否支持，直接查询缓存
 验证部分放到GetLink里

*/
type Source interface {
	Verify(*caches.Query) (string, bool)    // 验证是否可用 <查询参数> <rquery,ok>
	GetLink(*caches.Query) (string, string) // 查询获取链接 <查询参数> <链接,信息>
}

// 默认空接口
type NullSource struct{}

func (*NullSource) Verify(*caches.Query) (string, bool)    { return ``, false }
func (*NullSource) GetLink(*caches.Query) (string, string) { return ``, `NullSource` }

var UseSource Source = &NullSource{} // = &builtin.Source{}

// 统一错误
// type Error struct {
// 	msg string
// }

// func (e *Error) Error() string {
// 	return ztool.Str_FastConcat(e.msg)
// }
