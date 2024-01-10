package kw

import (
	"bytes"
	"lx-source/src/sources"

	"github.com/ZxwyWebSite/ztool/x/bytesconv"
)

var (
	fileInfo = map[string]struct {
		E string
		H string
	}{
		sources.Q_128k: {
			E: `mp3`,
			H: sources.Q_128k,
		},
		sources.Q_320k: {
			E: `mp3`,
			H: sources.Q_320k,
		},
		sources.Q_flac: {
			E: sources.Q_flac,
			H: `2000k`,
		},
	}
	qualityMapReverse = map[string]string{
		`128`:  sources.Q_128k,
		`320`:  sources.Q_320k,
		`2000`: sources.Q_flac,
	}
	header = map[string]string{
		`User-Agent`: `okhttp/3.10.0`,
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
		out[`_etc`] += bytesconv.BytesToString(pat[0]) + `; `
	}
	return out
}
