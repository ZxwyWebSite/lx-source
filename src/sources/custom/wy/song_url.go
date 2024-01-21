package wy

import (
	"net/http"
	"sort"
	"strconv"
	"strings"

	"github.com/ZxwyWebSite/ztool"
)

// type Query_song_url struct {
// 	Cookie map[string]string
// 	Ids    string
// 	Br     string
// 	RealIP string
// }

// 歌曲链接
func SongUrl(query ReqQuery) (*ReqAnswer, error) {
	if query.Cookie == nil {
		query.Cookie = make(map[string]string)
	}
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
	res, err := createRequest(
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
	// 根据id排序
	if length := len(ids); length > 1 && err == nil {
		indexOf := make(map[string]int, length)
		for i := 0; i < length; i++ {
			indexOf[ids[i]] = i
		}
		if data, ok := res.Body[`data`].([]interface{}); ok {
			sort.SliceStable(data, func(a, b int) bool {
				da, oa := data[a].(map[string]interface{})
				db, ob := data[b].(map[string]interface{})
				if oa && ob {
					ia, ka := da[`id`].(float64)
					ib, kb := db[`id`].(float64)
					if ka && kb {
						ta := strconv.FormatInt(int64(ia), 10)
						tb := strconv.FormatInt(int64(ib), 10)
						return indexOf[ta] < indexOf[tb]
					}
				}
				return false
			})
			res.Body[`data`] = data
		}
	}
	return res, err
}
