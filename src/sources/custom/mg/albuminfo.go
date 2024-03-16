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
	gob.Register(albumInfo{})
}

func getAlbumInfo(aid string) (info albumInfo, err error) {
	cquery := strings.Join([]string{`mg`, aid, `ainfo`}, `/`)
	if cdata, ok := env.Cache.Get(cquery); ok {
		if cinfo, ok := cdata.(albumInfo); ok {
			info = cinfo
			return
		}
	}
	err = ztool.Net_Request(
		http.MethodGet, ztool.Str_FastConcat(
			`https://app.c.nf.migu.cn/MIGUM2.0/v1.0/content/resourceinfo.do?needSimple=01&resourceId=`, aid, `&resourceType=2003`,
		), nil,
		[]ztool.Net_ReqHandlerFunc{ztool.Net_ReqAddHeaders()},
		[]ztool.Net_ResHandlerFunc{ztool.Net_ResToStruct(&info)},
	)
	if err == nil {
		if len(info.Resource) == 0 {
			err = errors.New(`no Album Resource`)
		} else {
			env.Cache.Set(cquery, info, sources.C_lx)
		}
	}
	return
}
