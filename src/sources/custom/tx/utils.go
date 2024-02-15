package tx

import (
	"bytes"
	"lx-source/src/sources"
	"net/http"

	"github.com/ZxwyWebSite/ztool"
)

var (
	fileInfo = map[string]struct {
		E string // 扩展名
		H string // 专用音质
	}{
		sources.Q_128k: {
			E: sources.X_mp3,
			H: `M500`,
		},
		sources.Q_320k: {
			E: sources.X_mp3,
			H: `M800`,
		},
		sources.Q_flac: {
			E: sources.Q_flac,
			H: `F000`,
		},
		sources.Q_fl24: {
			E: sources.Q_flac,
			H: `RS01`,
		},
		`dolby`: {
			E: sources.Q_flac,
			H: `Q000`,
		},
		`master`: {
			E: sources.Q_flac,
			H: `AI00`, // (~~母带音质大部分都是AI提上去的~~)
		},
	}
	// qualityMapReverse = map[string]string{
	// 	`M500`: sources.Q_128k,
	// 	`M800`: sources.Q_320k,
	// 	`F000`: sources.Q_flac,
	// 	`RS01`: sources.Q_fl24,
	// 	`Q000`: `dolby`,
	// 	`AI00`: `master`,
	// }
	header = map[string]string{
		`Referer`: `https://y.qq.com/`,
	}
)

func signRequest(data []byte, out any) error {
	s := sign(data)
	return ztool.Net_Request(http.MethodPost,
		ztool.Str_FastConcat(`https://u.y.qq.com/cgi-bin/musics.fcg?format=json&sign=`, s),
		bytes.NewReader(data),
		[]ztool.Net_ReqHandlerFunc{
			ztool.Net_ReqAddHeaders(header),
		},
		[]ztool.Net_ResHandlerFunc{
			ztool.Net_ResToStruct(out),
		},
	)
}
