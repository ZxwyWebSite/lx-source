package tx

import (
	"lx-source/src/env"
	"lx-source/src/sources"

	"github.com/ZxwyWebSite/ztool"
	"github.com/ZxwyWebSite/ztool/x/bytesconv"
)

/*
 音乐URL获取逻辑：
  if需要付费播放and无账号信息:
   试听获取
  el有账号信息or无需付费:
   正常获取
    if没有链接:
     尝试获取试听
      if还没有链接:
       报错
   返回结果
 注：
  以上逻辑暂时没想好怎么改，
  当前根据配置文件 [Custom].Tx_Enable 是否开启判断，
   if需要付费播放and未开启账号解析:
    试听获取
   el无需付费or有账号信息:
    正常获取
     if没有链接:
      报错
   返回结果
 更新：
  可通过 goto loop 实现，但可能会导致逻辑混乱 (想使用账号获取正常链接却返回试听链接)
 2024-03-16:
  正常获取->128k获取->试听获取
*/

func Url(songMid, quality string) (ourl, msg string) {
	loger := env.Loger.NewGroup(`Tx`)
	defer loger.Free()
	infoFile, ok := fileInfo[quality]
	if !ok || (!env.Config.Custom.Tx_Enable && quality != sources.Q_128k) {
		msg = sources.E_QNotSupport //`不支持的音质`
		return
	}
	infoBody, emsg := getMusicInfo(songMid)
	if emsg != `` {
		loger.Error(`GetInfo: %v`, emsg)
		msg = emsg
		return
	}
	loger.Debug(`infoBody: %+v`, infoBody)
	var uauthst, uuin string = env.Config.Custom.Tx_Ukey, env.Config.Custom.Tx_Uuin
	if uuin == `` || !env.Config.Custom.Tx_Enable {
		uuin = `1535153710`
	}
	var strFileName string
	tryLink := infoBody.TrackInfo.Pay.PayPlay == 1 && /*uauthst == ``&&*/ !env.Config.Custom.Tx_Enable
Loop:
	if tryLink {
		if infoBody.TrackInfo.Vs[0] == `` {
			msg = sources.ErrNoLink
			return
		}
		strFileName = ztool.Str_FastConcat(`RS02`, infoBody.TrackInfo.Vs[0], `.`, sources.X_mp3)
	} else {
		strFileName = ztool.Str_FastConcat(infoFile.H, infoBody.TrackInfo.File.MediaMid, `.`, infoFile.E)
	}
	requestBody := ztool.Str_FastConcat(
		`{"comm":{"authst":"`,
		uauthst,
		`","ct":"26","cv":"2010101","qq":"`,
		uuin,
		`","v":"2010101"},"req_0":{"method":"CgiGetVkey","module":"vkey.GetVkeyServer","param":{"filename":["`,
		strFileName,
		`"],"guid":"20211008","loginflag":1,"platform":"20","songmid":["`,
		songMid,
		`"],"songtype":[0],"uin":"`,
		uuin,
		`"}}}`,
	)
	var infoResp struct {
		Code int `json:"code"`
		// Ts      int64    `json:"ts"`
		// StartTs int64    `json:"start_ts"`
		// Traceid string   `json:"traceid"`
		Req0 playInfo `json:"req_0"`
	}
	err := signRequest(bytesconv.StringToBytes(requestBody), &infoResp)
	if err != nil {
		loger.Error(`Request: %s`, err)
		msg = err.Error()
		return
	}
	loger.Debug(`infoResp: %+v`, infoResp)
	if len(infoResp.Req0.Data.Midurlinfo) == 0 {
		msg = `No Data: 无返回数据`
		return
	}
	infoData := infoResp.Req0.Data.Midurlinfo[0]
	if infoData.Purl == `` {
		if env.Config.Source.ForceFallback && !tryLink {
			if quality != sources.Q_128k && infoBody.TrackInfo.Pay.PayPlay == 0 {
				msg = `Fallback to 128k`
				infoFile = fileInfo[sources.Q_128k]
				quality = sources.Q_128k
			} else {
				msg = `Fallbacked`
				tryLink = true
			}
			goto Loop
		}
		msg = sources.E_NoLink //`无法获取音乐链接`
		return
	}
	realQuality := ztool.Str_Before(infoData.Filename, `.`)[:4] //strings.Split(infoData.Filename, `.`)[0][:4]
	// if qualityMapReverse[realQuality] != quality && /*infoBody.TrackInfo.Pay.PayPlay == 0*/ !tryLink {
	// 	msg = sources.E_QNotMatch //`实际音质不匹配`
	// 	return
	// }
	if realQuality != infoFile.H && !tryLink {
		msg = sources.E_QNotMatch
		if !env.Config.Source.ForceFallback {
			return
		}
	}
	ourl = env.Config.Custom.Tx_CDNUrl + infoData.Purl
	return
}
