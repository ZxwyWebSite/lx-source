package localcache

import (
	"errors"
	"lx-source/src/caches"
	"lx-source/src/env"
	"net/url"
	"os"
	"strings"

	"github.com/ZxwyWebSite/ztool"
)

type Cache struct {
	Path  string // 本地缓存目录 cache
	Bind  string // Api地址，用于生成外链 http://192.168.10.22:1011/
	state bool   // 激活状态
}

var loger = caches.Loger.AppGroup(`local`)

func (c *Cache) getLink(q string) string {
	return ztool.Str_FastConcat(c.Bind, `/file/`, q) // c.Addr + `file/` + q
}

func (c *Cache) Get(q *caches.Query) string {
	// 加一层缓存，减少重复检测文件造成的性能损耗
	if _, ok := env.Cache.Get(q.Query()); !ok {
		if _, e := os.Stat(ztool.Str_FastConcat(c.Path, `/`, q.Query())); e != nil {
			return ``
		}
		env.Cache.Set(q.Query(), struct{}{}, 3600)
	}
	return c.getLink(q.Query())
	// fpath := filepath.Join(c.Path, q.Source, q.MusicID, q.Quality)
	// if _, e := os.Stat(fpath); e != nil {
	// 	return ``
	// }
	// return c.getLink(fpath)
}

func (c *Cache) Set(q *caches.Query, l string) string {
	err := ztool.Net_DownloadFile(l, ztool.Str_FastConcat(c.Path, `/`, q.Query()), nil)
	if err != nil {
		loger.Error(`DownloadFile: %v`, err)
		return ``
	}
	env.Cache.Set(q.Query(), struct{}{}, 3600)
	return c.getLink(q.Query())
	// fpath := filepath.Join(c.Path, q.String)
	// os.MkdirAll(filepath.Dir(fpath), fs.ModePerm)
	// g := c.Loger.NewGroup(`localcache`)
	// ret, err := ztool.Net_HttpReq(http.MethodGet, l, nil, nil, nil)
	// if err != nil {
	// 	g.Warn(`HttpReq: %s`, err)
	// 	return ``
	// }
	// if err := os.WriteFile(fpath, ret, fs.ModePerm); err != nil {
	// 	g.Warn(`WriteFile: %s`, err)
	// 	return ``
	// }
	// return c.getLink(fpath)
}

func (c *Cache) Stat() bool {
	return c.state
}

func (c *Cache) Init() error {
	if c.Bind == `` {
		return errors.New(`请输入Api地址以生成外链`)
	} else {
		ubj, err := url.Parse(c.Bind)
		if err != nil {
			return err
		}
		ubj.Path = strings.TrimSuffix(ubj.Path, `/`)
		c.Bind = ubj.String()
	}
	c.state = true
	return nil
}

// func New(path, addr string, loger *logs.Logger) *Cache {
// 	return &Cache{
// 		Path:  path,
// 		Addr:  addr,
// 		Loger: loger,
// 		emsg:  cache.ErrNotInited,
// 	}
// }
