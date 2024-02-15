package tx

import (
	"github.com/ZxwyWebSite/ztool"
	"github.com/ZxwyWebSite/ztool/x/bytesconv"
)

func getMusicInfo(songMid string) (infoBody musicInfo, emsg string) {
	infoReqBody := ztool.Str_FastConcat(`{"comm":{"ct":"19","cv":"1859","uin":"0"},"req":{"method":"get_song_detail_yqq","module":"music.pf_song_detail_svr","param":{"song_mid":"`, songMid, `","song_type":0}}}`)
	var infoResp struct {
		Code int `json:"code"`
		// Ts      int64  `json:"ts"`
		// StartTs int64  `json:"start_ts"`
		// Traceid string `json:"traceid"`
		Req struct {
			Code int       `json:"code"`
			Data musicInfo `json:"data"`
		} `json:"req"`
	}
	err := signRequest(bytesconv.StringToBytes(infoReqBody), &infoResp)
	if err != nil {
		emsg = err.Error()
		return //nil, err.Error()
	}
	if infoResp.Code != 0 || infoResp.Req.Code != 0 {
		emsg = `获取音乐信息失败`
		return //nil, `获取音乐信息失败`
	}
	infoBody = infoResp.Req.Data
	return //infoBody.Req.Data, ``
}
