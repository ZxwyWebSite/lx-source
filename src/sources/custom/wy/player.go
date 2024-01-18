package wy

import (
	"io"
	"lx-source/src/env"
	"lx-source/src/sources"
	"lx-source/src/sources/builtin"
	"lx-source/src/sources/custom/utils"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/ZxwyWebSite/ztool"
	"github.com/ZxwyWebSite/ztool/x/cookie"
)

func Url(songMid, quality string) (ourl, msg string) {
	loger := env.Loger.NewGroup(`Wy`)
	rquality, ok := brMap[quality]
	if !ok {
		msg = sources.E_QNotSupport
		return
	}
	cookies := cookie.Parse(env.Config.Custom.Wy_Cookie)
	answer, err := SongUrl(song_url_query{
		Cookie: cookie.ToMap(cookies),
		Ids:    songMid,
		Br:     rquality,
	})
	var body builtin.WyApi_Song
	if err == nil {
		err = ztool.Val_MapToStruct(answer.Body, &body)
	}
	if err != nil {
		loger.Error(`SongUrl: %s`, err)
		msg = sources.ErrHttpReq
		return
	}
	loger.Debug(`Resp: %+v`, body)
	if len(body.Data) == 0 {
		msg = `No Data：无返回数据`
		return
	}
	data := body.Data[0]
	br := strconv.Itoa(data.Br) // 注：由于flac返回br值不固定，暂无法进行比较
	if br != rquality && !ztool.Chk_IsMatch(br, sources.Q_flac, sources.Q_fl24) {
		msg = sources.E_QNotMatch
		return
	}
	ourl = utils.DelQuery(data.URL)
	return
}

func PyUrl(songMid, quality string) (ourl, msg string) {
	loger := env.Loger.NewGroup(`Wy`)
	rquality, ok := qualityMap[quality]
	if !ok {
		msg = sources.E_QNotSupport
		return
	}
	path := `/api/song/enhance/player/url/v1`
	requestUrl := `https://interface.music.163.com/eapi/song/enhance/player/url/v1`
	var body builtin.WyApi_Song
	text := ztool.Str_FastConcat(
		`{"encodeType":"flac","ids":["`, songMid, `"],"level":"`, rquality, `"}`,
	)
	var form url.Values = eapiEncrypt(path, text)
	// form, err := json.Marshal(eapiEncrypt(path, text))
	// if err == nil {
	err := ztool.Net_Request(
		http.MethodPost, requestUrl,
		strings.NewReader(form.Encode()), //bytes.NewReader(form),
		[]ztool.Net_ReqHandlerFunc{ztool.Net_ReqAddHeader(map[string]string{
			`Cookie`: env.Config.Custom.Wy_Cookie,
		})},
		[]ztool.Net_ResHandlerFunc{
			func(res *http.Response) error {
				body, err := io.ReadAll(res.Body)
				if err != nil {
					return err
				}
				loger.Info(`%s`, body)
				return ztool.Err_EsContinue
			},
			ztool.Net_ResToStruct(&body),
		},
	)
	// }
	if err != nil {
		loger.Error(`Request: %s`, err)
		msg = sources.ErrHttpReq
		return
	}
	loger.Debug(`Resp: %+v`, body)
	if len(body.Data) == 0 {
		msg = `No Data：无返回数据`
		return
	}
	data := body.Data[0]
	if data.Level != rquality {
		msg = sources.E_QNotMatch
		return
	}
	ourl = utils.DelQuery(data.URL)
	return
}
