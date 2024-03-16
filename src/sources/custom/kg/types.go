package kg

type (
	musicInfo struct {
		AlbumAudioID string `json:"album_audio_id"`
		// AudioID          string `json:"audio_id"`
		// AuthorName       string `json:"author_name"`
		// OriAudioName     string `json:"ori_audio_name"`
		// OfficialSongname string `json:"official_songname"`
		// Songname         string `json:"songname"`
		// Remark           string `json:"remark"`
		// SuffixAudioName  string `json:"suffix_audio_name"`
		// IsSearch         string `json:"is_search"`
		// IsOriginal       string `json:"is_original"`
		// IsPublish        string `json:"is_publish"`
		// MixsongType      string `json:"mixsong_type"`
		// Version          string `json:"version"`
		// Language         string `json:"language"`
		// Bpm              string `json:"bpm"`
		// BpmType          string `json:"bpm_type"`
		// BpmDesc          string `json:"bpm_desc"`
		// PublishDate      string `json:"publish_date"`
		// Extname          string `json:"extname"`
		AlbumInfo struct {
			AlbumID string `json:"album_id"`
			// AlbumName    string `json:"album_name"`
			// PublishDate  string `json:"publish_date"`
			// Category     string `json:"category"`
			// IsPublish    string `json:"is_publish"`
			// SizableCover string `json:"sizable_cover"`
		} `json:"album_info"`
		// Classification []struct {
		// 	Usage   string `json:"usage"`
		// 	Status  string `json:"status"`
		// 	Level   string `json:"level"`
		// 	ResType string `json:"res_type"`
		// 	Type    string `json:"type"`
		// 	ID      string `json:"id"`
		// 	ResID   string `json:"res_id"`
		// } `json:"classification"`
		AudioInfo struct {
			// IsFileHead             string `json:"is_file_head"`
			// IsFileHead320          string `json:"is_file_head_320"`
			// AudioID                string `json:"audio_id"`
			// Hash                   string `json:"hash"`
			// Filesize               string `json:"filesize"`
			// Timelength             string `json:"timelength"`
			// Bitrate                string `json:"bitrate"`
			Hash128 string `json:"hash_128"`
			// Filesize128            string `json:"filesize_128"`
			// Timelength128          string `json:"timelength_128"`
			Hash320 string `json:"hash_320"`
			// Filesize320            string `json:"filesize_320"`
			// Timelength320          string `json:"timelength_320"`
			HashFlac string `json:"hash_flac"`
			// FilesizeFlac           string `json:"filesize_flac"`
			// TimelengthFlac         string `json:"timelength_flac"`
			// BitrateFlac            string `json:"bitrate_flac"`
			HashHigh string `json:"hash_high"`
			// FilesizeHigh           string `json:"filesize_high"`
			// TimelengthHigh         string `json:"timelength_high"`
			// BitrateHigh            string `json:"bitrate_high"`
			HashSuper string `json:"hash_super"`
			// FilesizeSuper          string `json:"filesize_super"`
			// TimelengthSuper        string `json:"timelength_super"`
			// BitrateSuper           string `json:"bitrate_super"`
			// HashVinylrecord        string `json:"hash_vinylrecord"`
			// FilesizeVinylrecord    string `json:"filesize_vinylrecord"`
			// TimelengthVinylrecord  string `json:"timelength_vinylrecord"`
			// BitrateVinylrecord     string `json:"bitrate_vinylrecord"`
			// HashMultichannel       string `json:"hash_multichannel"`
			// FilesizeMultichannel   string `json:"filesize_multichannel"`
			// TimelengthMultichannel string `json:"timelength_multichannel"`
			// BitrateMultichannel    string `json:"bitrate_multichannel"`
			// HashDolby448           string `json:"hash_dolby_448"`
			// FilesizeDolby448       string `json:"filesize_dolby_448"`
			// TimelengthDolby448     string `json:"timelength_dolby_448"`
			// BitrateDolby448        string `json:"bitrate_dolby_448"`
			// HashDolby640           string `json:"hash_dolby_640"`
			// FilesizeDolby640       string `json:"filesize_dolby_640"`
			// TimelengthDolby640     string `json:"timelength_dolby_640"`
			// BitrateDolby640        string `json:"bitrate_dolby_640"`
			// HashDolby768           string `json:"hash_dolby_768"`
			// FilesizeDolby768       string `json:"filesize_dolby_768"`
			// TimelengthDolby768     string `json:"timelength_dolby_768"`
			// BitrateDolby768        string `json:"bitrate_dolby_768"`
			// AudioGroupID           string `json:"audio_group_id"`
			// ExtnameSuper           string `json:"extname_super"`
			// Extname                string `json:"extname"`
			// FailProcess            int    `json:"fail_process"`
			// PayType                int    `json:"pay_type"`
			// Type                   string `json:"type"`
			// OldCpy                 int    `json:"old_cpy"`
			// Privilege              string `json:"privilege"`
			// Privilege128           string `json:"privilege_128"`
			// Privilege320           string `json:"privilege_320"`
			// PrivilegeFlac          string `json:"privilege_flac"`
			// PrivilegeHigh          string `json:"privilege_high"`
			// PrivilegeSuper         string `json:"privilege_super"`
			// PrivilegeVinylrecord   string `json:"privilege_vinylrecord"`
			// PrivilegeMultichannel  string `json:"privilege_multichannel"`
			// PrivilegeDolby448      string `json:"privilege_dolby_448"`
			// PrivilegeDolby640      string `json:"privilege_dolby_640"`
			// PrivilegeDolby768      string `json:"privilege_dolby_768"`
			// TransParam             struct {
			// 	HashOffset struct {
			// 		StartByte  int    `json:"start_byte"`
			// 		EndByte    int    `json:"end_byte"`
			// 		StartMs    int    `json:"start_ms"`
			// 		EndMs      int    `json:"end_ms"`
			// 		OffsetHash string `json:"offset_hash"`
			// 		FileType   int    `json:"file_type"`
			// 		ClipHash   string `json:"clip_hash"`
			// 	} `json:"hash_offset"`
			// 	MusicpackAdvance int `json:"musicpack_advance"`
			// 	PayBlockTpl      int `json:"pay_block_tpl"`
			// 	Display          int `json:"display"`
			// 	DisplayRate      int `json:"display_rate"`
			// 	CpyGrade         int `json:"cpy_grade"`
			// 	CpyLevel         int `json:"cpy_level"`
			// 	Cid              int `json:"cid"`
			// 	CpyAttr0         int `json:"cpy_attr0"`
			// 	Classmap         struct {
			// 		Attr0 int `json:"attr0"`
			// 	} `json:"classmap"`
			// 	InitPubDay int `json:"init_pub_day"`
			// 	Qualitymap struct {
			// 		Attr0 int `json:"attr0"`
			// 	} `json:"qualitymap"`
			// 	Language string `json:"language"`
			// } `json:"trans_param"`
		} `json:"audio_info"`
	}
	playInfo struct {
		// Hash     string `json:"hash"`
		// Classmap struct {
		// 	Attr0 int `json:"attr0"`
		// } `json:"classmap"`
		Status int `json:"status"`
		// Volume      float64  `json:"volume"`
		// StdHashTime int      `json:"std_hash_time"`
		URL []string `json:"url"`
		// StdHash     string   `json:"std_hash"`
		// TransParam  struct {
		// 	Display     int `json:"display"`
		// 	DisplayRate int `json:"display_rate"`
		// } `json:"trans_param"`
		// FileHead   int     `json:"fileHead"`
		// VolumePeak float64 `json:"volume_peak"`
		// BitRate    int     `json:"bitRate"`
		// TimeLength int     `json:"timeLength"`
		// VolumeGain int     `json:"volume_gain"`
		// Q          int     `json:"q"`
		// FileName   string  `json:"fileName"`
		// ExtName    string  `json:"extName"`
		// FileSize   int     `json:"fileSize"`
	}
)

type rankInfo struct {
	Data struct {
		// Timestamp int `json:"timestamp"`
		// Total     int `json:"total"`
		Info []struct {
			// LastSort int `json:"last_sort"`
			// Authors  []struct {
			// 	SizableAvatar string `json:"sizable_avatar"`
			// 	IsPublish     int    `json:"is_publish"`
			// 	AuthorName    string `json:"author_name"`
			// 	AuthorID      int    `json:"author_id"`
			// } `json:"authors"`
			// RankCount         int    `json:"rank_count"`
			// RankIDPublishDate string `json:"rank_id_publish_date"`
			// Songname          string `json:"songname"`
			// TopicURL320       string `json:"topic_url_320"`
			// Sqhash            string `json:"sqhash"`
			// FailProcess       int    `json:"fail_process"`
			// PayType           int    `json:"pay_type"`
			// RecommendReason   string `json:"recommend_reason"`
			// RpType            string `json:"rp_type"`
			// AlbumID           string `json:"album_id"`
			// PrivilegeHigh     int    `json:"privilege_high"`
			// TopicURLSq        string `json:"topic_url_sq"`
			// RankCid           int    `json:"rank_cid"`
			// Inlist            int    `json:"inlist"`
			// Three20Filesize   int    `json:"320filesize"`
			// PkgPrice320       int    `json:"pkg_price_320"`
			// Feetype           int    `json:"feetype"`
			// Price320          int    `json:"price_320"`
			// DurationHigh      int    `json:"duration_high"`
			// FailProcess320    int    `json:"fail_process_320"`
			// Zone              string `json:"zone"`
			// TopicURL          string `json:"topic_url"`
			// RpPublish         int    `json:"rp_publish"`
			// TransObj          struct {
			// 	RankShowSort int `json:"rank_show_sort"`
			// } `json:"trans_obj"`
			// Hash              string      `json:"hash"`
			// Sqfilesize        int         `json:"sqfilesize"`
			// Sqprivilege       int         `json:"sqprivilege"`
			// PayTypeSq         int         `json:"pay_type_sq"`
			// Bitrate           int         `json:"bitrate"`
			// PkgPriceSq        int         `json:"pkg_price_sq"`
			// HasAccompany      int         `json:"has_accompany"`
			// Musical           interface{} `json:"musical"`
			// PayType320        int         `json:"pay_type_320"`
			// Issue             int         `json:"issue"`
			// ExtnameSuper      string      `json:"extname_super"`
			// DurationSuper     int         `json:"duration_super"`
			// BitrateSuper      int         `json:"bitrate_super"`
			// HashHigh          string      `json:"hash_high"`
			// Duration          int         `json:"duration"`
			// Three20Hash       string      `json:"320hash"`
			// PriceSq           int         `json:"price_sq"`
			// OldCpy            int         `json:"old_cpy"`
			// AlbumAudioID      int         `json:"album_audio_id"`
			// M4Afilesize       int         `json:"m4afilesize"`
			// PkgPrice          int         `json:"pkg_price"`
			// First             int         `json:"first"`
			// AudioID           int         `json:"audio_id"`
			// HashSuper         string      `json:"hash_super"`
			// Addtime           string      `json:"addtime"`
			// FilesizeHigh      int         `json:"filesize_high"`
			// Price             int         `json:"price"`
			// Privilege         int         `json:"privilege"`
			// AlbumSizableCover string      `json:"album_sizable_cover"`
			// Mvdata            []struct {
			// 	Typ  int    `json:"typ"`
			// 	Trk  string `json:"trk"`
			// 	Hash string `json:"hash"`
			// 	ID   string `json:"id"`
			// } `json:"mvdata,omitempty"`
			// Sort       int `json:"sort"`
			// TransParam struct {
			// 	CpyLevel int `json:"cpy_level"`
			// 	Classmap struct {
			// 		Attr0 int `json:"attr0"`
			// 	} `json:"classmap"`
			// 	CpyGrade   int `json:"cpy_grade"`
			// 	Qualitymap struct {
			// 		Attr0 int `json:"attr0"`
			// 	} `json:"qualitymap"`
			// 	PayBlockTpl    int    `json:"pay_block_tpl"`
			// 	Cid            int    `json:"cid"`
			// 	Language       string `json:"language"`
			// 	HashMultitrack string `json:"hash_multitrack"`
			// 	CpyAttr0       int    `json:"cpy_attr0"`
			// 	Ipmap          struct {
			// 		Attr0 int `json:"attr0"`
			// 	} `json:"ipmap"`
			// 	AppidBlock       string `json:"appid_block"`
			// 	MusicpackAdvance int    `json:"musicpack_advance"`
			// 	Display          int    `json:"display"`
			// 	DisplayRate      int    `json:"display_rate"`
			// } `json:"trans_param"`
			// FilesizeSuper    int    `json:"filesize_super"`
			Filename string `json:"filename"`
			// BitrateHigh      int    `json:"bitrate_high"`
			// Remark           string `json:"remark"`
			// Extname          string `json:"extname"`
			// Filesize         int    `json:"filesize"`
			// Isfirst          int    `json:"isfirst"`
			// Mvhash           string `json:"mvhash"`
			// Three20Privilege int    `json:"320privilege"`
			// PrivilegeSuper   int    `json:"privilege_super"`
			// FailProcessSq    int    `json:"fail_process_sq"`
		} `json:"info"`
	} `json:"data"`
	// Errcode int    `json:"errcode"`
	Status int    `json:"status"`
	Error  string `json:"error"`
}

type searchInfo struct {
	Data struct {
		// AlgPath     string `json:"AlgPath"`
		// Aggregation struct {
		// } `json:"aggregation"`
		// Allowerr          int    `json:"allowerr"`
		// Chinesecount      int    `json:"chinesecount"`
		// Correctionforce   int    `json:"correctionforce"`
		// Correctionrelate  string `json:"correctionrelate"`
		// Correctionsubject string `json:"correctionsubject"`
		// Correctiontip     string `json:"correctiontip"`
		// Correctiontype    int    `json:"correctiontype"`
		// From              int    `json:"from"`
		// Isshareresult     int    `json:"isshareresult"`
		// Istag             int    `json:"istag"`
		// Istagresult       int    `json:"istagresult"`
		Lists []struct {
			// A320Privilege  int           `json:"A320Privilege"`
			// ASQPrivilege   int           `json:"ASQPrivilege"`
			// Accompany      int           `json:"Accompany"`
			// AlbumAux       string        `json:"AlbumAux"`
			// AlbumID        string        `json:"AlbumID"`
			// AlbumName      string        `json:"AlbumName"`
			// AlbumPrivilege int           `json:"AlbumPrivilege"`
			// AudioCdn       int           `json:"AudioCdn"`
			// Audioid        int           `json:"Audioid"`
			// Auxiliary      string        `json:"Auxiliary"`
			// Bitrate        int           `json:"Bitrate"`
			// Category       int           `json:"Category"`
			// Duration       int           `json:"Duration"`
			// ExtName        string        `json:"ExtName"`
			// FailProcess    int           `json:"FailProcess"`
			// FileHash       string        `json:"FileHash"`
			// FileName       string        `json:"FileName"`
			// FileSize       int           `json:"FileSize"`
			// FoldType       int           `json:"FoldType"`
			// Grp            []interface{} `json:"Grp"`
			// HQBitrate      int           `json:"HQBitrate"`
			// HQDuration     int           `json:"HQDuration"`
			// HQExtName      string        `json:"HQExtName"`
			// HQFailProcess  int           `json:"HQFailProcess"`
			// HQFileHash     string        `json:"HQFileHash"`
			// HQFileSize     int           `json:"HQFileSize"`
			// HQPayType      int           `json:"HQPayType"`
			// HQPkgPrice     int           `json:"HQPkgPrice"`
			// HQPrice        int           `json:"HQPrice"`
			// HQPrivilege    int           `json:"HQPrivilege"`
			// HasAlbum       int           `json:"HasAlbum"`
			// HeatLevel      int           `json:"HeatLevel"`
			// HiFiQuality    int           `json:"HiFiQuality"`
			// ID             string        `json:"ID"`
			// Image          string        `json:"Image"`
			// IsOriginal     int           `json:"IsOriginal"`
			// M4ASize        int           `json:"M4aSize"`
			// MatchFlag      int           `json:"MatchFlag"`
			MixSongID string `json:"MixSongID"`
			// MvHash         string        `json:"MvHash"`
			// MvTrac         int           `json:"MvTrac"`
			// MvType         int           `json:"MvType"`
			// OldCpy         int           `json:"OldCpy"`
			// OriOtherName   string        `json:"OriOtherName"`
			// OriSongName    string        `json:"OriSongName"`
			// OtherName      string        `json:"OtherName"`
			// OwnerCount     int           `json:"OwnerCount"`
			// PayType        int           `json:"PayType"`
			// PkgPrice       int           `json:"PkgPrice"`
			// Price          int           `json:"Price"`
			// Privilege      int           `json:"Privilege"`
			// Publish        int           `json:"Publish"`
			// PublishAge     int           `json:"PublishAge"`
			// PublishTime    string        `json:"PublishTime"`
			// QualityLevel   int           `json:"QualityLevel"`
			// RankID         int           `json:"RankId"`
			// Res            struct {
			// 	FailProcess int `json:"FailProcess"`
			// 	PayType     int `json:"PayType"`
			// 	PkgPrice    int `json:"PkgPrice"`
			// 	Price       int `json:"Price"`
			// 	Privilege   int `json:"Privilege"`
			// } `json:"Res"`
			// ResBitrate    int    `json:"ResBitrate"`
			// ResDuration   int    `json:"ResDuration"`
			// ResFileHash   string `json:"ResFileHash"`
			// ResFileSize   int    `json:"ResFileSize"`
			// SQBitrate     int    `json:"SQBitrate"`
			// SQDuration    int    `json:"SQDuration"`
			// SQExtName     string `json:"SQExtName"`
			// SQFailProcess int    `json:"SQFailProcess"`
			// SQFileHash    string `json:"SQFileHash"`
			// SQFileSize    int    `json:"SQFileSize"`
			// SQPayType     int    `json:"SQPayType"`
			// SQPkgPrice    int    `json:"SQPkgPrice"`
			// SQPrice       int    `json:"SQPrice"`
			// SQPrivilege   int    `json:"SQPrivilege"`
			// Scid          int    `json:"Scid"`
			// ShowingFlag   int    `json:"ShowingFlag"`
			// SingerID      []int  `json:"SingerId"`
			// SingerName    string `json:"SingerName"`
			// Singers       []struct {
			// 	ID   int    `json:"id"`
			// 	IPID int    `json:"ip_id"`
			// 	Name string `json:"name"`
			// } `json:"Singers"`
			// SongLabel     string `json:"SongLabel"`
			// SongName      string `json:"SongName"`
			// Source        string `json:"Source"`
			// SourceID      int    `json:"SourceID"`
			// Suffix        string `json:"Suffix"`
			// SuperBitrate  int    `json:"SuperBitrate"`
			// SuperDuration int    `json:"SuperDuration"`
			// SuperExtName  string `json:"SuperExtName"`
			// SuperFileHash string `json:"SuperFileHash"`
			// SuperFileSize int    `json:"SuperFileSize"`
			// TagContent    string `json:"TagContent"`
			// TagDetails    []struct {
			// 	Content string `json:"content"`
			// 	Rankid  int    `json:"rankid"`
			// 	Type    int    `json:"type"`
			// 	Version int    `json:"version"`
			// } `json:"TagDetails"`
			// TopID           int    `json:"TopID"`
			// TopicRemark     string `json:"TopicRemark"`
			// TopicURL        string `json:"TopicUrl"`
			// Type            string `json:"Type"`
			// Uploader        string `json:"Uploader"`
			// UploaderContent string `json:"UploaderContent"`
			// MvTotal         int    `json:"mvTotal"`
			// Mvdata          []struct {
			// 	Hash string `json:"hash"`
			// 	ID   string `json:"id"`
			// 	Trk  string `json:"trk"`
			// 	Typ  int    `json:"typ"`
			// } `json:"mvdata"`
			// RecommendType int `json:"recommend_type"`
			// TransParam    struct {
			// 	AppidBlock string `json:"appid_block"`
			// 	Cid        int    `json:"cid"`
			// 	Classmap   struct {
			// 		Attr0 int `json:"attr0"`
			// 	} `json:"classmap"`
			// 	CpyAttr0       int    `json:"cpy_attr0"`
			// 	CpyGrade       int    `json:"cpy_grade"`
			// 	CpyLevel       int    `json:"cpy_level"`
			// 	Display        int    `json:"display"`
			// 	DisplayRate    int    `json:"display_rate"`
			// 	HashMultitrack string `json:"hash_multitrack"`
			// 	Ipmap          struct {
			// 		Attr0 int `json:"attr0"`
			// 	} `json:"ipmap"`
			// 	Language         string `json:"language"`
			// 	MusicpackAdvance int    `json:"musicpack_advance"`
			// 	PayBlockTpl      int    `json:"pay_block_tpl"`
			// 	Qualitymap       struct {
			// 		Attr0 int `json:"attr0"`
			// 	} `json:"qualitymap"`
			// 	SongnameSuffix string `json:"songname_suffix"`
			// } `json:"trans_param"`
			// Vvid string `json:"vvid"`
		} `json:"lists"`
		// Page       int `json:"page"`
		// Pagesize   int `json:"pagesize"`
		// Searchfull int `json:"searchfull"`
		// SecAggre   struct {
		// } `json:"sec_aggre"`
		// SecAggreV2 []interface{} `json:"sec_aggre_v2"`
		// SectagInfo struct {
		// 	IsSectag int `json:"is_sectag"`
		// } `json:"sectag_info"`
		// Size        int `json:"size"`
		// Subjecttype int `json:"subjecttype"`
		Total int `json:"total"`
	} `json:"data"`
	// ErrorCode int    `json:"error_code"`
	ErrorMsg string `json:"error_msg"`
	Status   int    `json:"status"`
}

type refreshInfo struct {
	// Data      string `json:"data"`
	ErrorCode int    `json:"error_code"`
	ErrorMsg  string `json:"error_msg"`
	Status    int    `json:"status"`
}
