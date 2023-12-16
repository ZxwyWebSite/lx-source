package builtin

type (
	// FongerData 数据更新
	// FgData_Headers struct {
	// 	Kw struct {
	// 		Headers string `json:"headers"`
	// 	} `json:"kw"`
	// 	Mg struct {
	// 		Referer   string `json:"Referer"`
	// 		UserAgent string `json:"User-Agent"`
	// 		By        string `json:"By"`
	// 		Channel   string `json:"channel"`
	// 		Cookie    string `json:"Cookie"`
	// 	} `json:"mg"`
	// 	Wy struct {
	// 		Cookie string `json:"Cookie"`
	// 	} `json:"wy"`
	// }

	// FongerApi 方格音乐接口
	FyApi_Song struct {
		Code int `json:"code"`
		Data []struct {
			ID                 int         `json:"id"`
			URL                string      `json:"url"`
			Br                 int         `json:"br"`
			Size               int         `json:"size"`
			Md5                string      `json:"md5"`
			Code               int         `json:"code"`
			Expi               int         `json:"expi"`
			Type               string      `json:"type"`
			Gain               float64     `json:"gain"`
			Peak               float64     `json:"peak"`
			Fee                int         `json:"fee"`
			Uf                 interface{} `json:"uf"`
			Payed              int         `json:"payed"`
			Flag               int         `json:"flag"`
			CanExtend          bool        `json:"canExtend"`
			FreeTrialInfo      interface{} `json:"freeTrialInfo"`
			Level              string      `json:"level"`
			EncodeType         string      `json:"encodeType"`
			FreeTrialPrivilege struct {
				ResConsumable      bool        `json:"resConsumable"`
				UserConsumable     bool        `json:"userConsumable"`
				ListenType         interface{} `json:"listenType"`
				CannotListenReason interface{} `json:"cannotListenReason"`
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
	}
	// MiguApi 咪咕音乐接口
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
	// BodianApi 波点音乐接口
	KwApi_Song struct {
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

const (
	// FongerData
	// fgdata = `http://api.fonger.top/pc/`
	// fgdata_banner  = `banner.json` // 新闻
	// fgdata_update  = `update.json` // 更新
	// fgdata_channel = `channel.json` // 可用源
	// fgdata_headers = `headers.json` // VipCookie
	// FongerHeader
	// fyhdr_ua = `Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.84 Safari/537.36 HBPC/12.1.2.300`

	// FongerApi
	// fyapi      = `http://nm.fyapi.site/`
	// fyapi_song = `song/url/v1` // (id: 网易云ID, level: 音质[HQ: exhigh, SQ: lossless, ZQ: hires]) ?id=1885551650&level=exhigh&noCookie=true

	// Source
	s_wy = `wy`
	s_mg = `mg`
	s_kw = `kw`
	// s_kg = `kg`
	// s_tx = `tx`
)

var (
	// def_headers FgData_Headers

	// 音质列表 ( [通用音质][音乐平台]对应音质 )
	qualitys = map[string]map[string]string{
		`128k`: {
			s_wy: `standard`,
			s_mg: `1`,
			s_kw: `128k`,
		},
		`320k`: {
			s_wy: `exhigh`,
			s_mg: `2`,
			s_kw: `320k`,
		},
		`flac`: {
			s_wy: `lossless`,
			s_mg: `3`,
			s_kw: `2000k`,
		},
		`flac24bit`: {
			s_wy: `hires`,
			s_mg: `4`,
		},
	}
	// Headers
	header_mg = map[string]string{
		`Referer`: `https://m.music.migu.cn/v4/`,
		`By`:      `04f81461a98c7af557fea3cf28c4ea15`,
		`channel`: `014000D`,
		`Cookie`:  `SESSION=ZTIwODkyMDQtOTE1NS00MDhlLThhMWEtMjQ0N2Y2Mzk2OTAz`,
	}
	header_wy = map[string]string{
		// MUSIC_U=000A32B5F2905E3227DBEFFC5C36250FC49DE0CF33A49B5FC6998B8507664B1E5408FC29A5C06EA23100E83D8E4C239090993406AB1F27ED03A7A978B4836527AF9189CB3BA0449C16AD634A2D50A78323B240368E04E05968460671EF377EFFA4B07319A6768D8D8A974B0E70E6F94195A52D77FC145049F05C1320401D0CE974C0604A1622C3EC5B7E5478B3E9F8004758E8C78D7900180F53F16BE9E5424E493FCAF122D8B3CB1C16CAACD7567F886790583AEB8B5D455EE1B48FBEEC1FB3F1C4BF5CEF685D718709C00DB1C76007D3BC32D5E5DB26927731DD4116F750356DB71380EF3523BCD47BD27A31C340B8444A4497AE277811AFD3B519DB585F85985EE7AF85765A567B54360FD59C54228CAF283D8D821251B94B09DB4ADC4F412951484B9150E9271166B475E2388BA75628912359A3DC5FDF64C68255225D3D070F1633447571ADC27909D3A5A3DF072A
		`Cookie`: `MUSIC_U=00B4C1E3FD77410780EF1C0840D08F3F5E7030E2D052CA8EC98A7368F7A7F6649B216E9533A1A174D72CCADF99554228E852DE46BBD2EA2A6B2A1433A3DF48B62EAA76FC18CD59256FEF6E76D39FB42DF76CE5068C69E3944E3A6E8E3C26135DBE0D9791FCE0BD524BD27F6226FD6460B05646A549A5C429F5E01EBA4E2D8D615BD715A7D245B13D9E570E87D0ADA608A607F2FAEF22AF8EE94F827AF150E9E1C517CB0F1588EF8F1D61947C43784985CF74F69458748960CE92053CA72B5FEF92C93F12F36714F0B346C2EAF89FAA516A8974E8CF53D5492DE95ED8591CCCF45AEB627C93B0CD370AEFB656EADAD031F688A6BB2CE3C9FA31BD6166A16ABEBEDADFCFEFBDCED5D4E12FFF1403C4F2B5A3F2422EF9D0878C0B52D08967D58E2E9DACE754404E2D6E1F81F52A1F1735CA9FBB85D758F81E0A7CBA41C5739D29E284F68430EB13E4F493890840031D3BD27E`,
	}
	header_kw = map[string]string{
		// `headers`: `Secret:6c3e1759abe6bd58f56bb713f6aee0bb738189eae7837be83636389b96fd4d7104c13520&&&Cookie:Hm_Iuvt_cdb524f42f0ce19b169b8072123a4727=2bm5QbPQKPZSRHyFN4pbZnGcNJ4J2DZJ`,
		`channel`: `qq`,
		`plat`:    `ar`,
		`net`:     `wifi`,
		`ver`:     `3.1.2`,
		`uid`:     ``,
		`devId`:   `0`,
	}
)

// func init() {
// 	json.Unmarshal([]byte(`{"kw":{"headers":"Secret:6c3e1759abe6bd58f56bb713f6aee0bb738189eae7837be83636389b96fd4d7104c13520&&&Cookie:Hm_Iuvt_cdb524f42f0ce19b169b8072123a4727=2bm5QbPQKPZSRHyFN4pbZnGcNJ4J2DZJ"},"mg":{"Referer":"https://m.music.migu.cn/v4/","User-Agent":"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.84 Safari/537.36 HBPC/12.1.2.300","By":"04f81461a98c7af557fea3cf28c4ea15","channel":"014000D","Cookie":"SESSION=ZTIwODkyMDQtOTE1NS00MDhlLThhMWEtMjQ0N2Y2Mzk2OTAz"},"wy":{"Cookie":"MUSIC_U=000A32B5F2905E3227DBEFFC5C36250FC49DE0CF33A49B5FC6998B8507664B1E5408FC29A5C06EA23100E83D8E4C239090993406AB1F27ED03A7A978B4836527AF9189CB3BA0449C16AD634A2D50A78323B240368E04E05968460671EF377EFFA4B07319A6768D8D8A974B0E70E6F94195A52D77FC145049F05C1320401D0CE974C0604A1622C3EC5B7E5478B3E9F8004758E8C78D7900180F53F16BE9E5424E493FCAF122D8B3CB1C16CAACD7567F886790583AEB8B5D455EE1B48FBEEC1FB3F1C4BF5CEF685D718709C00DB1C76007D3BC32D5E5DB26927731DD4116F750356DB71380EF3523BCD47BD27A31C340B8444A4497AE277811AFD3B519DB585F85985EE7AF85765A567B54360FD59C54228CAF283D8D821251B94B09DB4ADC4F412951484B9150E9271166B475E2388BA75628912359A3DC5FDF64C68255225D3D070F1633447571ADC27909D3A5A3DF072A"}}`), &def_headers)
// }
