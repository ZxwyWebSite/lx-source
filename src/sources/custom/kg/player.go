package kg

import (
	"lx-source/src/env"
	"lx-source/src/sources"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func Url(songMid, quality string) (ourl, msg string) {
	loger := env.Loger.NewGroup(`Kg`)
	defer loger.Free()
	rquality, ok := qualityMap[quality]
	if !ok {
		msg = sources.E_QNotSupport
		return
	}
	info, emsg := getMusicInfo(strings.ToLower(songMid))
	if emsg != `` {
		loger.Error(`GetInfo: %v`, emsg)
		msg = emsg
		return
	}
	loger.Debug(`Info: %+v`, info)
	var tHash string
	switch quality {
	case sources.Q_128k:
		tHash = info.AudioInfo.Hash128
	case sources.Q_320k:
		tHash = info.AudioInfo.Hash320
	case sources.Q_flac:
		tHash = info.AudioInfo.HashFlac
	case sources.Q_fl24:
		tHash = info.AudioInfo.HashHigh
	}
	if tHash == `` {
		msg = sources.E_QNotMatch
		return
	}
	tHash = strings.ToLower(tHash)
	now := time.Now()
	params := map[string]string{
		`album_id`:       info.AlbumInfo.AlbumID,
		`userid`:         env.Config.Custom.Kg_userId,
		`area_code`:      `1`,
		`hash`:           tHash,
		`module`:         ``,
		`mid`:            mid,
		`appid`:          env.Config.Custom.Kg_Client_AppId,
		`ssa_flag`:       `is_fromtrack`,
		`clientver`:      env.Config.Custom.Kg_Client_Version,
		`open_time`:      now.Format(`20060102`),
		`vipType`:        `6`,
		`ptype`:          `0`,
		`token`:          env.Config.Custom.Kg_token,
		`auth`:           ``,
		`mtype`:          `0`,
		`album_audio_id`: info.AlbumAudioID,
		`behavior`:       `play`,
		`clienttime`:     strconv.FormatInt(now.Unix(), 10),
		`pid`:            `2`,
		`key`:            getKey(tHash),
		`dfid`:           `-`,
		`pidversion`:     `3001`,

		`quality`: rquality,
		// `IsFreePart`: `1`,
	}
	if !env.Config.Custom.Kg_Enable {
		params[`IsFreePart`] = `1` // 仅游客登录时允许获取试听
	}
	headers := map[string]string{
		`User-Agent`: `Android712-AndroidPhone-8983-18-0-NetMusic-wifi`,
		`KG-THash`:   `3e5ec6b`,
		`KG-Rec`:     `1`,
		`KG-RC`:      `1`,

		`x-router`: `tracker.kugou.com`,
	}
	var resp playInfo
	err := signRequest(http.MethodGet, url, nil, params, headers, &resp)
	if err != nil {
		loger.Error(`Request: %s`, err)
		msg = sources.ErrHttpReq
		return
	}
	loger.Debug(`Resp: %+v`, resp)
	switch resp.Status {
	case 3:
		msg = `该歌曲在酷狗没有版权，请换源播放`
	case 2:
		msg = `链接获取失败：请检查账号是否有会员或数字专辑是否购买`
	}
	if resp.Status != 1 {
		if msg == `` {
			msg = `链接获取失败，可能是数字专辑或者api失效，Status: ` + strconv.Itoa(resp.Status)
		}
		return
	}
	ourl = resp.URL[len(resp.URL)-1]
	return
}
