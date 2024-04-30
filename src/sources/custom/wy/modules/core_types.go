package wy

type (
	// 音乐URL
	PlayInfo struct {
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
			FreeTrialInfo *struct {
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
	// 音乐是否可用
	VerifyInfo struct {
		Code    int16  `json:"code"`
		Success bool   `json:"success"`
		Message string `json:"message"`
	}
	// 音质数据
	QualityData struct {
		Br   int     `json:"br"`   // 比特率 Bit Rate
		Fid  int     `json:"fid"`  // ?
		Size int     `json:"size"` // 文件大小
		Vd   float64 `json:"vd"`   // Volume Delta
		Sr   int     `json:"sr"`   // 采样率 Sample Rate
	}
	// 歌曲音质详情
	QualityDetail struct {
		Data struct {
			SongID int         `json:"songId"`
			H      QualityData `json:"h"`  // 高质量文件信息
			M      QualityData `json:"m"`  // 中质量文件信息
			L      QualityData `json:"l"`  // 低质量文件信息
			Sq     QualityData `json:"sq"` // 无损质量文件信息
			Hr     QualityData `json:"hr"` // Hi-Res质量文件信息
			Db     QualityData `json:"db"` // 杜比音质
			Jm     QualityData `json:"jm"` // jymaster(超清母带)
			Je     QualityData `json:"je"` // jyeffect(高清环绕声)
			Sk     QualityData `json:"sk"` // sky(沉浸环绕声)
		} `json:"data"`
		Code    int    `json:"code"`
		Message string `json:"message"`
		Success bool   `json:"success"`
		Error   bool   `json:"error"`
	}
	// 扫码登录请求
	QrKey struct {
		Code   int    `json:"code"`
		UniKey string `json:"unikey"`
	}
	// 扫码登录结果
	QrCheck struct {
		Code      int    `json:"code"`
		Message   string `json:"message"`
		AvatarUrl string `json:"avatarUrl"`
		NickName  string `json:"nickname"`
	}
)
