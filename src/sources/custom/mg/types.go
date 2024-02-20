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
