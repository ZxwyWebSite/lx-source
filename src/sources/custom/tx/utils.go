package tx

import (
	"bytes"
	"lx-source/src/sources"
	"net/http"

	"github.com/ZxwyWebSite/ztool"
)

var (
	fileInfo = map[string]struct {
		E string
		H string
	}{
		sources.Q_128k: {
			E: `.mp3`,
			H: `M500`,
		},
		sources.Q_320k: {
			E: `.mp3`,
			H: `M800`,
		},
		sources.Q_flac: {
			E: `.flac`,
			H: `F000`,
		},
		sources.Q_fl24: {
			E: `.flac`,
			H: `RS01`,
		},
		`dolby`: {
			E: `.flac`,
			H: `Q000`,
		},
		`master`: {
			E: `.flac`,
			H: `AI00`,
		},
	}
	qualityMapReverse = map[string]string{
		`M500`: sources.Q_128k,
		`M800`: sources.Q_320k,
		`F000`: sources.Q_flac,
		`RS01`: sources.Q_fl24,
		`Q000`: `dolby`,
		`AI00`: `master`,
	}
)

func signRequest(data []byte, out any) error {
	s := sign(data)
	return ztool.Net_Request(http.MethodPost,
		ztool.Str_FastConcat(`https://u.y.qq.com/cgi-bin/musics.fcg?format=json&sign=`, s),
		bytes.NewReader(data),
		[]ztool.Net_ReqHandlerFunc{
			ztool.Net_ReqAddHeaders(map[string]string{
				`Referer`: `https://y.qq.com/`,
			}),
		},
		[]ztool.Net_ResHandlerFunc{
			ztool.Net_ResToStruct(out),
		},
	)
}
