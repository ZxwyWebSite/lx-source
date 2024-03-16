//go:build extapp

package server

import (
	"lx-source/src/middleware/resp"

	"github.com/ZxwyWebSite/ztool"
	"github.com/gin-gonic/gin"
)

type queryObj struct {
	Name      string `json:"name"`      // 歌名
	Singer    string `json:"singer"`    // 歌手
	Source    string `json:"source"`    // 平台
	Songmid   string `json:"songmid"`   // 音乐ID
	Interval  string `json:"interval"`  // 时长
	AlbumName string `json:"albumName"` // 专辑
	Img       string `json:"img"`       // 封面
	// TypeURL   struct {
	// } `json:"typeUrl"` // 未知
	AlbumID string `json:"albumId"` // 专辑ID
	// 支持音质
	Types []struct {
		Type string `json:"type"` // 音质
		Size string `json:"size"` // 大小
		Hash string `json:"hash"` // 哈希(kg only)
	} `json:"types"`
	// tx
	StrMediaMid string `json:"strMediaMid"` // 当前文件ID
	AlbumMid    string `json:"albumMid"`    // 专辑ID
	SongID      int    `json:"songId"`      // 音乐ID
	// mg
	CopyrightID string `json:"copyrightId"` // 音乐ID
	LrcURL      string `json:"lrcUrl"`      // lrc歌词
	MrcURL      string `json:"mrcUrl"`      // mrc歌词
	TrcURL      string `json:"trcUrl"`      // trc歌词
	// kg
	Hash string `json:"hash"` // 文件哈希(kg only)
}

func loadLxMusic(lx *gin.RouterGroup) {
	// 获取链接
	lx.POST(`/link/:q`, func(c *gin.Context) {
		resp.Wrap(c, func() *resp.Resp {
			var obj queryObj
			if err := c.ShouldBindJSON(&obj); err != nil {
				return &resp.Resp{Code: 6, Msg: `解析错误: ` + err.Error()}
			}
			pams := map[string]string{
				`s`:  obj.Source,
				`id`: ztool.Str_Select(obj.Hash, obj.CopyrightID, obj.Songmid),
			}
			for k, v := range pams {
				c.Params = append(c.Params, gin.Param{Key: k, Value: v})
			}
			return nil
		})
	}, linkHandler)
	// lx.GET(`/info`)
}
