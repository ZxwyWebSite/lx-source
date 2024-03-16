package mg

import (
	"errors"
	"lx-source/src/env"
	"net/http"

	"github.com/ZxwyWebSite/ztool"
	"github.com/ZxwyWebSite/ztool/logs"
)

func refresh(loger *logs.Logger, now int64) error {
	var out map[string]any
	err := ztool.Net_Request(
		http.MethodPost,
		`https://m.music.migu.cn/migumusic/h5/user/auth/userActiveNotice`,
		nil,
		[]ztool.Net_ReqHandlerFunc{ztool.Net_ReqAddHeaders(mgheader)},
		[]ztool.Net_ResHandlerFunc{ztool.Net_ResToStruct(&out)},
	)
	if err == nil {
		if out[`code`].(int) != http.StatusOK {
			return errors.New(out[`msg`].(string))
		} else {
			loger.Info(`咪咕session保活成功`)
		}
	}
	return err
}

func init() {
	env.Inits.Add(func() {
		if env.Config.Custom.Mg_Refresh_Enable && false {
			env.Tasker.Add(`mg_refresh`, refresh, 86000, true)
		}
	})
}
