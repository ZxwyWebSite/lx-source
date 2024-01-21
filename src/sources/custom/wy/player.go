package wy

import (
	"lx-source/src/env"
	"lx-source/src/sources"
	"lx-source/src/sources/custom/utils"

	"github.com/ZxwyWebSite/ztool"
	"github.com/ZxwyWebSite/ztool/x/cookie"
)

type playInfo struct {
	Data []struct {
		ID            int         `json:"id"`
		URL           string      `json:"url"`
		Br            int         `json:"br"`
		Size          int         `json:"size"`
		Md5           string      `json:"md5"`
		Code          int         `json:"code"`
		Expi          int         `json:"expi"`
		Type          string      `json:"type"`
		Gain          float64     `json:"gain"`
		Peak          float64     `json:"peak"`
		Fee           int         `json:"fee"`
		Uf            interface{} `json:"uf"`
		Payed         int         `json:"payed"`
		Flag          int         `json:"flag"`
		CanExtend     bool        `json:"canExtend"`
		FreeTrialInfo struct {
			AlgData      interface{} `json:"algData"`
			End          int         `json:"end"`
			FragmentType int         `json:"fragmentType"`
			Start        int         `json:"start"`
		} `json:"freeTrialInfo"`
		Level              string `json:"level"`
		EncodeType         string `json:"encodeType"`
		FreeTrialPrivilege struct {
			ResConsumable      bool        `json:"resConsumable"`
			UserConsumable     bool        `json:"userConsumable"`
			ListenType         int         `json:"listenType"`
			CannotListenReason int         `json:"cannotListenReason"`
			PlayReason         interface{} `json:"playReason"`
		} `json:"freeTrialPrivilege"`
		FreeTimeTrialPrivilege struct {
			ResConsumable  bool `json:"resConsumable"`
			UserConsumable bool `json:"userConsumable"`
			Type           int  `json:"type"`
			RemainTime     int  `json:"remainTime"`
		} `json:"freeTimeTrialPrivilege"`
		URLSource   int         `json:"urlSource"`
		RightSource int         `json:"rightSource"`
		PodcastCtrp interface{} `json:"podcastCtrp"`
		EffectTypes interface{} `json:"effectTypes"`
		Time        int         `json:"time"`
	} `json:"data"`
	Code int `json:"code"`
}

func Url(songMid, quality string) (ourl, msg string) {
	loger := env.Loger.NewGroup(`Wy`)
	rquality, ok := qualityMap[quality]
	if !ok {
		msg = sources.E_QNotSupport
		return
	}
	cookies := cookie.Parse(env.Config.Custom.Wy_Cookie)
	answer, err := SongUrlV1(ReqQuery{
		Cookie: cookie.ToMap(cookies),
		Ids:    songMid,
		// Br:     rquality,
		Level: rquality,
	})
	var body playInfo
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
	if data.Level != rquality {
		msg = ztool.Str_FastConcat(`实际音质不匹配: `, rquality, ` <= `, data.Level)
		return
	}
	// br := strconv.Itoa(data.Br) // 注：由于flac返回br值不固定，暂无法进行比较
	// if br != rquality && !ztool.Chk_IsMatch(br, sources.Q_flac, sources.Q_fl24) {
	// 	msg = sources.E_QNotMatch
	// 	return
	// }
	ourl = utils.DelQuery(data.URL)
	return
}

// func PyUrl(songMid, quality string) (ourl, msg string) {
// 	loger := env.Loger.NewGroup(`Wy`)
// 	rquality, ok := qualityMap[quality]
// 	if !ok {
// 		msg = sources.E_QNotSupport
// 		return
// 	}
// 	path := `/api/song/enhance/player/url/v1`
// 	requestUrl := `https://interface.music.163.com/eapi/song/enhance/player/url/v1`
// 	var body builtin.WyApi_Song
// 	text := ztool.Str_FastConcat(
// 		`{"encodeType":"flac","ids":["`, songMid, `"],"level":"`, rquality, `"}`,
// 	)
// 	var form url.Values = eapiEncrypt(path, text)
// 	// form, err := json.Marshal(eapiEncrypt(path, text))
// 	// if err == nil {
// 	err := ztool.Net_Request(
// 		http.MethodPost, requestUrl,
// 		strings.NewReader(form.Encode()), //bytes.NewReader(form),
// 		[]ztool.Net_ReqHandlerFunc{ztool.Net_ReqAddHeader(map[string]string{
// 			`Cookie`: env.Config.Custom.Wy_Cookie,
// 		})},
// 		[]ztool.Net_ResHandlerFunc{
// 			func(res *http.Response) error {
// 				body, err := io.ReadAll(res.Body)
// 				if err != nil {
// 					return err
// 				}
// 				loger.Info(`%s`, body)
// 				return ztool.Err_EsContinue
// 			},
// 			ztool.Net_ResToStruct(&body),
// 		},
// 	)
// 	// }
// 	if err != nil {
// 		loger.Error(`Request: %s`, err)
// 		msg = sources.ErrHttpReq
// 		return
// 	}
// 	loger.Debug(`Resp: %+v`, body)
// 	if len(body.Data) == 0 {
// 		msg = `No Data：无返回数据`
// 		return
// 	}
// 	data := body.Data[0]
// 	if data.Level != rquality {
// 		msg = sources.E_QNotMatch
// 		return
// 	}
// 	ourl = utils.DelQuery(data.URL)
// 	return
// }
