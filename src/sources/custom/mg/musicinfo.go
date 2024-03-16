package mg

import (
	"encoding/gob"
	"errors"
	"lx-source/src/env"
	"lx-source/src/sources"
	"net/http"
	"strings"

	"github.com/ZxwyWebSite/ztool"
)

func init() {
	gob.Register(musicInfo{})
}

func getMusicInfo(cid string) (info musicInfo, err error) {
	cquery := strings.Join([]string{`mg`, cid, `minfo`}, `/`)
	if cdata, ok := env.Cache.Get(cquery); ok {
		if cinfo, ok := cdata.(musicInfo); ok {
			info = cinfo
			return
		}
	}
	err = ztool.Net_Request(
		http.MethodGet, ztool.Str_FastConcat(
			`https://c.musicapp.migu.cn/MIGUM2.0/v1.0/content/resourceinfo.do?copyrightId=`, cid, `&resourceType=2`,
		), nil,
		[]ztool.Net_ReqHandlerFunc{ztool.Net_ReqAddHeaders()},
		[]ztool.Net_ResHandlerFunc{ztool.Net_ResToStruct(&info)},
	)
	if err == nil {
		if len(info.Resource) == 0 {
			err = errors.New(`no Music Resource`)
		} else {
			env.Cache.Set(cquery, info, sources.C_lx)
		}
	}
	return
}
