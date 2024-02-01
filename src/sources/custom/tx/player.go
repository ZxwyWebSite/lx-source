package tx

import (
	"lx-source/src/env"
	"lx-source/src/sources"

	"github.com/ZxwyWebSite/ztool"
	"github.com/ZxwyWebSite/ztool/x/bytesconv"
)

type playInfo struct {
	Code int `json:"code"`
	Data struct {
		Uin          string   `json:"uin"`
		Retcode      int      `json:"retcode"`
		VerifyType   int      `json:"verify_type"`
		LoginKey     string   `json:"login_key"`
		Msg          string   `json:"msg"`
		Sip          []string `json:"sip"`
		Thirdip      []string `json:"thirdip"`
		Testfile2G   string   `json:"testfile2g"`
		Testfilewifi string   `json:"testfilewifi"`
		Midurlinfo   []struct {
			Songmid           string `json:"songmid"`
			Filename          string `json:"filename"`
			Purl              string `json:"purl"`
			Errtype           string `json:"errtype"`
			P2Pfromtag        int    `json:"p2pfromtag"`
			Qmdlfromtag       int    `json:"qmdlfromtag"`
			CommonDownfromtag int    `json:"common_downfromtag"`
			VipDownfromtag    int    `json:"vip_downfromtag"`
			Pdl               int    `json:"pdl"`
			Premain           int    `json:"premain"`
			Hisdown           int    `json:"hisdown"`
			Hisbuy            int    `json:"hisbuy"`
			UIAlert           int    `json:"uiAlert"`
			Isbuy             int    `json:"isbuy"`
			Pneedbuy          int    `json:"pneedbuy"`
			Pneed             int    `json:"pneed"`
			Isonly            int    `json:"isonly"`
			Onecan            int    `json:"onecan"`
			Result            int    `json:"result"`
			Tips              string `json:"tips"`
			Opi48Kurl         string `json:"opi48kurl"`
			Opi96Kurl         string `json:"opi96kurl"`
			Opi192Kurl        string `json:"opi192kurl"`
			Opiflackurl       string `json:"opiflackurl"`
			Opi128Kurl        string `json:"opi128kurl"`
			Opi192Koggurl     string `json:"opi192koggurl"`
			Wififromtag       string `json:"wififromtag"`
			Flowfromtag       string `json:"flowfromtag"`
			Wifiurl           string `json:"wifiurl"`
			Flowurl           string `json:"flowurl"`
			Vkey              string `json:"vkey"`
			Opi30Surl         string `json:"opi30surl"`
			Ekey              string `json:"ekey"`
			AuthSwitch        int    `json:"auth_switch"`
			Subcode           int    `json:"subcode"`
			Opi96Koggurl      string `json:"opi96koggurl"`
			AuthSwitch2       int    `json:"auth_switch2"`
		} `json:"midurlinfo"`
		Servercheck string `json:"servercheck"`
		Expiration  int    `json:"expiration"`
	} `json:"data"`
}

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
	loger.Debug(`infoBody: %+v`, infoBody)
	if emsg != `` {
		msg = emsg
		return
	}
	var uauthst, uuin string = env.Config.Custom.Tx_Ukey, env.Config.Custom.Tx_Uuin
	if uuin == `` || !env.Config.Custom.Tx_Enable {
		uuin = `1535153710`
	}
	var strFileName string
	tryLink := infoBody.TrackInfo.Pay.PayPlay == 1 && /*uauthst == ``&&*/ !env.Config.Custom.Tx_Enable
	if tryLink {
		strFileName = ztool.Str_FastConcat(`RS02`, infoBody.TrackInfo.Vs[0], `.mp3`)
	} else {
		strFileName = ztool.Str_FastConcat(infoFile.H, infoBody.TrackInfo.File.MediaMid, infoFile.E)
	}
	requestBody := ztool.Str_FastConcat(`{"comm":{"authst":"`, uauthst, `","ct":"26","cv":"2010101","qq":"`, uuin, `","v":"2010101"},"req_0":{"method":"CgiGetVkey","module":"vkey.GetVkeyServer","param":{"filename":["`, strFileName, `"],"guid":"114514","loginflag":1,"platform":"20","songmid":["`, songMid, `"],"songtype":[0],"uin":"10086"}}}`)
	var infoResp struct {
		Code int `json:"code"`
		// Ts      int64    `json:"ts"`
		// StartTs int64    `json:"start_ts"`
		// Traceid string   `json:"traceid"`
		Req0 playInfo `json:"req_0"`
	}
	err := signRequest(bytesconv.StringToBytes(requestBody), &infoResp)
	if err != nil {
		msg = err.Error()
		return
	}
	loger.Debug(`infoResp: %+v`, infoResp)
	infoData := infoResp.Req0.Data.Midurlinfo[0]
	if infoData.Purl == `` {
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
		return
	}
	ourl = ztool.Str_FastConcat(`https://ws.stream.qqmusic.qq.com/`, infoData.Purl)
	return
}
