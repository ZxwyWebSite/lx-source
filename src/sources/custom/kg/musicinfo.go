package kg

import (
	"encoding/gob"
	"lx-source/src/env"
	"lx-source/src/sources"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/ZxwyWebSite/ztool"
)

func init() {
	gob.Register(musicInfo{})
}

func getMusicInfo(hash_ string) (info musicInfo, emsg string) {
	cquery := strings.Join([]string{`kg`, hash_, `info`}, `/`)
	if cdata, ok := env.Cache.Get(cquery); ok {
		if cinfo, ok := cdata.(musicInfo); ok {
			info = cinfo
			return
		}
	}
	body := ztool.Str_FastConcat(
		`{"area_code":"1","show_privilege":"1","show_album_info":"1","is_publish":"","appid":1005,"clientver":11451,"mid":"211008","dfid":"-","clienttime":"`,
		strconv.FormatInt(time.Now().Unix(), 10),
		`","key":"OIlwlieks28dk2k092lksi2UIkp","data":[{"hash":"`,
		hash_,
		`"}]}`,
	)
	var infoResp struct {
		Status    int           `json:"status"`
		ErrorCode int           `json:"error_code"`
		Errmsg    string        `json:"errmsg"`
		Data      [][]musicInfo `json:"data"`
	}
	err := ztool.Net_Request(
		http.MethodPost,
		`http://gateway.kugou.com/v3/album_audio/audio`,
		strings.NewReader(body),
		[]ztool.Net_ReqHandlerFunc{ztool.Net_ReqAddHeader(map[string]string{
			`KG-THash`:   `13a3164`,
			`KG-RC`:      `1`,
			`KG-Fake`:    `0`,
			`KG-RF`:      `00869891`,
			`User-Agent`: `Android712-AndroidPhone-11451-376-0-FeeCacheUpdate-wifi`,
			`x-router`:   `kmr.service.kugou.com`,
		})},
		[]ztool.Net_ResHandlerFunc{ztool.Net_ResToStruct(&infoResp)},
	)
	if err != nil {
		emsg = err.Error()
		return
	}
	if len(infoResp.Data) == 0 {
		if infoResp.Errmsg != `` {
			emsg = infoResp.Errmsg
		} else {
			emsg = `No Data`
		}
		return
	}
	info = infoResp.Data[0][0]
	emsg = infoResp.Errmsg
	env.Cache.Set(cquery, info, sources.C_lx)
	return
}
