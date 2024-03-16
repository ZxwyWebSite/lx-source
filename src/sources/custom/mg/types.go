package mg

type playInfo struct {
	Code string `json:"code"`
	Data struct {
		AudioFormatType string `json:"audioFormatType"`
		FreeListenType  string `json:"freeListenType"`
		HaveVisualMv    bool   `json:"haveVisualMv"`
		LrcURL          string `json:"lrcUrl"`
		MrcURL          string `json:"mrcUrl"`
		Song            struct {
			Album           string   `json:"album"`
			AlbumID         string   `json:"albumId"`
			AlbumPinyin     string   `json:"albumPinyin"`
			ContentID       string   `json:"contentId"`
			CopyrightID     string   `json:"copyrightId"`
			CopyrightType   int      `json:"copyrightType"`
			Duration        int      `json:"duration"`
			ForeverListen   bool     `json:"foreverListen"`
			HaveShockRing   int      `json:"haveShockRing"`
			Img1            string   `json:"img1"`
			Img2            string   `json:"img2"`
			Img3            string   `json:"img3"`
			MvCopyrightType int      `json:"mvCopyrightType"`
			ResourceType    string   `json:"resourceType"`
			RestrictType    int      `json:"restrictType"`
			RingToneID      string   `json:"ringToneId"`
			ShowTags        []string `json:"showTags"`
			SingerList      []struct {
				ID           string `json:"id"`
				Img          string `json:"img"`
				Name         string `json:"name"`
				NameSpelling string `json:"nameSpelling"`
			} `json:"singerList"`
			SongID     string `json:"songId"`
			SongName   string `json:"songName"`
			SongPinyin string `json:"songPinyin"`
		} `json:"song"`
		TrcURL  string `json:"trcUrl"`
		URL     string `json:"url"`
		Version string `json:"version"`
	} `json:"data"`
	Info string `json:"info"`
}

// func (resp *playInfo) Url(rquality string) (ourl, emsg string) {
// 	if resp.Code != `000000` {
// 		emsg = resp.Info
// 		return
// 	}
// 	if resp.Data.URL == `` {
// 		emsg = `No Data: 无返回链接`
// 		return
// 	}
// 	if resp.Data.AudioFormatType != rquality {
// 		emsg = ztool.Str_FastConcat(`实际音质不匹配: `, rquality, ` <= `, resp.Data.AudioFormatType)
// 		if !env.Config.Source.ForceFallback {
// 			return
// 		}
// 	}
// 	ourl = utils.DelQuery(resp.Data.URL)
// 	return
// }

type mgApi_Song struct {
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

// func (resp *mgApi_Song) Url(string) (ourl, emsg string) {
// 	if resp.Data.PlayURL != `` {
// 		ourl = `https:` + utils.DelQuery(resp.Data.PlayURL)
// 	} else {
// 		emsg = ztool.Str_FastConcat(resp.Code, `: `, resp.Msg)
// 	}
// 	return
// }

type musicInfo struct {
	// Code     string `json:"code"`
	// Info     string `json:"info"`
	Resource []struct {
		// ResourceType string `json:"resourceType"`
		// RefID        string `json:"refId"`
		// CopyrightID  string `json:"copyrightId"`
		// ContentID    string `json:"contentId"`
		// SongID       string `json:"songId"`
		// SongName     string `json:"songName"`
		// SingerID     string `json:"singerId"`
		// Singer       string `json:"singer"`
		AlbumID string `json:"albumId"`
		// Album        string `json:"album"`
		// AlbumImgs    []struct {
		// 	ImgSizeType string `json:"imgSizeType"`
		// 	Img         string `json:"img"`
		// 	FileID      string `json:"fileId"`
		// 	WebpImg     string `json:"webpImg"`
		// } `json:"albumImgs"`
		// OpNumItem struct {
		// 	PlayNum                   int    `json:"playNum"`
		// 	PlayNumDesc               string `json:"playNumDesc"`
		// 	KeepNum                   int    `json:"keepNum"`
		// 	KeepNumDesc               string `json:"keepNumDesc"`
		// 	CommentNum                int    `json:"commentNum"`
		// 	CommentNumDesc            string `json:"commentNumDesc"`
		// 	ShareNum                  int    `json:"shareNum"`
		// 	ShareNumDesc              string `json:"shareNumDesc"`
		// 	OrderNumByWeek            int    `json:"orderNumByWeek"`
		// 	OrderNumByWeekDesc        string `json:"orderNumByWeekDesc"`
		// 	OrderNumByTotal           int    `json:"orderNumByTotal"`
		// 	OrderNumByTotalDesc       string `json:"orderNumByTotalDesc"`
		// 	ThumbNum                  int    `json:"thumbNum"`
		// 	ThumbNumDesc              string `json:"thumbNumDesc"`
		// 	FollowNum                 int    `json:"followNum"`
		// 	FollowNumDesc             string `json:"followNumDesc"`
		// 	SubscribeNum              int    `json:"subscribeNum"`
		// 	SubscribeNumDesc          string `json:"subscribeNumDesc"`
		// 	LivePlayNum               int    `json:"livePlayNum"`
		// 	LivePlayNumDesc           string `json:"livePlayNumDesc"`
		// 	PopularNum                int    `json:"popularNum"`
		// 	PopularNumDesc            string `json:"popularNumDesc"`
		// 	BookingNum                int    `json:"bookingNum"`
		// 	BookingNumDesc            string `json:"bookingNumDesc"`
		// 	SettingNum                int    `json:"settingNum"`
		// 	SettingNumDesc            string `json:"settingNumDesc"`
		// 	CallNum                   int    `json:"callNum"`
		// 	CallNumDesc               string `json:"callNumDesc"`
		// 	CallingPlayNum            int    `json:"callingPlayNum"`
		// 	CallingPlayNumDesc        string `json:"callingPlayNumDesc"`
		// 	CallingPlayDuration       int    `json:"callingPlayDuration"`
		// 	CallingPlayDurationDesc   string `json:"callingPlayDurationDesc"`
		// 	CalledPlayDuration        int    `json:"calledPlayDuration"`
		// 	CalledPlayDurationDesc    string `json:"calledPlayDurationDesc"`
		// 	RingtoneAppPlayNum        int    `json:"ringtoneAppPlayNum"`
		// 	RingtoneAppPlayNumDesc    string `json:"ringtoneAppPlayNumDesc"`
		// 	RingtoneAppSettingNum     int    `json:"ringtoneAppSettingNum"`
		// 	RingtoneAppSettingNumDesc string `json:"ringtoneAppSettingNumDesc"`
		// } `json:"opNumItem"`
		// ToneControl  string `json:"toneControl"`
		// RelatedSongs []struct {
		// 	ResourceType     string `json:"resourceType"`
		// 	ResourceTypeName string `json:"resourceTypeName"`
		// 	CopyrightID      string `json:"copyrightId"`
		// 	ProductID        string `json:"productId"`
		// } `json:"relatedSongs"`
		// RateFormats []struct {
		// 	ResourceType         string `json:"resourceType"`
		// 	FormatType           string `json:"formatType"`
		// 	Format               string `json:"format"`
		// 	Size                 string `json:"size"`
		// 	FileType             string `json:"fileType,omitempty"`
		// 	Price                string `json:"price"`
		// 	AndroidFileType      string `json:"androidFileType,omitempty"`
		// 	IosFileType          string `json:"iosFileType,omitempty"`
		// 	IosSize              string `json:"iosSize,omitempty"`
		// 	AndroidSize          string `json:"androidSize,omitempty"`
		// 	IosFormat            string `json:"iosFormat,omitempty"`
		// 	AndroidFormat        string `json:"androidFormat,omitempty"`
		// 	IosAccuracyLevel     string `json:"iosAccuracyLevel,omitempty"`
		// 	AndroidAccuracyLevel string `json:"androidAccuracyLevel,omitempty"`
		// } `json:"rateFormats"`
		NewRateFormats []struct {
			// ResourceType         string `json:"resourceType"`
			FormatType string `json:"formatType"`
			// Format               string `json:"format"`
			// Size                 string `json:"size"`
			// FileType             string `json:"fileType,omitempty"`
			// Price                string `json:"price"`
			// AndroidFileType      string `json:"androidFileType,omitempty"`
			// IosFileType          string `json:"iosFileType,omitempty"`
			// IosSize              string `json:"iosSize,omitempty"`
			// AndroidSize          string `json:"androidSize,omitempty"`
			// IosFormat            string `json:"iosFormat,omitempty"`
			// AndroidFormat        string `json:"androidFormat,omitempty"`
			// IosAccuracyLevel     string `json:"iosAccuracyLevel,omitempty"`
			// AndroidAccuracyLevel string `json:"androidAccuracyLevel,omitempty"`
		} `json:"newRateFormats"`
		// LrcURL  string `json:"lrcUrl"`
		// TagList []struct {
		// 	ResourceType string `json:"resourceType"`
		// 	TagID        string `json:"tagId"`
		// 	TagName      string `json:"tagName"`
		// 	TagDesc      string `json:"tagDesc,omitempty"`
		// } `json:"tagList"`
		// DigitalColumnID    string `json:"digitalColumnId"`
		// Copyright          string `json:"copyright"`
		// ValidStatus        bool   `json:"validStatus"`
		// SongDescs          string `json:"songDescs"`
		// SongAliasName      string `json:"songAliasName"`
		// IsInDAlbum         string `json:"isInDAlbum"`
		// IsInSideDalbum     string `json:"isInSideDalbum"`
		// IsInSalesPeriod    string `json:"isInSalesPeriod"`
		// SongType           string `json:"songType"`
		// MrcURL             string `json:"mrcUrl"`
		// InvalidateDate     string `json:"invalidateDate"`
		// DalbumID           string `json:"dalbumId"`
		// TrcURL             string `json:"trcUrl"`
		// VipType            string `json:"vipType"`
		// ScopeOfcopyright   string `json:"scopeOfcopyright"`
		// AuditionsType      string `json:"auditionsType"`
		// FirstIcon          string `json:"firstIcon"`
		// TranslateName      string `json:"translateName"`
		// ChargeAuditions    string `json:"chargeAuditions"`
		// OldChargeAuditions string `json:"oldChargeAuditions"`
		// SongIcon           string `json:"songIcon"`
		// CodeRate           struct {
		// 	PQ struct {
		// 		CodeRateChargeAuditions string `json:"codeRateChargeAuditions"`
		// 		IsCodeRateDownload      string `json:"isCodeRateDownload"`
		// 		CodeRateFileSize        string `json:"codeRateFileSize"`
		// 	} `json:"PQ"`
		// 	HQ struct {
		// 		CodeRateChargeAuditions string `json:"codeRateChargeAuditions"`
		// 		IsCodeRateDownload      string `json:"isCodeRateDownload"`
		// 	} `json:"HQ"`
		// 	SQ struct {
		// 		CodeRateChargeAuditions string `json:"codeRateChargeAuditions"`
		// 		IsCodeRateDownload      string `json:"isCodeRateDownload"`
		// 		ContentIDSQ             string `json:"contentIdSQ"`
		// 	} `json:"SQ"`
		// } `json:"codeRate"`
		// IsDownload    string `json:"isDownload"`
		// CopyrightType string `json:"copyrightType"`
		// HasMv         string `json:"hasMv"`
		// TopQuality    string `json:"topQuality"`
		// PreSale       string `json:"preSale"`
		// IsShare       string `json:"isShare"`
		// IsCollection  string `json:"isCollection"`
		// Length        string `json:"length"`
		// SingerImg     struct {
		// 	Num421 struct {
		// 		SingerName   string `json:"singerName"`
		// 		MiguImgItems []struct {
		// 			ImgSizeType string `json:"imgSizeType"`
		// 			Img         string `json:"img"`
		// 			FileID      string `json:"fileId"`
		// 			WebpImg     string `json:"webpImg"`
		// 		} `json:"miguImgItems"`
		// 	} `json:"421"`
		// } `json:"singerImg"`
		// SongNamePinyin  string `json:"songNamePinyin"`
		// AlbumNamePinyin string `json:"albumNamePinyin"`
		// Artists         []struct {
		// 	ID           string `json:"id"`
		// 	Name         string `json:"name"`
		// 	NameSpelling string `json:"nameSpelling"`
		// } `json:"artists"`
		// LandscapImg         string   `json:"landscapImg"`
		// VipLogo             string   `json:"vipLogo"`
		// VipDownload         string   `json:"vipDownload"`
		// FirstPublish        string   `json:"firstPublish"`
		// ShowTag             []string `json:"showTag"`
		// MaterialValidStatus bool     `json:"materialValidStatus"`
		// NeedEncrypt         string   `json:"needEncrypt"`
		// ForeverListen       bool     `json:"foreverListen"`
		// HasAssociatedRing   bool     `json:"hasAssociatedRing"`
	} `json:"resource"`
}

type albumInfo struct {
	// Code     string `json:"code"`
	// Info     string `json:"info"`
	Resource []struct {
		// ResourceType string `json:"resourceType"`
		// AlbumID      string `json:"albumId"`
		// ImgItems     []struct {
		// 	ImgSizeType string `json:"imgSizeType"`
		// 	Img         string `json:"img"`
		// 	FileID      string `json:"fileId"`
		// 	WebpImg     string `json:"webpImg"`
		// } `json:"imgItems"`
		// Title      string `json:"title"`
		// Singer     string `json:"singer"`
		// SingerID   string `json:"singerId"`
		// SingerImgs []struct {
		// 	ImgSizeType string `json:"imgSizeType"`
		// 	Img         string `json:"img"`
		// 	FileID      string `json:"fileId"`
		// 	WebpImg     string `json:"webpImg"`
		// } `json:"singerImgs"`
		// Summary     string `json:"summary"`
		// TotalCount  string `json:"totalCount"`
		// PublishTime string `json:"publishTime"`
		// PublishCorp string `json:"publishCorp"`
		// OpNumItem   struct {
		// 	PlayNum                   int    `json:"playNum"`
		// 	PlayNumDesc               string `json:"playNumDesc"`
		// 	KeepNum                   int    `json:"keepNum"`
		// 	KeepNumDesc               string `json:"keepNumDesc"`
		// 	CommentNum                int    `json:"commentNum"`
		// 	CommentNumDesc            string `json:"commentNumDesc"`
		// 	ShareNum                  int    `json:"shareNum"`
		// 	ShareNumDesc              string `json:"shareNumDesc"`
		// 	OrderNumByWeek            int    `json:"orderNumByWeek"`
		// 	OrderNumByWeekDesc        string `json:"orderNumByWeekDesc"`
		// 	OrderNumByTotal           int    `json:"orderNumByTotal"`
		// 	OrderNumByTotalDesc       string `json:"orderNumByTotalDesc"`
		// 	ThumbNum                  int    `json:"thumbNum"`
		// 	ThumbNumDesc              string `json:"thumbNumDesc"`
		// 	FollowNum                 int    `json:"followNum"`
		// 	FollowNumDesc             string `json:"followNumDesc"`
		// 	SubscribeNum              int    `json:"subscribeNum"`
		// 	SubscribeNumDesc          string `json:"subscribeNumDesc"`
		// 	LivePlayNum               int    `json:"livePlayNum"`
		// 	LivePlayNumDesc           string `json:"livePlayNumDesc"`
		// 	PopularNum                int    `json:"popularNum"`
		// 	PopularNumDesc            string `json:"popularNumDesc"`
		// 	BookingNum                int    `json:"bookingNum"`
		// 	BookingNumDesc            string `json:"bookingNumDesc"`
		// 	SettingNum                int    `json:"settingNum"`
		// 	SettingNumDesc            string `json:"settingNumDesc"`
		// 	CallNum                   int    `json:"callNum"`
		// 	CallNumDesc               string `json:"callNumDesc"`
		// 	CallingPlayNum            int    `json:"callingPlayNum"`
		// 	CallingPlayNumDesc        string `json:"callingPlayNumDesc"`
		// 	CallingPlayDuration       int    `json:"callingPlayDuration"`
		// 	CallingPlayDurationDesc   string `json:"callingPlayDurationDesc"`
		// 	CalledPlayDuration        int    `json:"calledPlayDuration"`
		// 	CalledPlayDurationDesc    string `json:"calledPlayDurationDesc"`
		// 	RingtoneAppPlayNum        int    `json:"ringtoneAppPlayNum"`
		// 	RingtoneAppPlayNumDesc    string `json:"ringtoneAppPlayNumDesc"`
		// 	RingtoneAppSettingNum     int    `json:"ringtoneAppSettingNum"`
		// 	RingtoneAppSettingNumDesc string `json:"ringtoneAppSettingNumDesc"`
		// } `json:"opNumItem"`
		// Tags []struct {
		// 	ResourceType string `json:"resourceType"`
		// 	TagID        string `json:"tagId"`
		// 	TagName      string `json:"tagName"`
		// 	TagDesc      string `json:"tagDesc,omitempty"`
		// } `json:"tags"`
		// AlbumAliasName string `json:"albumAliasName"`
		// AlbumClass     string `json:"albumClass"`
		// Language       string `json:"language"`
		// PublishCompany string `json:"publishCompany"`
		// PublishDate    string `json:"publishDate"`
		// TranslateName  string `json:"translateName"`
		SongItems []struct {
			// ResourceType string `json:"resourceType"`
			// RefID        string `json:"refId"`
			CopyrightID string `json:"copyrightId"`
			// ContentID    string `json:"contentId"`
			// SongID       string `json:"songId"`
			// SongName     string `json:"songName"`
			// SingerID     string `json:"singerId"`
			// Singer       string `json:"singer"`
			// AlbumID      string `json:"albumId"`
			// Album        string `json:"album"`
			// AlbumImgs    []struct {
			// 	ImgSizeType string `json:"imgSizeType"`
			// 	Img         string `json:"img"`
			// 	FileID      string `json:"fileId"`
			// 	WebpImg     string `json:"webpImg"`
			// } `json:"albumImgs"`
			// OpNumItem struct {
			// 	PlayNum                   int    `json:"playNum"`
			// 	PlayNumDesc               string `json:"playNumDesc"`
			// 	KeepNum                   int    `json:"keepNum"`
			// 	KeepNumDesc               string `json:"keepNumDesc"`
			// 	CommentNum                int    `json:"commentNum"`
			// 	CommentNumDesc            string `json:"commentNumDesc"`
			// 	ShareNum                  int    `json:"shareNum"`
			// 	ShareNumDesc              string `json:"shareNumDesc"`
			// 	OrderNumByWeek            int    `json:"orderNumByWeek"`
			// 	OrderNumByWeekDesc        string `json:"orderNumByWeekDesc"`
			// 	OrderNumByTotal           int    `json:"orderNumByTotal"`
			// 	OrderNumByTotalDesc       string `json:"orderNumByTotalDesc"`
			// 	ThumbNum                  int    `json:"thumbNum"`
			// 	ThumbNumDesc              string `json:"thumbNumDesc"`
			// 	FollowNum                 int    `json:"followNum"`
			// 	FollowNumDesc             string `json:"followNumDesc"`
			// 	SubscribeNum              int    `json:"subscribeNum"`
			// 	SubscribeNumDesc          string `json:"subscribeNumDesc"`
			// 	LivePlayNum               int    `json:"livePlayNum"`
			// 	LivePlayNumDesc           string `json:"livePlayNumDesc"`
			// 	PopularNum                int    `json:"popularNum"`
			// 	PopularNumDesc            string `json:"popularNumDesc"`
			// 	BookingNum                int    `json:"bookingNum"`
			// 	BookingNumDesc            string `json:"bookingNumDesc"`
			// 	SettingNum                int    `json:"settingNum"`
			// 	SettingNumDesc            string `json:"settingNumDesc"`
			// 	CallNum                   int    `json:"callNum"`
			// 	CallNumDesc               string `json:"callNumDesc"`
			// 	CallingPlayNum            int    `json:"callingPlayNum"`
			// 	CallingPlayNumDesc        string `json:"callingPlayNumDesc"`
			// 	CallingPlayDuration       int    `json:"callingPlayDuration"`
			// 	CallingPlayDurationDesc   string `json:"callingPlayDurationDesc"`
			// 	CalledPlayDuration        int    `json:"calledPlayDuration"`
			// 	CalledPlayDurationDesc    string `json:"calledPlayDurationDesc"`
			// 	RingtoneAppPlayNum        int    `json:"ringtoneAppPlayNum"`
			// 	RingtoneAppPlayNumDesc    string `json:"ringtoneAppPlayNumDesc"`
			// 	RingtoneAppSettingNum     int    `json:"ringtoneAppSettingNum"`
			// 	RingtoneAppSettingNumDesc string `json:"ringtoneAppSettingNumDesc"`
			// } `json:"opNumItem"`
			// ToneControl  string `json:"toneControl"`
			// RelatedSongs []struct {
			// 	ResourceType     string `json:"resourceType"`
			// 	ResourceTypeName string `json:"resourceTypeName"`
			// 	CopyrightID      string `json:"copyrightId"`
			// 	ProductID        string `json:"productId"`
			// } `json:"relatedSongs"`
			// RateFormats []struct {
			// 	ResourceType         string `json:"resourceType"`
			// 	FormatType           string `json:"formatType"`
			// 	URL                  string `json:"url,omitempty"`
			// 	Format               string `json:"format"`
			// 	Size                 string `json:"size"`
			// 	FileType             string `json:"fileType,omitempty"`
			// 	Price                string `json:"price"`
			// 	IosURL               string `json:"iosUrl,omitempty"`
			// 	AndroidURL           string `json:"androidUrl,omitempty"`
			// 	AndroidFileType      string `json:"androidFileType,omitempty"`
			// 	IosFileType          string `json:"iosFileType,omitempty"`
			// 	IosSize              string `json:"iosSize,omitempty"`
			// 	AndroidSize          string `json:"androidSize,omitempty"`
			// 	IosFormat            string `json:"iosFormat,omitempty"`
			// 	AndroidFormat        string `json:"androidFormat,omitempty"`
			// 	IosAccuracyLevel     string `json:"iosAccuracyLevel,omitempty"`
			// 	AndroidAccuracyLevel string `json:"androidAccuracyLevel,omitempty"`
			// } `json:"rateFormats"`
			NewRateFormats []struct {
				// ResourceType         string `json:"resourceType"`
				FormatType string `json:"formatType"`
				URL        string `json:"url,omitempty"`
				// Format               string `json:"format"`
				// Size                 string `json:"size"`
				// FileType             string `json:"fileType,omitempty"`
				// Price                string `json:"price"`
				// IosURL               string `json:"iosUrl,omitempty"`
				AndroidURL string `json:"androidUrl,omitempty"`
				// AndroidFileType      string `json:"androidFileType,omitempty"`
				// IosFileType          string `json:"iosFileType,omitempty"`
				// IosSize              string `json:"iosSize,omitempty"`
				// AndroidSize          string `json:"androidSize,omitempty"`
				// IosFormat            string `json:"iosFormat,omitempty"`
				// AndroidFormat        string `json:"androidFormat,omitempty"`
				// IosAccuracyLevel     string `json:"iosAccuracyLevel,omitempty"`
				// AndroidAccuracyLevel string `json:"androidAccuracyLevel,omitempty"`
			} `json:"newRateFormats"`
			// LrcURL  string `json:"lrcUrl"`
			// TagList []struct {
			// 	ResourceType string `json:"resourceType"`
			// 	TagID        string `json:"tagId"`
			// 	TagName      string `json:"tagName"`
			// 	TagDesc      string `json:"tagDesc,omitempty"`
			// } `json:"tagList"`
			// DigitalColumnID    string `json:"digitalColumnId"`
			// Copyright          string `json:"copyright"`
			// ValidStatus        bool   `json:"validStatus"`
			// SongDescs          string `json:"songDescs"`
			// SongAliasName      string `json:"songAliasName"`
			// IsInDAlbum         string `json:"isInDAlbum"`
			// IsInSideDalbum     string `json:"isInSideDalbum"`
			// IsInSalesPeriod    string `json:"isInSalesPeriod"`
			// SongType           string `json:"songType"`
			// MrcURL             string `json:"mrcUrl"`
			// InvalidateDate     string `json:"invalidateDate"`
			// DalbumID           string `json:"dalbumId"`
			// TrackNumber        string `json:"trackNumber"`
			// TrcURL             string `json:"trcUrl"`
			// Disc               string `json:"disc"`
			// VipType            string `json:"vipType"`
			// ScopeOfcopyright   string `json:"scopeOfcopyright"`
			// AuditionsType      string `json:"auditionsType"`
			// FirstIcon          string `json:"firstIcon"`
			// TranslateName      string `json:"translateName,omitempty"`
			// ChargeAuditions    string `json:"chargeAuditions"`
			// OldChargeAuditions string `json:"oldChargeAuditions"`
			// SongIcon           string `json:"songIcon"`
			// CodeRate           struct {
			// 	PQ struct {
			// 		CodeRateChargeAuditions string `json:"codeRateChargeAuditions"`
			// 		IsCodeRateDownload      string `json:"isCodeRateDownload"`
			// 		CodeRateFileSize        string `json:"codeRateFileSize"`
			// 	} `json:"PQ"`
			// 	HQ struct {
			// 		CodeRateChargeAuditions string `json:"codeRateChargeAuditions"`
			// 		IsCodeRateDownload      string `json:"isCodeRateDownload"`
			// 	} `json:"HQ"`
			// 	SQ struct {
			// 		CodeRateChargeAuditions string `json:"codeRateChargeAuditions"`
			// 		IsCodeRateDownload      string `json:"isCodeRateDownload"`
			// 		ContentIDSQ             string `json:"contentIdSQ"`
			// 	} `json:"SQ"`
			// } `json:"codeRate"`
			// IsDownload    string `json:"isDownload"`
			// CopyrightType string `json:"copyrightType"`
			// HasMv         string `json:"hasMv"`
			// TopQuality    string `json:"topQuality"`
			// PreSale       string `json:"preSale"`
			// IsShare       string `json:"isShare"`
			// IsCollection  string `json:"isCollection"`
			// Length        string `json:"length"`
			// SingerImg     struct {
			// 	Num421 struct {
			// 		SingerName   string `json:"singerName"`
			// 		MiguImgItems []struct {
			// 			ImgSizeType string `json:"imgSizeType"`
			// 			Img         string `json:"img"`
			// 			FileID      string `json:"fileId"`
			// 			WebpImg     string `json:"webpImg"`
			// 		} `json:"miguImgItems"`
			// 	} `json:"421"`
			// } `json:"singerImg"`
			// SongNamePinyin  string `json:"songNamePinyin"`
			// AlbumNamePinyin string `json:"albumNamePinyin"`
			// Artists         []struct {
			// 	ID           string `json:"id"`
			// 	Name         string `json:"name"`
			// 	NameSpelling string `json:"nameSpelling"`
			// } `json:"artists"`
			// LandscapImg         string   `json:"landscapImg"`
			// VipLogo             string   `json:"vipLogo"`
			// VipDownload         string   `json:"vipDownload"`
			// FirstPublish        string   `json:"firstPublish"`
			// ShowTag             []string `json:"showTag"`
			// MaterialValidStatus bool     `json:"materialValidStatus"`
			// NeedEncrypt         string   `json:"needEncrypt"`
			// ForeverListen       bool     `json:"foreverListen"`
			// HasAssociatedRing   bool     `json:"hasAssociatedRing"`
			// LoginListenFlag     string   `json:"loginListenFlag,omitempty"`
			// MvCopyright         string   `json:"mvCopyright,omitempty"`
			// ForeverListenFlag   string   `json:"foreverListenFlag,omitempty"`
		} `json:"songItems"`
	} `json:"resource"`
}
