package kg

type (
	musicInfo struct {
		AlbumAudioID     string `json:"album_audio_id"`
		AudioID          string `json:"audio_id"`
		AuthorName       string `json:"author_name"`
		OriAudioName     string `json:"ori_audio_name"`
		OfficialSongname string `json:"official_songname"`
		Songname         string `json:"songname"`
		Remark           string `json:"remark"`
		SuffixAudioName  string `json:"suffix_audio_name"`
		IsSearch         string `json:"is_search"`
		IsOriginal       string `json:"is_original"`
		IsPublish        string `json:"is_publish"`
		MixsongType      string `json:"mixsong_type"`
		Version          string `json:"version"`
		Language         string `json:"language"`
		Bpm              string `json:"bpm"`
		BpmType          string `json:"bpm_type"`
		BpmDesc          string `json:"bpm_desc"`
		PublishDate      string `json:"publish_date"`
		Extname          string `json:"extname"`
		AlbumInfo        struct {
			AlbumID      string `json:"album_id"`
			AlbumName    string `json:"album_name"`
			PublishDate  string `json:"publish_date"`
			Category     string `json:"category"`
			IsPublish    string `json:"is_publish"`
			SizableCover string `json:"sizable_cover"`
		} `json:"album_info"`
		Classification []struct {
			Usage   string `json:"usage"`
			Status  string `json:"status"`
			Level   string `json:"level"`
			ResType string `json:"res_type"`
			Type    string `json:"type"`
			ID      string `json:"id"`
			ResID   string `json:"res_id"`
		} `json:"classification"`
		AudioInfo struct {
			IsFileHead             string `json:"is_file_head"`
			IsFileHead320          string `json:"is_file_head_320"`
			AudioID                string `json:"audio_id"`
			Hash                   string `json:"hash"`
			Filesize               string `json:"filesize"`
			Timelength             string `json:"timelength"`
			Bitrate                string `json:"bitrate"`
			Hash128                string `json:"hash_128"`
			Filesize128            string `json:"filesize_128"`
			Timelength128          string `json:"timelength_128"`
			Hash320                string `json:"hash_320"`
			Filesize320            string `json:"filesize_320"`
			Timelength320          string `json:"timelength_320"`
			HashFlac               string `json:"hash_flac"`
			FilesizeFlac           string `json:"filesize_flac"`
			TimelengthFlac         string `json:"timelength_flac"`
			BitrateFlac            string `json:"bitrate_flac"`
			HashHigh               string `json:"hash_high"`
			FilesizeHigh           string `json:"filesize_high"`
			TimelengthHigh         string `json:"timelength_high"`
			BitrateHigh            string `json:"bitrate_high"`
			HashSuper              string `json:"hash_super"`
			FilesizeSuper          string `json:"filesize_super"`
			TimelengthSuper        string `json:"timelength_super"`
			BitrateSuper           string `json:"bitrate_super"`
			HashVinylrecord        string `json:"hash_vinylrecord"`
			FilesizeVinylrecord    string `json:"filesize_vinylrecord"`
			TimelengthVinylrecord  string `json:"timelength_vinylrecord"`
			BitrateVinylrecord     string `json:"bitrate_vinylrecord"`
			HashMultichannel       string `json:"hash_multichannel"`
			FilesizeMultichannel   string `json:"filesize_multichannel"`
			TimelengthMultichannel string `json:"timelength_multichannel"`
			BitrateMultichannel    string `json:"bitrate_multichannel"`
			HashDolby448           string `json:"hash_dolby_448"`
			FilesizeDolby448       string `json:"filesize_dolby_448"`
			TimelengthDolby448     string `json:"timelength_dolby_448"`
			BitrateDolby448        string `json:"bitrate_dolby_448"`
			HashDolby640           string `json:"hash_dolby_640"`
			FilesizeDolby640       string `json:"filesize_dolby_640"`
			TimelengthDolby640     string `json:"timelength_dolby_640"`
			BitrateDolby640        string `json:"bitrate_dolby_640"`
			HashDolby768           string `json:"hash_dolby_768"`
			FilesizeDolby768       string `json:"filesize_dolby_768"`
			TimelengthDolby768     string `json:"timelength_dolby_768"`
			BitrateDolby768        string `json:"bitrate_dolby_768"`
			AudioGroupID           string `json:"audio_group_id"`
			ExtnameSuper           string `json:"extname_super"`
			Extname                string `json:"extname"`
			FailProcess            int    `json:"fail_process"`
			PayType                int    `json:"pay_type"`
			Type                   string `json:"type"`
			OldCpy                 int    `json:"old_cpy"`
			Privilege              string `json:"privilege"`
			Privilege128           string `json:"privilege_128"`
			Privilege320           string `json:"privilege_320"`
			PrivilegeFlac          string `json:"privilege_flac"`
			PrivilegeHigh          string `json:"privilege_high"`
			PrivilegeSuper         string `json:"privilege_super"`
			PrivilegeVinylrecord   string `json:"privilege_vinylrecord"`
			PrivilegeMultichannel  string `json:"privilege_multichannel"`
			PrivilegeDolby448      string `json:"privilege_dolby_448"`
			PrivilegeDolby640      string `json:"privilege_dolby_640"`
			PrivilegeDolby768      string `json:"privilege_dolby_768"`
			TransParam             struct {
				HashOffset struct {
					StartByte  int    `json:"start_byte"`
					EndByte    int    `json:"end_byte"`
					StartMs    int    `json:"start_ms"`
					EndMs      int    `json:"end_ms"`
					OffsetHash string `json:"offset_hash"`
					FileType   int    `json:"file_type"`
					ClipHash   string `json:"clip_hash"`
				} `json:"hash_offset"`
				MusicpackAdvance int `json:"musicpack_advance"`
				PayBlockTpl      int `json:"pay_block_tpl"`
				Display          int `json:"display"`
				DisplayRate      int `json:"display_rate"`
				CpyGrade         int `json:"cpy_grade"`
				CpyLevel         int `json:"cpy_level"`
				Cid              int `json:"cid"`
				CpyAttr0         int `json:"cpy_attr0"`
				Classmap         struct {
					Attr0 int `json:"attr0"`
				} `json:"classmap"`
				InitPubDay int `json:"init_pub_day"`
				Qualitymap struct {
					Attr0 int `json:"attr0"`
				} `json:"qualitymap"`
				Language string `json:"language"`
			} `json:"trans_param"`
		} `json:"audio_info"`
	}
	playInfo struct {
		Hash     string `json:"hash"`
		Classmap struct {
			Attr0 int `json:"attr0"`
		} `json:"classmap"`
		Status      int      `json:"status"`
		Volume      float64  `json:"volume"`
		StdHashTime int      `json:"std_hash_time"`
		URL         []string `json:"url"`
		StdHash     string   `json:"std_hash"`
		TransParam  struct {
			Display     int `json:"display"`
			DisplayRate int `json:"display_rate"`
		} `json:"trans_param"`
		FileHead   int     `json:"fileHead"`
		VolumePeak float64 `json:"volume_peak"`
		BitRate    int     `json:"bitRate"`
		TimeLength int     `json:"timeLength"`
		VolumeGain int     `json:"volume_gain"`
		Q          int     `json:"q"`
		FileName   string  `json:"fileName"`
		ExtName    string  `json:"extName"`
		FileSize   int     `json:"fileSize"`
	}
)
