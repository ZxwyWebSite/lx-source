package kw

import (
	"bytes"
	"lx-source/src/sources"

	"github.com/ZxwyWebSite/ztool/x/bytesconv"
)

var (
	fileInfo = map[string]struct {
		E string // 扩展名
		H string // 专用音质
	}{
		sources.Q_128k: {
			E: sources.X_mp3,
			H: sources.Q_128k,
		},
		sources.Q_320k: {
			E: sources.X_mp3,
			H: sources.Q_320k,
		},
		sources.Q_flac: {
			E: sources.Q_flac,
			H: `2000k`,
		},
		sources.Q_fl24: {
			E: sources.Q_flac,
			H: `4000k`,
		},
	}
	// 注：这个还是有规律的，加上或去掉k即可直接比较
	// qualityMapReverse = map[string]string{
	// 	`128`:  sources.Q_128k,
	// 	`320`:  sources.Q_320k,
	// 	`2000`: sources.Q_flac,
	// 	`4000`: sources.Q_fl24,
	// }
	desheader = map[string]string{
		// `User-Agent`: `okhttp/3.10.0`,
	}
	bdheader = map[string]string{
		`channel`: `qq`,
		`plat`:    `ar`,
		`net`:     `wifi`,
		`ver`:     `3.1.2`,
		// `uid`:     ``,
		// `devId`:   `0`,
	}
)

func mkMap(data []byte) map[string]string {
	out := make(map[string]string)
	sep := bytes.Split(data, []byte{13, 10})
	for i, r := 0, len(sep); i < r; i++ {
		pat := bytes.Split(sep[i], []byte{61})
		if len(pat) == 2 {
			out[bytesconv.BytesToString(pat[0])] = bytesconv.BytesToString(pat[1])
			continue
		}
		out[`_`] += bytesconv.BytesToString(pat[0]) + `;`
	}
	return out
}
