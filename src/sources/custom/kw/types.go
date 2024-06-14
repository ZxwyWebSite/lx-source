package kw

type (
	// KwDES json
	playInfo struct {
		Code       int    `json:"code"`
		Locationid string `json:"locationid"`
		Data       struct {
			Bitrate          int    `json:"bitrate"`
			User             string `json:"user"`
			Sig              string `json:"sig"`
			Type             string `json:"type"`
			Format           string `json:"format"`
			P2PAudiosourceid string `json:"p2p_audiosourceid"`
			Rid              int    `json:"rid"`
			// Source string `json:"source"`
			URL string `json:"url"`
		} `json:"data"`
		Msg string `json:"msg"`
	}
	// 酷我音乐接口 (波点)
	kwApi_Song struct {
		Code  int    `json:"code"`
		Msg   string `json:"msg"`
		ReqID string `json:"reqId"`
		Data  struct {
			Duration  int `json:"duration"`
			AudioInfo struct {
				Bitrate string `json:"bitrate"`
				Format  string `json:"format"`
				Level   string `json:"level"`
				Size    string `json:"size"`
			} `json:"audioInfo"`
			URL string `json:"url"`
		} `json:"data"`
		ProfileID string `json:"profileId"`
		CurTime   int64  `json:"curTime"`
	}
)
