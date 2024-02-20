// 账号解析源
package custom

import (
	"lx-source/src/env"
	"lx-source/src/sources/custom/kg"
	"lx-source/src/sources/custom/kw"
	"lx-source/src/sources/custom/mg"
	"lx-source/src/sources/custom/tx"
	"lx-source/src/sources/custom/wy"
	"strconv"
)

type (
	// 源定义
	UrlFunc func(string, string) (string, string)
	LrcFunc func(string) (string, string)
	PicFunc func(string) (string, string)
	VefFunc func(string) bool
	// 源接口
	Source interface {
		Url(string, string) (string, string) // 外链
		Lrc(string) (string, string)         // 歌词
		Pic(string) (string, string)         // 封面
		Vef(string) bool                     // 验证
	}
)

func notSupport(string) (string, string) { return ``, `不支持的平台或功能` }

// 接口封装
type WrapSource struct {
	UrlFunc
	LrcFunc
	PicFunc
	VefFunc
}

func (ws *WrapSource) Url(songMid, quality string) (string, string) {
	return ws.UrlFunc(songMid, quality)
}
func (ws *WrapSource) Lrc(songMid string) (string, string) {
	return ws.LrcFunc(songMid)
}
func (ws *WrapSource) Pic(songMid string) (string, string) {
	return ws.PicFunc(songMid)
}
func (ws *WrapSource) Vef(songMid string) bool {
	return ws.VefFunc(songMid)
}

var (
	WySource Source
	MgSource Source
	KwSource Source
	KgSource Source
	TxSource Source
	LxSource Source
)

func init() {
	env.Inits.Add(func() {
		if env.Config.Source.Enable_Wy {
			WySource = &WrapSource{
				UrlFunc: wy.Url,
				LrcFunc: notSupport,
				PicFunc: notSupport,
				VefFunc: func(songMid string) bool {
					_, err := strconv.ParseUint(songMid, 10, 0)
					return err == nil
				},
			}
		}
		if env.Config.Source.Enable_Mg {
			MgSource = &WrapSource{
				UrlFunc: mg.Url,
				LrcFunc: notSupport,
				PicFunc: notSupport,
				VefFunc: func(songMid string) bool { return len(songMid) == 11 },
			}
		}
		if env.Config.Source.Enable_Kw {
			KwSource = &WrapSource{
				UrlFunc: kw.Url,
				LrcFunc: notSupport,
				PicFunc: notSupport,
				VefFunc: func(songMid string) bool {
					_, err := strconv.ParseUint(songMid, 10, 0)
					return err == nil
				},
			}
		}
		if env.Config.Source.Enable_Kg {
			KgSource = &WrapSource{
				UrlFunc: kg.Url,
				LrcFunc: notSupport,
				PicFunc: notSupport,
				VefFunc: func(songMid string) bool { return len(songMid) == 32 },
			}
		}
		if env.Config.Source.Enable_Tx {
			TxSource = &WrapSource{
				UrlFunc: tx.Url,
				LrcFunc: notSupport,
				PicFunc: notSupport,
				VefFunc: func(songMid string) bool { return len(songMid) == 14 },
			}
		}
		if env.Config.Source.Enable_Lx {
			LxSource = nil
		}
	})
}
