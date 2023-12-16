package caches

import (
	"lx-source/src/env"

	"github.com/ZxwyWebSite/ztool"
)

type (
	// 查询参数
	Query struct {
		Source  string // source 平台 wy, mg
		MusicID string // sid 音乐ID wy: songmid, mg: copyrightId
		Quality string // quality 音质 128k / 320k / flac / flac24bit
		Extname string // rext 扩展名 mp3 / flac (没有前缀点)
		query   string // 查询字符串缓存
	}
	// 缓存需实现以下接口
	Cache interface {
		// 获取缓存 (查询参数 query)(外链)
		/* `wy/10086/128k.mp3`->`http://192.168.10.22:1011/file/wy/10086/128k.mp3` */
		Get(*Query) string
		// 设置缓存 (查询参数 query, 音乐直链 link)(外链)
		/* (`wy/10086/128k.mp3`,`https://xxx.xxxx.xx/file.mp3`)->`http://192.168.10.22:1011/file/wy/10086/128k.mp3` */
		Set(*Query, string) string
		// 可用状态 true/false
		Stat() bool
		// 初始化 ()(错误)
		Init() error
	}
)

// 默认无缓存的缓存
type Nullcache struct{}

func (*Nullcache) Get(*Query) string         { return `` }
func (*Nullcache) Set(*Query, string) string { return `` }
func (*Nullcache) Stat() bool                { return false }
func (*Nullcache) Init() error               { return nil }

var (
	Loger = env.Loger.NewGroup(`Caches`)

	UseCache Cache = &Nullcache{}

	// ErrNotInited = errors.New(`缓存策略未初始化`)
)

// 根据音质判断文件后缀
func rext(q string) string {
	if q == `128k` || q == `320k` {
		return `mp3`
	}
	return `flac`
}

// 生成查询参数 (必须使用此函数初始化)
func NewQuery(s, id, q string) *Query {
	return &Query{
		Source:  s,
		MusicID: id,
		Quality: q,
		Extname: rext(q),
	}
}

// 获取旧版查询字符串
func (c *Query) Query() string {
	if c.query == `` {
		c.query = ztool.Str_FastConcat(c.Source, `/`, c.MusicID, `/`, c.Quality, `.`, c.Extname)
	}
	return c.query
}

// 初始化缓存
func New(c Cache) (Cache, error) {
	err := c.Init()
	return c, err
	// if err != nil {
	// 	return nil, err
	// }
	// return c, nil
}
func MustNew(c Cache) Cache {
	out, err := New(c)
	if err != nil {
		panic(err)
	}
	return out
}
