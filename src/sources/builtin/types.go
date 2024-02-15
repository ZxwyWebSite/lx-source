package builtin

import (
	"encoding/base64"
	"lx-source/src/sources"

	"github.com/ZxwyWebSite/ztool"
	"github.com/ZxwyWebSite/ztool/zcypt"
)

type (
	// 网易音乐接口 (方格/简繁)
	// WyApi_Song struct {
	// 	Data []struct {
	// 		ID                 int         `json:"id"`
	// 		URL                string      `json:"url"`
	// 		Br                 int         `json:"br"`
	// 		Size               int         `json:"size"`
	// 		Md5                string      `json:"md5"`
	// 		Code               int         `json:"code"`
	// 		Expi               int         `json:"expi"`
	// 		Type               string      `json:"type"`
	// 		Gain               float64     `json:"gain"`
	// 		Peak               float64     `json:"peak"`
	// 		Fee                int         `json:"fee"`
	// 		Uf                 interface{} `json:"uf"`
	// 		Payed              int         `json:"payed"`
	// 		Flag               int         `json:"flag"`
	// 		CanExtend          bool        `json:"canExtend"`
	// 		FreeTrialInfo      interface{} `json:"freeTrialInfo"`
	// 		Level              string      `json:"level"`
	// 		EncodeType         string      `json:"encodeType"`
	// 		FreeTrialPrivilege struct {
	// 			ResConsumable      bool        `json:"resConsumable"`
	// 			UserConsumable     bool        `json:"userConsumable"`
	// 			ListenType         interface{} `json:"listenType"`
	// 			CannotListenReason interface{} `json:"cannotListenReason"`
	// 			PlayReason         interface{} `json:"playReason"`
	// 		} `json:"freeTrialPrivilege"`
	// 		FreeTimeTrialPrivilege struct {
	// 			ResConsumable  bool `json:"resConsumable"`
	// 			UserConsumable bool `json:"userConsumable"`
	// 			Type           int  `json:"type"`
	// 			RemainTime     int  `json:"remainTime"`
	// 		} `json:"freeTimeTrialPrivilege"`
	// 		URLSource   int         `json:"urlSource"`
	// 		RightSource int         `json:"rightSource"`
	// 		PodcastCtrp interface{} `json:"podcastCtrp"`
	// 		EffectTypes interface{} `json:"effectTypes"`
	// 		Time        int         `json:"time"`
	// 	} `json:"data"`
	// 	Code int `json:"code"`
	// }
	// WyApi_Vef struct {
	// 	Code    int16  `json:"code"`
	// 	Success bool   `json:"success"`
	// 	Message string `json:"message"`
	// }
	// 咪咕音乐接口
	MgApi_Song struct {
		Code string `json:"code"`
		Msg  string `json:"msg"`
		Data struct {
			PlayURL         string `json:"playUrl"`
			FormatID        string `json:"formatId"`
			SalePrice       string `json:"salePrice"`
			BizType         string `json:"bizType"`
			BizCode         string `json:"bizCode"`
			AuditionsLength int    `json:"auditionsLength"`
		} `json:"data"`
	}
	// 酷我音乐接口 (波点)
	// KwApi_Song struct {
	// 	Code  int    `json:"code"`
	// 	Msg   string `json:"msg"`
	// 	ReqID string `json:"reqId"`
	// 	Data  struct {
	// 		Duration  int `json:"duration"`
	// 		AudioInfo struct {
	// 			Bitrate string `json:"bitrate"`
	// 			Format  string `json:"format"`
	// 			Level   string `json:"level"`
	// 			Size    string `json:"size"`
	// 		} `json:"audioInfo"`
	// 		URL string `json:"url"`
	// 	} `json:"data"`
	// 	ProfileID string `json:"profileId"`
	// 	CurTime   int64  `json:"curTime"`
	// }
	// 酷狗试听接口
	// KgApi_Song struct {
	// 	Status  int `json:"status"`
	// 	ErrCode int `json:"err_code"`
	// 	Data    any `json:"data"`
	// }
	// KgApi_Data struct {
	// 	// Hash       string `json:"hash"`
	// 	// Timelength int    `json:"timelength"`
	// 	// Filesize   int    `json:"filesize"`
	// 	// AudioName  string `json:"audio_name"`
	// 	// HaveAlbum  int    `json:"have_album"`
	// 	// AlbumName  string `json:"album_name"`
	// 	// AlbumID    any    `json:"album_id"`
	// 	// Img        string `json:"img"`
	// 	// HaveMv     int    `json:"have_mv"`
	// 	// VideoID    any    `json:"video_id"`
	// 	// AuthorName string `json:"author_name"`
	// 	// SongName   string `json:"song_name"`
	// 	// Lyrics     string `json:"lyrics"`
	// 	// AuthorID   any    `json:"author_id"`
	// 	// Privilege  int    `json:"privilege"`
	// 	// Privilege2 string `json:"privilege2"`
	// 	PlayURL string `json:"play_url"`
	// 	// Authors    []struct {
	// 	// 	AuthorID      any    `json:"author_id"`
	// 	// 	AuthorName    string `json:"author_name"`
	// 	// 	IsPublish     string `json:"is_publish"`
	// 	// 	SizableAvatar string `json:"sizable_avatar"`
	// 	// 	EAuthorID     string `json:"e_author_id"`
	// 	// 	Avatar        string `json:"avatar"`
	// 	// } `json:"authors"`
	// 	// IsFreePart         int    `json:"is_free_part"`
	// 	// Bitrate            int    `json:"bitrate"`
	// 	// RecommendAlbumID   string `json:"recommend_album_id"`
	// 	// StoreType          string `json:"store_type"`
	// 	// AlbumAudioID       int    `json:"album_audio_id"`
	// 	// IsPublish          int    `json:"is_publish"`
	// 	// EAuthorID          string `json:"e_author_id"`
	// 	// AudioID            any    `json:"audio_id"`
	// 	// HasPrivilege       bool   `json:"has_privilege"`
	// 	PlayBackupURL string `json:"play_backup_url"`
	// 	// SmallLibrarySong   int    `json:"small_library_song"`
	// 	// EncodeAlbumID      string `json:"encode_album_id"`
	// 	// EncodeAlbumAudioID string `json:"encode_album_audio_id"`
	// 	// EVideoID           string `json:"e_video_id"`
	// }
	// 腾讯试听接口
	// res_tx struct {
	// 	Code int `json:"code"`
	// 	// Ts      int64  `json:"ts"`
	// 	// StartTs int64  `json:"start_ts"`
	// 	// Traceid string `json:"traceid"`
	// 	// Req     struct {
	// 	// 	Code int `json:"code"`
	// 	// 	Data struct {
	// 	// 		Expiration    int      `json:"expiration"`
	// 	// 		Freeflowsip   []string `json:"freeflowsip"`
	// 	// 		Keepalivefile string   `json:"keepalivefile"`
	// 	// 		Msg           string   `json:"msg"`
	// 	// 		Retcode       int      `json:"retcode"`
	// 	// 		Servercheck   string   `json:"servercheck"`
	// 	// 		Sip           []string `json:"sip"`
	// 	// 		Testfile2G    string   `json:"testfile2g"`
	// 	// 		Testfilewifi  string   `json:"testfilewifi"`
	// 	// 		Uin           string   `json:"uin"`
	// 	// 		Userip        string   `json:"userip"`
	// 	// 		Vkey          string   `json:"vkey"`
	// 	// 	} `json:"data"`
	// 	// } `json:"req"`
	// 	Req0 struct {
	// 		Code int `json:"code"`
	// 		Data struct {
	// 			// Uin          string   `json:"uin"`
	// 			// Retcode      int      `json:"retcode"`
	// 			// VerifyType   int      `json:"verify_type"`
	// 			// LoginKey     string   `json:"login_key"`
	// 			// Msg          string   `json:"msg"`
	// 			// Sip          []string `json:"sip"`
	// 			// Thirdip      []string `json:"thirdip"`
	// 			// Testfile2G   string   `json:"testfile2g"`
	// 			// Testfilewifi string   `json:"testfilewifi"`
	// 			Midurlinfo []struct {
	// 				// Songmid           string `json:"songmid"`
	// 				// Filename          string `json:"filename"`
	// 				Purl string `json:"purl"`
	// 				// Errtype           string `json:"errtype"`
	// 				// P2Pfromtag        int    `json:"p2pfromtag"`
	// 				// Qmdlfromtag       int    `json:"qmdlfromtag"`
	// 				// CommonDownfromtag int    `json:"common_downfromtag"`
	// 				// VipDownfromtag    int    `json:"vip_downfromtag"`
	// 				// Pdl               int    `json:"pdl"`
	// 				// Premain           int    `json:"premain"`
	// 				// Hisdown           int    `json:"hisdown"`
	// 				// Hisbuy            int    `json:"hisbuy"`
	// 				// UIAlert           int    `json:"uiAlert"`
	// 				// Isbuy             int    `json:"isbuy"`
	// 				// Pneedbuy          int    `json:"pneedbuy"`
	// 				// Pneed             int    `json:"pneed"`
	// 				// Isonly            int    `json:"isonly"`
	// 				// Onecan            int    `json:"onecan"`
	// 				// Result            int    `json:"result"`
	// 				// Tips              string `json:"tips"`
	// 				// Opi48Kurl         string `json:"opi48kurl"`
	// 				// Opi96Kurl         string `json:"opi96kurl"`
	// 				// Opi192Kurl        string `json:"opi192kurl"`
	// 				// Opiflackurl       string `json:"opiflackurl"`
	// 				// Opi128Kurl        string `json:"opi128kurl"`
	// 				// Opi192Koggurl     string `json:"opi192koggurl"`
	// 				// Wififromtag       string `json:"wififromtag"`
	// 				// Flowfromtag       string `json:"flowfromtag"`
	// 				// Wifiurl           string `json:"wifiurl"`
	// 				// Flowurl           string `json:"flowurl"`
	// 				// Vkey              string `json:"vkey"`
	// 				// Opi30Surl         string `json:"opi30surl"`
	// 				// Ekey              string `json:"ekey"`
	// 				// AuthSwitch        int    `json:"auth_switch"`
	// 				// Subcode           int    `json:"subcode"`
	// 				// Opi96Koggurl      string `json:"opi96koggurl"`
	// 				// AuthSwitch2       int    `json:"auth_switch2"`
	// 			} `json:"midurlinfo"`
	// 			// Servercheck string `json:"servercheck"`
	// 			// Expiration  int    `json:"expiration"`
	// 		} `json:"data"`
	// 	} `json:"req_0"`
	// }
)

// const (
// 	// Source
// 	s_wy = `wy`
// 	s_mg = `mg`
// 	s_kw = `kw`
// 	s_kg = `kg`
// 	s_tx = `tx`
// 	// s_lx = `lx`
// )

var (
	// 音质列表 ( [通用音质][音乐平台]对应音质 )
	/*
	 注: kg源使用对应hash匹配音质，故为空
	*/
	qualitys = map[string]map[string]string{
		sources.Q_128k: {
			sources.S_wy: `standard`,
			sources.S_mg: `1`,
			sources.S_kw: sources.Q_128k,
			sources.S_kg: ``,
			sources.S_tx: `M500`,
		},
		sources.Q_320k: {
			sources.S_wy: `exhigh`,
			sources.S_mg: `2`,
			sources.S_kw: sources.Q_320k,
			// sources.S_kg: ``,
			sources.S_tx: `M800`,
		},
		sources.Q_flac: {
			sources.S_wy: `lossless`,
			sources.S_mg: `3`,
			sources.S_kw: `2000k`,
			sources.S_tx: `F000`,
		},
		sources.Q_fl24: {
			sources.S_wy: `hires`,
			sources.S_mg: `4`,
			// sources.S_tx: `RS01`,
		},
		// `fl24`: {
		// 	s_wy: `hires`,
		// 	s_mg: `4`,
		// },
	}
	// ApiAddr
	api_wy string
	api_mg string
	// api_kw string
	// api_kg string = `https://wwwapi.kugou.com/yy/index.php?r=play/getdata&platid=4&mid=1`
	// api_tx string = `https://u.y.qq.com/cgi-bin/musicu.fcg?data=`
	vef_wy string
	// Headers
	header_wy map[string]string
	header_mg map[string]string
	// header_kw map[string]string
	// header_tx = map[string]string{`Referer`: `https://y.qq.com/`}
)

func init() {
	// InitBuiltInSource
	var initdata = struct {
		Api_Wy    *string
		Api_Mg    *string
		Vef_Wy    *string
		Header_Wy *map[string]string
		Header_Mg *map[string]string
	}{
		Api_Wy:    &api_wy,
		Api_Mg:    &api_mg,
		Vef_Wy:    &vef_wy,
		Header_Wy: &header_wy,
		Header_Mg: &header_mg,
	}
	data := []byte{0x53, 0x6e, 0x38, 0x44, 0x41, 0x51, 0x4c, 0x2f, 0x67, 0x41, 0x41, 0x42, 0x42, 0x51, 0x45, 0x47, 0x51, 0x58, 0x42, 0x70, 0x58, 0x31, 0x64, 0x35, 0x41, 0x51, 0x77, 0x41, 0x41, 0x51, 0x5a, 0x42, 0x63, 0x47, 0x6c, 0x66, 0x54, 0x57, 0x63, 0x42, 0x44, 0x41, 0x41, 0x42, 0x42, 0x6c, 0x5a, 0x6c, 0x5a, 0x6c, 0x39, 0x58, 0x65, 0x51, 0x45, 0x4d, 0x41, 0x41, 0x45, 0x4a, 0x53, 0x47, 0x56, 0x68, 0x5a, 0x47, 0x56, 0x79, 0x58, 0x31, 0x64, 0x35, 0x41, 0x66, 0x2b, 0x43, 0x41, 0x41, 0x45, 0x4a, 0x53, 0x47, 0x56, 0x68, 0x5a, 0x47, 0x56, 0x79, 0x58, 0x30, 0x31, 0x6e, 0x41, 0x66, 0x2b, 0x43, 0x41, 0x41, 0x41, 0x41, 0x49, 0x66, 0x2b, 0x42, 0x42, 0x41, 0x45, 0x42, 0x45, 0x57, 0x31, 0x68, 0x63, 0x46, 0x74, 0x7a, 0x64, 0x48, 0x4a, 0x70, 0x62, 0x6d, 0x64, 0x64, 0x63, 0x33, 0x52, 0x79, 0x61, 0x57, 0x35, 0x6e, 0x41, 0x66, 0x2b, 0x43, 0x41, 0x41, 0x45, 0x4d, 0x41, 0x51, 0x77, 0x41, 0x41, 0x50, 0x34, 0x42, 0x72, 0x76, 0x2b, 0x41, 0x41, 0x53, 0x52, 0x6a, 0x63, 0x32, 0x30, 0x75, 0x63, 0x32, 0x46, 0x35, 0x63, 0x58, 0x6f, 0x75, 0x59, 0x32, 0x39, 0x74, 0x4c, 0x32, 0x46, 0x77, 0x61, 0x53, 0x38, 0x2f, 0x64, 0x48, 0x6c, 0x77, 0x5a, 0x54, 0x31, 0x68, 0x63, 0x47, 0x6c, 0x54, 0x62, 0x32, 0x35, 0x6e, 0x56, 0x58, 0x4a, 0x73, 0x56, 0x6a, 0x45, 0x42, 0x4e, 0x6d, 0x30, 0x75, 0x62, 0x58, 0x56, 0x7a, 0x61, 0x57, 0x4d, 0x75, 0x62, 0x57, 0x6c, 0x6e, 0x64, 0x53, 0x35, 0x6a, 0x62, 0x69, 0x39, 0x74, 0x61, 0x57, 0x64, 0x31, 0x62, 0x58, 0x56, 0x7a, 0x61, 0x57, 0x4d, 0x76, 0x61, 0x44, 0x55, 0x76, 0x63, 0x47, 0x78, 0x68, 0x65, 0x53, 0x39, 0x68, 0x64, 0x58, 0x52, 0x6f, 0x4c, 0x32, 0x64, 0x6c, 0x64, 0x46, 0x4e, 0x76, 0x62, 0x6d, 0x64, 0x51, 0x62, 0x47, 0x46, 0x35, 0x53, 0x57, 0x35, 0x6d, 0x62, 0x77, 0x45, 0x6c, 0x59, 0x33, 0x4e, 0x74, 0x4c, 0x6e, 0x4e, 0x68, 0x65, 0x58, 0x46, 0x36, 0x4c, 0x6d, 0x4e, 0x76, 0x62, 0x53, 0x39, 0x68, 0x63, 0x47, 0x6b, 0x76, 0x50, 0x33, 0x52, 0x35, 0x63, 0x47, 0x55, 0x39, 0x59, 0x33, 0x4e, 0x74, 0x51, 0x32, 0x68, 0x6c, 0x59, 0x57, 0x74, 0x4e, 0x64, 0x58, 0x4e, 0x70, 0x59, 0x77, 0x45, 0x43, 0x43, 0x6c, 0x56, 0x7a, 0x5a, 0x58, 0x49, 0x74, 0x51, 0x57, 0x64, 0x6c, 0x62, 0x6e, 0x52, 0x65, 0x54, 0x57, 0x39, 0x36, 0x61, 0x57, 0x78, 0x73, 0x59, 0x53, 0x38, 0x31, 0x4c, 0x6a, 0x41, 0x67, 0x4b, 0x46, 0x64, 0x70, 0x62, 0x6d, 0x52, 0x76, 0x64, 0x33, 0x4d, 0x67, 0x54, 0x6c, 0x51, 0x67, 0x4e, 0x69, 0x34, 0x78, 0x4f, 0x79, 0x42, 0x58, 0x54, 0x31, 0x63, 0x32, 0x4e, 0x43, 0x6b, 0x67, 0x51, 0x58, 0x42, 0x77, 0x62, 0x47, 0x56, 0x58, 0x5a, 0x57, 0x4a, 0x4c, 0x61, 0x58, 0x51, 0x76, 0x4e, 0x54, 0x4d, 0x33, 0x4c, 0x6a, 0x4d, 0x32, 0x49, 0x43, 0x68, 0x4c, 0x53, 0x46, 0x52, 0x4e, 0x54, 0x43, 0x77, 0x67, 0x62, 0x47, 0x6c, 0x72, 0x5a, 0x53, 0x42, 0x48, 0x5a, 0x57, 0x4e, 0x72, 0x62, 0x79, 0x6b, 0x67, 0x51, 0x32, 0x68, 0x79, 0x62, 0x32, 0x31, 0x6c, 0x4c, 0x7a, 0x55, 0x77, 0x4c, 0x6a, 0x41, 0x75, 0x4d, 0x6a, 0x59, 0x32, 0x4d, 0x53, 0x34, 0x34, 0x4e, 0x78, 0x42, 0x59, 0x4c, 0x56, 0x4a, 0x6c, 0x63, 0x58, 0x56, 0x6c, 0x63, 0x33, 0x52, 0x6c, 0x5a, 0x43, 0x31, 0x58, 0x61, 0x58, 0x52, 0x6f, 0x44, 0x6c, 0x68, 0x4e, 0x54, 0x45, 0x68, 0x30, 0x64, 0x48, 0x42, 0x53, 0x5a, 0x58, 0x46, 0x31, 0x5a, 0x58, 0x4e, 0x30, 0x41, 0x51, 0x51, 0x48, 0x55, 0x6d, 0x56, 0x6d, 0x5a, 0x58, 0x4a, 0x6c, 0x63, 0x68, 0x74, 0x6f, 0x64, 0x48, 0x52, 0x77, 0x63, 0x7a, 0x6f, 0x76, 0x4c, 0x32, 0x30, 0x75, 0x62, 0x58, 0x56, 0x7a, 0x61, 0x57, 0x4d, 0x75, 0x62, 0x57, 0x6c, 0x6e, 0x64, 0x53, 0x35, 0x6a, 0x62, 0x69, 0x39, 0x32, 0x4e, 0x43, 0x38, 0x43, 0x51, 0x6e, 0x6b, 0x67, 0x4d, 0x44, 0x52, 0x6d, 0x4f, 0x44, 0x45, 0x30, 0x4e, 0x6a, 0x46, 0x68, 0x4f, 0x54, 0x68, 0x6a, 0x4e, 0x32, 0x46, 0x6d, 0x4e, 0x54, 0x55, 0x33, 0x5a, 0x6d, 0x56, 0x68, 0x4d, 0x32, 0x4e, 0x6d, 0x4d, 0x6a, 0x68, 0x6a, 0x4e, 0x47, 0x56, 0x68, 0x4d, 0x54, 0x55, 0x48, 0x59, 0x32, 0x68, 0x68, 0x62, 0x6d, 0x35, 0x6c, 0x62, 0x41, 0x63, 0x77, 0x4d, 0x54, 0x51, 0x77, 0x4d, 0x44, 0x42, 0x45, 0x42, 0x6b, 0x4e, 0x76, 0x62, 0x32, 0x74, 0x70, 0x5a, 0x54, 0x68, 0x54, 0x52, 0x56, 0x4e, 0x54, 0x53, 0x55, 0x39, 0x4f, 0x50, 0x56, 0x70, 0x55, 0x53, 0x58, 0x64, 0x50, 0x52, 0x47, 0x74, 0x35, 0x54, 0x55, 0x52, 0x52, 0x64, 0x45, 0x39, 0x55, 0x52, 0x54, 0x46, 0x4f, 0x55, 0x7a, 0x41, 0x77, 0x54, 0x55, 0x52, 0x6f, 0x62, 0x45, 0x78, 0x55, 0x61, 0x47, 0x68, 0x4e, 0x56, 0x30, 0x56, 0x30, 0x54, 0x57, 0x70, 0x52, 0x4d, 0x45, 0x34, 0x79, 0x57, 0x54, 0x4a, 0x4e, 0x65, 0x6d, 0x73, 0x79, 0x54, 0x31, 0x52, 0x42, 0x65, 0x67, 0x41, 0x3d}
	dec, _ := zcypt.Base64Decode(base64.StdEncoding, data)
	ztool.Val_GobDecode(dec, &initdata)
}
