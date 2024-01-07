package tx

import (
	"lx-source/src/env"

	"github.com/ZxwyWebSite/ztool/logs"
)

// func Info(songMid string) (info any, msg string) {
// 	req, emsg := getMusicInfo(songMid)
// 	if emsg != `` {
// 		msg = emsg
// 		return
// 	}
// 	var singerList []any
// 	for _, s := range req.TrackInfo.Singer {
// 		// item := new(struct{
// 		// 	ID int `json:"id"`
// 		// 	Str string
// 		// })
// 		// item.ID = s.ID
// 		singerList = append(singerList, s)
// 	}
// 	var file_info map[string]struct{
// 		Size interface{}
// 	}
// 	if req.TrackInfo.File.Size128Mp3 != 0 {
// 		file_info[`128k`].Size =
// 	}
// 	return
// }

func Init() {
	if env.Config.Custom.Tx_Refresh_Enable {
		env.Tasker.Add(`refresh_login`, func(l *logs.Logger) error {
			refresh(l)
			return nil
		}, 86000, true)
	}
}
