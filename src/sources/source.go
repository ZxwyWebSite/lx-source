package sources

import (
	"lx-source/src/caches"
)

// var Loger = env.Loger.NewGroup(`Sources`) // JieXiApis

const (
	Err_Verify = `Verify Failed`
	// 通用音质
	Q_128k = `128k`
	Q_320k = `320k`
	Q_flac = `flac`
	Q_fl24 = `flac24bit`
	// 文件扩展
	// X_aac  = `aac`
	X_mp3 = `mp3`
	// X_flac = Q_flac
	// 通用平台
	S_wy = `wy` // 小芸
	S_mg = `mg` // 小蜜
	S_kw = `kw` // 小蜗
	S_kg = `kg` // 小枸
	S_tx = `tx` // 小秋
	S_lx = `lx` // 小洛 (预留)
	// 常用错误
	E_QNotSupport = `不支持的音质`
	E_QNotMatch   = `实际音质不匹配`
	E_NoLink      = `无法获取音乐链接`
	E_VefMusicId  = `音乐ID校验失败`
	// 内置错误
	ErrHttpReq = `无法连接解析接口`
	ErrNoLink  = `无法获取试听链接`
	ErrDisable = `该音乐源已被禁用`
)

const (
	I_wy = iota
	I_mg
	I_kw
	I_kg
	I_tx
	I_lx
)

var (
	S_al = []string{S_wy, S_mg, S_kw, S_kg, S_tx, S_lx} // 全部平台
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
// type (
// 	ErrDef struct {
// 		Typ string
// 		Msg string
// 	}
// 	ErrQul struct {
// 		Need string
// 		Real string
// 	}
// )

// func (e *ErrDef) Error() string {
// 	return ztool.Str_FastConcat(e.Typ, `: `, e.Msg)
// }
// func (e *ErrQul) Error() string {
// 	return ztool.Str_FastConcat(`实际音质不匹配: `, e.Need, ` <= `, e.Real)
// }

// 验证失败(Verify Failed)
// func ErrVerify(msg string) error {
// 	return &ErrDef{
// 		Typ: Err_Verify,
// 		Msg: msg,
// 	}
// }

// 实际音质不匹配
// func ErrQuality(need, real string) error {
// 	return &ErrQul{
// 		Need: need,
// 		Real: real,
// 	}
// }

// 无返回数据(No Data)
// func ErrNoData() error {
// 	return &ErrDef{
// 		Typ: `No Data`,
// 		Msg: ``,
// 	}
// }
