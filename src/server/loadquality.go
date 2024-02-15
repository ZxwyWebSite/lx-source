package server

import (
	"lx-source/src/env"
	"lx-source/src/sources"
)

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
			m[sources.I_wy] = defQuality
		} else {
			m[sources.I_wy] = tstQuality
		}
	}
	// 1.mg
	if env.Config.Source.Enable_Mg {
		if env.Config.Custom.Mg_Enable {
			m[sources.I_mg] = defQuality
		} else {
			m[sources.I_mg] = tstQuality
		}
	}
	// 2.kw
	if env.Config.Source.Enable_Kw {
		if env.Config.Custom.Kw_Enable {
			m[sources.I_kw] = stdQuality
		}
	}
	// 3.kg
	if env.Config.Source.Enable_Kg {
		if env.Config.Custom.Kg_Enable {
			m[sources.I_kg] = defQuality
		} else {
			m[sources.I_kg] = tstQuality
		}
	}
	// 4.tx
	if env.Config.Source.Enable_Tx {
		if env.Config.Custom.Tx_Enable {
			m[sources.I_tx] = stdQuality
		} else {
			m[sources.I_tx] = tstQuality
		}
	}
	// 5.lx
	if env.Config.Source.Enable_Lx {
		m[sources.I_lx] = defQuality
	}
	return m
}
