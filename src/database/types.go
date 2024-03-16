//go:build gorm

package database

// MusicFree 数据结构
// type (
// 	// 其他
// 	IExtra map[string]interface{}
// 	// 音乐
// 	IMusicItem struct {
// 		Artist   string `json:"artist"`             // 作者
// 		Title    string `json:"title"`              // 歌曲标题
// 		Duration int    `json:"duration,omitempty"` // 时长(s)
// 		Album    string `json:"album,omitempty"`    // 专辑名
// 		Artwork  string `json:"artwork,omitempty"`  // 专辑封面图
// 		Url      string `json:"url,omitempty"`      // 默认音源
// 		Lrc      string `json:"lrc,omitempty"`      // 歌词URL
// 		RawLrc   string `json:"rawLrc,omitempty"`   // 歌词文本（lrc格式 带时间戳）
// 		Other    IExtra `json:"extra,omitempty"`    // 其他
// 	}
// 	// 歌单
// 	IMusicSheetItem struct {
// 		Artwork     string       `json:"artwork,omitempty"`     // 封面图
// 		Title       string       `json:"title"`                 // 标题
// 		Description string       `json:"description,omitempty"` // 描述
// 		WorksNum    int          `json:"worksNum,omitempty"`    // 作品总数
// 		PlayCount   int          `json:"playCount,omitempty"`   // 播放次数
// 		MusicList   []IMusicItem `json:"musicList,omitempty"`   // 播放列表
// 		CreateAt    int64        `json:"createAt,omitempty"`    // 歌单创建日期
// 		Artist      string       `json:"artist,omitempty"`      // 歌单作者
// 		Other       IExtra       `json:"extra,omitempty"`       // 其他
// 	}
// 	// 专辑
// 	IAlbumItem IMusicSheetItem
// 	// 作者
// 	IArtistItem struct {
// 		Platform    string       `json:"platform,omitempty"`    // 插件名
// 		ID          interface{}  `json:"id"`                    // 唯一id
// 		Name        string       `json:"name"`                  // 姓名
// 		Fans        int          `json:"fans,omitempty"`        // 粉丝数
// 		Description string       `json:"description,omitempty"` // 简介
// 		Avatar      string       `json:"avatar,omitempty"`      // 头像
// 		WorksNum    int          `json:"worksNum,omitempty"`    // 作品数目
// 		MusicList   []IMusicItem `json:"musicList,omitempty"`   // 作者的音乐列表
// 		AlbumList   []IAlbumItem `json:"albumList,omitempty"`   // 作者的专辑列表
// 		Other       IExtra       `json:"extra,omitempty"`       // 其他
// 	}
// )

// 结构表
// type (
// 	// 重复
// 	XPublicKeys struct {
// 		ID  string `json:"id" gorm:"primaryKey"`  // 唯一ID
// 		Exp int64  `json:"exp" gorm:"column:exp"` // 过期时间
// 	}
// 	// 音乐
// 	XMusicItem struct {
// 		ID   string `json:"id" gorm:"primaryKey"`    // 唯一ID
// 		Name string `json:"name" gorm:"column:name"` // 歌曲名称
// 	}
// 	// 作者
// 	// XArtistItem struct {
// 	// 	ID   string `json:"id" gorm:"primaryKey"`
// 	// 	Name string `json:"name" gorm:"column:name"`
// 	// }
// 	// 歌词
// 	XLyricItem struct {
// 		ID      string `json:"id" gorm:"primaryKey"`
// 		Lyric   string `json:"lyric" gorm:"column:lyric"`     // 歌曲歌词
// 		TLyric  string `json:"tlyric" gorm:"column:tlyric"`   // 翻译歌词，没有可为 null
// 		RLyric  string `json:"rlyric" gorm:"column:rlyric"`   // 罗马音歌词，没有可为 null
// 		LxLyric string `json:"lxlyric" gorm:"column:lxlyric"` // lx 逐字歌词，没有可为 null
// 		// 歌词格式为 [分钟:秒.毫秒]<开始时间（基于该句）,持续时间>歌词文字
// 		// 例如： [00:00.000]<0,36>测<36,36>试<50,60>歌<80,75>词
// 	}
// 	// 视频
// 	// XMovieItem struct {
// 	// 	ID   string `json:"id" gorm:"primaryKey"`
// 	// 	Name string `json:"name" gorm:"column:name"`
// 	// }
// 	// 链接
// 	XLinkItem struct {
// 		ID string `json:"id" gorm:"primaryKey"`
// 	}
// )

// const (
// 	T_artist = `artist`
// 	T_detail = `detail`
// 	T_lyric  = `lyric`
// 	T_music  = `music`
// )

// 分源表
// type (
// 	// Music
// 	WyMusic MusicItem
// 	MgMusic MusicItem
// 	KwMusic MusicItem
// 	KgMusic MusicItem
// 	TxMusic MusicItem
// 	LxMusic MusicItem
// 	// Artist
// 	WyArtist ArtistItem
// 	MgArtist ArtistItem
// 	KwArtist ArtistItem
// 	KgArtist ArtistItem
// 	TxArtist ArtistItem
// 	LxArtist ArtistItem
// 	// Lyric
// 	WyLyric LyricItem
// 	MgLyric LyricItem
// 	KwLyric LyricItem
// 	KgLyric LyricItem
// 	TxLyric LyricItem
// 	LxLyric LyricItem
// )
