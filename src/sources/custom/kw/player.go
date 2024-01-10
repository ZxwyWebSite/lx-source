package kw

import (
	"io"
	"lx-source/src/env"
	"lx-source/src/sources"
	"net/http"
	"strings"

	"github.com/ZxwyWebSite/ztool"
)

func Url(songMid, quality string) (ourl, msg string) {
	loger := env.Loger.NewGroup(`Kw`)
	infoFile, ok := fileInfo[quality]
	if !ok {
		msg = sources.E_QNotSupport
		return
	}
	target_url := ztool.Str_FastConcat(
		`https://mobi.kuwo.cn/mobi.s?f=kuwo&q=`,
		Base64_encrypt(ztool.Str_FastConcat(
			`user=0&android_id=0&prod=kwplayer_ar_8.5.5.0&corp=kuwo&newver=3&vipver=8.5.5.0&source=kwplayer_ar_8.5.5.0_apk_keluze.apk&p2p=1&notrace=0&type=convert_url2`,
			`&br=`, infoFile.H, infoFile.E,
			`&format=`, infoFile.E,
			`&rid=`, songMid,
			`&priority=bitrate&loginUid=0&network=WIFI&loginSid=0&mode=down`,
		)),
	)
	ztool.Net_Request(http.MethodGet, target_url, nil,
		[]ztool.Net_ReqHandlerFunc{
			ztool.Net_ReqAddHeader(header),
		},
		[]ztool.Net_ResHandlerFunc{
			func(res *http.Response) (err error) {
				data, err := io.ReadAll(res.Body)
				if err != nil {
					msg = err.Error()
					return
				}
				if res.StatusCode != http.StatusOK {
					msg = res.Status
					loger.Debug(`failed: %v`, res.Status)
					return
				}
				infoData := mkMap(data)
				loger.Debug(`infoData: %+v`, infoData)
				realQuality := qualityMapReverse[infoData[`bitrate`]]
				if realQuality != quality {
					msg = sources.E_QNotMatch
					return
				}
				ourl = infoData[`url`][:strings.Index(infoData[`url`], `?`)]
				return
			},
		},
	)
	return
}
