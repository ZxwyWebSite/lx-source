package server

import "lx-source/src/env"

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
	if env.Config.Source.Enable_Wy {
		if env.Config.Custom.Wy_Enable {
			m[0] = defQuality
		} else {
			m[0] = tstQuality
		}
	}
	// 1.mg
	if env.Config.Source.Enable_Mg {
		if env.Config.Custom.Mg_Enable {
			m[1] = defQuality
		}
	}
	// 2.kw
	if env.Config.Source.Enable_Kw {
		if env.Config.Custom.Kw_Enable {
			m[2] = stdQuality
		}
	}
	// 3.kg
	if env.Config.Source.Enable_Kg {
		if env.Config.Custom.Kg_Enable {
			m[3] = tstQuality
		}
	}
	// 4.tx
	if env.Config.Source.Enable_Tx {
		if env.Config.Custom.Tx_Enable {
			m[4] = stdQuality
		} else {
			m[4] = tstQuality
		}
	}
	// 5.lx
	// if env.Config.Source.Enable_Lx {
	// 	m[5] = defQuality
	// }
	return m
}
