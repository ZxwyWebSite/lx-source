package database

// 结构表
// type (
// 	PublicKeys struct {
// 		ID string `json:"id" gorm:"primaryKey"`
// 	}
// 	// 音乐
// 	MusicItem struct {
// 		ID     string `json:"id" gorm:"primaryKey"` // 唯一ID
// 		Name   string `json:"name" gorm:"column:name"` //
// 		// Source string `json:"source" gorm:"-:all"`
// 	}
// 	// 作者
// 	ArtistItem struct {
// 		ID   string `json:"id" gorm:"primaryKey"`
// 		Name string `json:"name" gorm:"column:name"`
// 	}
// 	// 歌词
// 	LyricItem struct {
// 		ID      string `json:"id" gorm:"primaryKey"`
// 		Lyric   string `json:"lyric" gorm:"column:lyric"`     // 歌曲歌词
// 		TLyric  string `json:"tlyric" gorm:"column:tlyric"`   // 翻译歌词，没有可为 null
// 		RLyric  string `json:"rlyric" gorm:"column:rlyric"`   // 罗马音歌词，没有可为 null
// 		LxLyric string `json:"lxlyric" gorm:"column:lxlyric"` // lx 逐字歌词，没有可为 null
// 		// 歌词格式为 [分钟:秒.毫秒]<开始时间（基于该句）,持续时间>歌词文字
// 		// 例如： [00:00.000]<0,36>测<36,36>试<50,60>歌<80,75>词
// 	}
// 	// 视频
// 	MovieItem struct {
// 		ID   string `json:"id" gorm:"primaryKey"`
// 		Name string `json:"name" gorm:"column:name"`
// 	}
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
