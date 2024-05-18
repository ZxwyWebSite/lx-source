package cloudcache

import (
	"lx-source/src/caches"
	"lx-source/src/env"
	"net/http"
	"strings"

	cr "github.com/ZxwyWebSite/cr-go-sdk"
	"github.com/ZxwyWebSite/cr-go-sdk/service/explorer"
	"github.com/ZxwyWebSite/ztool"
)

type Cache struct {
	Site  *cr.SiteObj
	Path  string
	state bool
}

func (c *Cache) Get(q *caches.Query) string {
	var b strings.Builder
	b.WriteString(c.Path)
	b.WriteByte('/')
	b.WriteString(q.Source)
	b.WriteByte('/')
	b.WriteString(q.MusicID)
	list, err := c.Site.Directory(b.String())
	if err != nil {
		caches.Loger.Debug(`列出目录: %v`, err)
		return ``
	}
	name := q.Quality + `.` + q.Extname
	var id string
	for _, v := range list.Objects {
		if v.Name == name && v.Type == `file` {
			id = v.ID
			break
		}
	}
	if id == `` {
		caches.Loger.Debug(`文件不存在`)
		return ``
	}
	srcs, err := c.Site.FileSource(cr.GenerateSrc(false, id))
	if err != nil {
		caches.Loger.Debug(`生成外链: %v`, err)
		return ``
	}
	return (*srcs)[0].URL
	/*link, err := c.Site.FileDownload(id)
	if err != nil {
		caches.Loger.Debug(`下载文件: %v`, err)
		return ``
	}
	if (*link)[0] == '/' {
		return c.Site.Addr + (*link)[1:]
	}
	return *link*/
}

func (c *Cache) Set(q *caches.Query, l string) string {
	var b strings.Builder
	b.WriteString(c.Path)
	b.WriteByte('/')
	b.WriteString(q.Source)
	b.WriteByte('/')
	b.WriteString(q.MusicID)
	dir := b.String()
	err := c.Site.DirectoryNew(&explorer.DirectoryService{
		Path: dir,
	})
	if err != nil {
		caches.Loger.Debug(`创建目录: %v`, err)
		return ``
	}
	/*var buf bytes.Buffer
	err = ztool.Net_Download(l, &buf, nil)
	if err != nil {
		caches.Loger.Debug(`下载文件: %v`, err)
		return ``
	}*/
	name := q.Quality + `.` + q.Extname
	err = ztool.Net_Request(
		http.MethodGet, l, nil,
		[]ztool.Net_ReqHandlerFunc{ztool.Net_ReqAddHeaders()},
		[]ztool.Net_ResHandlerFunc{func(res *http.Response) error {
			return (&cr.UploadTask{
				Site: c.Site,
				File: res.Body,
				Size: uint64(res.ContentLength),
				Name: name,
				Mime: `audio/mpeg`,
			}).Do(dir)
		}},
	)
	if err != nil {
		caches.Loger.Debug(`上传文件: %v`, err)
		return ``
	}
	return c.Get(q)
}

func (c *Cache) Stat() bool {
	return c.state
}

func (c *Cache) Init() error {
	cr.Cr_Debug = env.Config.Main.Debug
	err := c.Site.SdkInit()
	if err != nil {
		return err
	}
	if c.Site.Users.Cookie == nil || c.Site.Config.User.Anonymous {
		err = c.Site.SdkLogin()
		if err != nil {
			return err
		}
	}
	c.state = true
	return nil
}
