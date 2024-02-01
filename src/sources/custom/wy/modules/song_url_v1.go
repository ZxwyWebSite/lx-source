package wy

import (
	"net/http"

	"github.com/ZxwyWebSite/ztool"
)

// 歌曲链接 - v1
// 此版本不再采用 br 作为音质区分的标准
// 而是采用 standard, exhigh, lossless, hires, jyeffect(高清环绕声), sky(沉浸环绕声), jymaster(超清母带) 进行音质判断
func SongUrlV1(query ReqQuery) (*ReqAnswer, error) {
	if query.Cookie == nil {
		query.Cookie = make(map[string]string)
	}
	query.Cookie[`os`] = `android`
	query.Cookie[`appver`] = `8.10.05`
	data := map[string]any{
		`ids`:        ztool.Str_FastConcat(`[`, query.Ids, `]`),
		`level`:      query.Level,
		`encodeType`: `flac`,
	}
	if query.Level == `sky` /*|| query.Level == `jysky`*/ {
		data[`immerseType`] = `c51`
	}
	return createRequest(
		http.MethodPost,
		`https://interface.music.163.com/eapi/song/enhance/player/url/v1`,
		data,
		reqOptions{
			Crypto: `eapi`,
			Cookie: query.Cookie,
			RealIP: query.RealIP,
			Url:    `/api/song/enhance/player/url/v1`,
		},
	)
}
