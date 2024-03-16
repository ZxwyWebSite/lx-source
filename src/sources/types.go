//go:build extapp

package sources

// MusicFree 数据结构
type (
	// 其他
	IExtra map[string]interface{}
	// 音乐
	IMusicItem struct {
		Artist   string `json:"artist"`             // 作者
		Title    string `json:"title"`              // 歌曲标题
		Duration int    `json:"duration,omitempty"` // 时长(s)
		Album    string `json:"album,omitempty"`    // 专辑名
		Artwork  string `json:"artwork,omitempty"`  // 专辑封面图
		Url      string `json:"url,omitempty"`      // 默认音源
		Lrc      string `json:"lrc,omitempty"`      // 歌词URL
		RawLrc   string `json:"rawLrc,omitempty"`   // 歌词文本（lrc格式 带时间戳）
		Other    IExtra `json:"extra,omitempty"`    // 其他
	}
	// 歌单
	IMusicSheetItem struct {
		Artwork     string       `json:"artwork,omitempty"`     // 封面图
		Title       string       `json:"title"`                 // 标题
		Description string       `json:"description,omitempty"` // 描述
		WorksNum    int          `json:"worksNum,omitempty"`    // 作品总数
		PlayCount   int          `json:"playCount,omitempty"`   // 播放次数
		MusicList   []IMusicItem `json:"musicList,omitempty"`   // 播放列表
		CreateAt    int64        `json:"createAt,omitempty"`    // 歌单创建日期
		Artist      string       `json:"artist,omitempty"`      // 歌单作者
		Other       IExtra       `json:"extra,omitempty"`       // 其他
	}
	// 专辑
	IAlbumItem IMusicSheetItem
	// 作者
	IArtistItem struct {
		Platform    string       `json:"platform,omitempty"`    // 插件名
		ID          interface{}  `json:"id"`                    // 唯一id
		Name        string       `json:"name"`                  // 姓名
		Fans        int          `json:"fans,omitempty"`        // 粉丝数
		Description string       `json:"description,omitempty"` // 简介
		Avatar      string       `json:"avatar,omitempty"`      // 头像
		WorksNum    int          `json:"worksNum,omitempty"`    // 作品数目
		MusicList   []IMusicItem `json:"musicList,omitempty"`   // 作者的音乐列表
		AlbumList   []IAlbumItem `json:"albumList,omitempty"`   // 作者的专辑列表
		Other       IExtra       `json:"extra,omitempty"`       // 其他
	}
)
