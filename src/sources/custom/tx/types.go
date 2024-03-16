package tx

type musicInfo struct {
	// Info struct {
	// 	Company struct {
	// 		Title   string `json:"title"`
	// 		Type    string `json:"type"`
	// 		Content []struct {
	// 			ID        int    `json:"id"`
	// 			Value     string `json:"value"`
	// 			Mid       string `json:"mid"`
	// 			Type      int    `json:"type"`
	// 			ShowType  int    `json:"show_type"`
	// 			IsParent  int    `json:"is_parent"`
	// 			Picurl    string `json:"picurl"`
	// 			ReadCnt   int    `json:"read_cnt"`
	// 			Author    string `json:"author"`
	// 			Jumpurl   string `json:"jumpurl"`
	// 			OriPicurl string `json:"ori_picurl"`
	// 		} `json:"content"`
	// 		Pos         int    `json:"pos"`
	// 		More        int    `json:"more"`
	// 		Selected    string `json:"selected"`
	// 		UsePlatform int    `json:"use_platform"`
	// 	} `json:"company"`
	// 	Genre struct {
	// 		Title   string `json:"title"`
	// 		Type    string `json:"type"`
	// 		Content []struct {
	// 			ID        int    `json:"id"`
	// 			Value     string `json:"value"`
	// 			Mid       string `json:"mid"`
	// 			Type      int    `json:"type"`
	// 			ShowType  int    `json:"show_type"`
	// 			IsParent  int    `json:"is_parent"`
	// 			Picurl    string `json:"picurl"`
	// 			ReadCnt   int    `json:"read_cnt"`
	// 			Author    string `json:"author"`
	// 			Jumpurl   string `json:"jumpurl"`
	// 			OriPicurl string `json:"ori_picurl"`
	// 		} `json:"content"`
	// 		Pos         int    `json:"pos"`
	// 		More        int    `json:"more"`
	// 		Selected    string `json:"selected"`
	// 		UsePlatform int    `json:"use_platform"`
	// 	} `json:"genre"`
	// 	Lan struct {
	// 		Title   string `json:"title"`
	// 		Type    string `json:"type"`
	// 		Content []struct {
	// 			ID        int    `json:"id"`
	// 			Value     string `json:"value"`
	// 			Mid       string `json:"mid"`
	// 			Type      int    `json:"type"`
	// 			ShowType  int    `json:"show_type"`
	// 			IsParent  int    `json:"is_parent"`
	// 			Picurl    string `json:"picurl"`
	// 			ReadCnt   int    `json:"read_cnt"`
	// 			Author    string `json:"author"`
	// 			Jumpurl   string `json:"jumpurl"`
	// 			OriPicurl string `json:"ori_picurl"`
	// 		} `json:"content"`
	// 		Pos         int    `json:"pos"`
	// 		More        int    `json:"more"`
	// 		Selected    string `json:"selected"`
	// 		UsePlatform int    `json:"use_platform"`
	// 	} `json:"lan"`
	// } `json:"info"`
	// Extras struct {
	// 	Name      string `json:"name"`
	// 	Transname string `json:"transname"`
	// 	Subtitle  string `json:"subtitle"`
	// 	From      string `json:"from"`
	// 	Wikiurl   string `json:"wikiurl"`
	// } `json:"extras"`
	TrackInfo struct {
		// ID       int    `json:"id"`
		// Type     int    `json:"type"`
		// Mid      string `json:"mid"`
		// Name     string `json:"name"`
		// Title    string `json:"title"`
		// Subtitle string `json:"subtitle"`
		// Singer   []struct {
		// 	ID    int    `json:"id"`
		// 	Mid   string `json:"mid"`
		// 	Name  string `json:"name"`
		// 	Title string `json:"title"`
		// 	Type  int    `json:"type"`
		// 	Uin   int    `json:"uin"`
		// } `json:"singer"`
		// Album struct {
		// 	ID         int    `json:"id"`
		// 	Mid        string `json:"mid"`
		// 	Name       string `json:"name"`
		// 	Title      string `json:"title"`
		// 	Subtitle   string `json:"subtitle"`
		// 	TimePublic string `json:"time_public"`
		// 	Pmid       string `json:"pmid"`
		// } `json:"album"`
		// Mv struct {
		// 	ID    int    `json:"id"`
		// 	Vid   string `json:"vid"`
		// 	Name  string `json:"name"`
		// 	Title string `json:"title"`
		// 	Vt    int    `json:"vt"`
		// } `json:"mv"`
		// Interval   int    `json:"interval"`
		// Isonly     int    `json:"isonly"`
		// Language   int    `json:"language"`
		// Genre      int    `json:"genre"`
		// IndexCd    int    `json:"index_cd"`
		// IndexAlbum int    `json:"index_album"`
		// TimePublic string `json:"time_public"`
		// Status     int    `json:"status"`
		// Fnote      int    `json:"fnote"`
		File struct {
			MediaMid string `json:"media_mid"`
			// Size24Aac     int           `json:"size_24aac"`
			// Size48Aac     int           `json:"size_48aac"`
			// Size96Aac     int           `json:"size_96aac"`
			// Size192Ogg    int           `json:"size_192ogg"`
			// Size192Aac    int           `json:"size_192aac"`
			// Size128Mp3    int           `json:"size_128mp3"`
			// Size320Mp3    int           `json:"size_320mp3"`
			// SizeApe       int           `json:"size_ape"`
			// SizeFlac      int           `json:"size_flac"`
			// SizeDts       int           `json:"size_dts"`
			// SizeTry       int           `json:"size_try"`
			// TryBegin      int           `json:"try_begin"`
			// TryEnd        int           `json:"try_end"`
			// URL           string        `json:"url"`
			// SizeHires     int           `json:"size_hires"`
			// HiresSample   int           `json:"hires_sample"`
			// HiresBitdepth int           `json:"hires_bitdepth"`
			// B30S          int           `json:"b_30s"`
			// E30S          int           `json:"e_30s"`
			// Size96Ogg     int           `json:"size_96ogg"`
			// Size360Ra     []interface{} `json:"size_360ra"`
			// SizeDolby     int           `json:"size_dolby"`
			// SizeNew       []int         `json:"size_new"`
		} `json:"file"`
		Pay struct {
			// PayMonth   int `json:"pay_month"`
			// PriceTrack int `json:"price_track"`
			// PriceAlbum int `json:"price_album"`
			PayPlay int `json:"pay_play"`
			// PayDown    int `json:"pay_down"`
			// PayStatus  int `json:"pay_status"`
			// TimeFree   int `json:"time_free"`
		} `json:"pay"`
		// Action struct {
		// 	Switch   int `json:"switch"`
		// 	Msgid    int `json:"msgid"`
		// 	Alert    int `json:"alert"`
		// 	Icons    int `json:"icons"`
		// 	Msgshare int `json:"msgshare"`
		// 	Msgfav   int `json:"msgfav"`
		// 	Msgdown  int `json:"msgdown"`
		// 	Msgpay   int `json:"msgpay"`
		// 	Switch2  int `json:"switch2"`
		// 	Icon2    int `json:"icon2"`
		// } `json:"action"`
		// Ksong struct {
		// 	ID  int    `json:"id"`
		// 	Mid string `json:"mid"`
		// } `json:"ksong"`
		// Volume struct {
		// 	Gain float64 `json:"gain"`
		// 	Peak float64 `json:"peak"`
		// 	Lra  float64 `json:"lra"`
		// } `json:"volume"`
		// Label       string    `json:"label"`
		// URL         string    `json:"url"`
		// Bpm         int       `json:"bpm"`
		// Version     int       `json:"version"`
		// Trace       string    `json:"trace"`
		// DataType    int       `json:"data_type"`
		// ModifyStamp int       `json:"modify_stamp"`
		// Pingpong    string    `json:"pingpong"`
		// Ppurl       string    `json:"ppurl"`
		// Tid         int       `json:"tid"`
		// Ov          int       `json:"ov"`
		// Sa          int       `json:"sa"`
		// Es          string    `json:"es"`
		Vs []string `json:"vs"`
		// Vi          []int     `json:"vi"`
		// Ktag        string    `json:"ktag"`
		// Vf          []float64 `json:"vf"`
	} `json:"track_info"`
}

type playInfo struct {
	// Code int `json:"code"`
	Data struct {
		// Uin          string   `json:"uin"`
		// Retcode      int      `json:"retcode"`
		// VerifyType   int      `json:"verify_type"`
		// LoginKey     string   `json:"login_key"`
		// Msg          string   `json:"msg"`
		// Sip          []string `json:"sip"`
		// Thirdip      []string `json:"thirdip"`
		// Testfile2G   string   `json:"testfile2g"`
		// Testfilewifi string   `json:"testfilewifi"`
		Midurlinfo []struct {
			// Songmid           string `json:"songmid"`
			Filename string `json:"filename"`
			Purl     string `json:"purl"`
			// Errtype           string `json:"errtype"`
			// P2Pfromtag        int    `json:"p2pfromtag"`
			// Qmdlfromtag       int    `json:"qmdlfromtag"`
			// CommonDownfromtag int    `json:"common_downfromtag"`
			// VipDownfromtag    int    `json:"vip_downfromtag"`
			// Pdl               int    `json:"pdl"`
			// Premain           int    `json:"premain"`
			// Hisdown           int    `json:"hisdown"`
			// Hisbuy            int    `json:"hisbuy"`
			// UIAlert           int    `json:"uiAlert"`
			// Isbuy             int    `json:"isbuy"`
			// Pneedbuy          int    `json:"pneedbuy"`
			// Pneed             int    `json:"pneed"`
			// Isonly            int    `json:"isonly"`
			// Onecan            int    `json:"onecan"`
			// Result            int    `json:"result"`
			// Tips              string `json:"tips"`
			// Opi48Kurl         string `json:"opi48kurl"`
			// Opi96Kurl         string `json:"opi96kurl"`
			// Opi192Kurl        string `json:"opi192kurl"`
			// Opiflackurl       string `json:"opiflackurl"`
			// Opi128Kurl        string `json:"opi128kurl"`
			// Opi192Koggurl     string `json:"opi192koggurl"`
			// Wififromtag       string `json:"wififromtag"`
			// Flowfromtag       string `json:"flowfromtag"`
			// Wifiurl           string `json:"wifiurl"`
			// Flowurl           string `json:"flowurl"`
			// Vkey              string `json:"vkey"`
			// Opi30Surl         string `json:"opi30surl"`
			// Ekey              string `json:"ekey"`
			// AuthSwitch        int    `json:"auth_switch"`
			// Subcode           int    `json:"subcode"`
			// Opi96Koggurl      string `json:"opi96koggurl"`
			// AuthSwitch2       int    `json:"auth_switch2"`
		} `json:"midurlinfo"`
		// Servercheck string `json:"servercheck"`
		// Expiration  int    `json:"expiration"`
	} `json:"data"`
}
