package wy

import (
	"net/http"
	"strings"

	"github.com/ZxwyWebSite/ztool"
)

type song_url_query struct {
	Cookie map[string]string
	Ids    string
	Br     string
	RealIP string
}

func SongUrl(query song_url_query) (*reqAnswer, error) {
	query.Cookie[`os`] = `pc`
	if query.Br == `` {
		query.Br = `999000`
	}
	ids := strings.Split(query.Ids, `,`)
	// idj, err := json.Marshal(ids)
	// if err != nil {
	// 	return nil, err
	// }
	data := map[string]any{
		`ids`: ztool.Str_FastConcat(`["`, strings.Join(ids, `","`), `"]`), //bytesconv.BytesToString(idj), //`["1998644237"]`,
		`br`:  query.Br,                                                   //ztool.Str_Select(query.Br, `999000`),
	}
	return createRequest(
		http.MethodPost,
		`https://interface3.music.163.com/eapi/song/enhance/player/url`,
		data,
		reqOptions{
			Crypto: `eapi`,
			Cookie: query.Cookie,
			RealIP: query.RealIP,
			Url:    `/api/song/enhance/player/url`,
		},
	)
}
