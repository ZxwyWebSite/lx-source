package main

import (
	"encoding/base64"
	"lx-source/src/caches"
	"lx-source/src/caches/localcache"
	"lx-source/src/env"
	"lx-source/src/sources"
	"lx-source/src/sources/builtin"
	"net/http"
	stdurl "net/url"
	"path/filepath"

	"github.com/ZxwyWebSite/ztool"
	"github.com/ZxwyWebSite/ztool/logs"
	"github.com/ZxwyWebSite/ztool/zcypt"
	"github.com/gin-gonic/gin"
)

// 生成连接码
func genAuth() {
	ga := env.Loger.NewGroup(`LxM-Auth`)
	// 检测Key是否存在, 否则生成并保存
	if env.Config.Auth.ApiKey_Value == `` {
		pass := zcypt.Base64ToString(base64.StdEncoding, zcypt.RandomBytes(4*4))
		env.Config.Auth.ApiKey_Value = pass // env.Config.Apis.LxM_Auth
		ga.Info(`已生成默认连接码: %q`, pass)
		ga.Info(`可在配置文件 [Auth].ApiKey_Value 中修改`) //可在配置文件 [Apis].LxM_Auth 中修改, 填写 "null" 关闭验证
		if err := env.Cfg.Save(``); err != nil {
			ga.Error(`写入配置文件失败: %s, 将导致下次启动时连接码发生变化`, err)
		}
	}
	if !env.Config.Auth.ApiKey_Enable {
		ga.Warn(`已关闭Key验证, 公开部署可能导致安全隐患`)
	} else {
		ga.Warn(`已开启Key验证, 记得在脚本中填写 apipass=%q`, env.Config.Auth.ApiKey_Value)
	}
	ga.Free()
}

// 加载文件日志 (请在初始化配置文件后调用)
func loadFileLoger() {
	// 最后加载FileLoger保证必要日志已输出 (Debug模式强制在控制台输出日志)
	if env.Config.Main.LogPath != `` {
		lg := env.Loger.NewGroup(`FileLoger`)
		printout := env.Config.Main.Print // || env.Config.Main.Debug
		f, do, err := env.Loger.SetOutFile(ztool.Str_FastConcat(env.RunPath, env.Config.Main.LogPath), printout)
		if err == nil {
			// env.Defer.Add(do)
			env.Defer.Add(func() { do(); f.Close() })
			env.Tasker.Add(`flog_flush`, func(loger *logs.Logger, now int64) error {
				loger.Debug(`已写入文件并清理日志缓存`)
				return do()
			}, 3600, false)
			gin.DefaultWriter = env.Loger.GetOutput()
			gin.ForceConsoleColor()
			// lg.Info(`文件日志初始化成功`)
		} else {
			lg.Error(`文件日志初始化失败：%v`, err)
		}
		lg.Free()
	}
}

// 初始化基础功能
func initMain() {
	// 初始化代理
	ipr := env.Loger.NewGroup(`InitProxy`)
	switch env.Config.Source.FakeIP_Mode {
	case `0`, `off`:
		break
	case `1`, `req`:
		ipr.Fatal(`暂未实现此功能`)
	case `2`, `val`:
		if env.Config.Source.FakeIP_Value != `` {
			ipr.Info(`已开启伪装IP，当前值: %v`, env.Config.Source.FakeIP_Value)
			ztool.Net_header[`X-Real-IP`] = env.Config.Source.FakeIP_Value
			ztool.Net_header[`X-Forwarded-For`] = env.Config.Source.FakeIP_Value
		} else {
			ipr.Error(`伪装IP为空，请检查配置 [Source].FakeIP_Value`)
		}
	default:
		ipr.Error(`未定义的代理模式，请检查配置 [Source].FakeIP_Mode，本次启动禁用IP伪装`)
	}
	if env.Config.Source.Proxy_Enable {
		ipr.Debug(`ProxyAddr: %v`, env.Config.Source.Proxy_Address)
		addr, err := stdurl.Parse(env.Config.Source.Proxy_Address)
		if err != nil {
			ipr.Error(`代理Url解析失败: %s, 将禁用代理功能`, err)
		} else {
			type chkRegion struct {
				AmapFlag    int    `json:"amap_flag"`
				IPFlag      int    `json:"ip_flag"`
				AmapAddress string `json:"amap_address"`
				Country     string `json:"country"`
				Flag        int    `json:"flag"`
				Errcode     int    `json:"errcode"`
				Status      int    `json:"status"`
				Error       string `json:"error"`
			}
			var out chkRegion
			oldval := ztool.Net_client.Transport
			ztool.Net_client.Transport = &http.Transport{Proxy: http.ProxyURL(addr)}
			err := ztool.Net_Request(http.MethodGet,
				`https://mips.kugou.com/check/iscn?&format=json`, nil,
				[]ztool.Net_ReqHandlerFunc{ztool.Net_ReqAddHeader(ztool.Net_header)},
				[]ztool.Net_ResHandlerFunc{ztool.Net_ResToStruct(&out)},
			)
			if err != nil {
				ztool.Net_client.Transport = oldval
				ipr.Error(`地区验证失败: %s, 已恢复默认配置`, err)
			} else {
				ipr.Debug(`Resp: %+v`, out)
				if out.Flag != 1 {
					ipr.Warn(`您正在使用非中国大陆(%v)代理，可能导致部分音乐不可用`, out.Country)
				} else {
					ipr.Warn(`代理开启成功，请注意潜在的Cookie泄露问题`)
				}
			}
		}
	}
	ipr.Free()

	// 初始化缓存
	icl := env.Loger.NewGroup(`InitCache`)
	switch env.Config.Cache.Mode {
	case `0`, `off`:
		// NothingToDo... (已默认禁用缓存)
		break
	case `1`, `local`:
		// 注：由于需要修改LocalCachePath参数，必须在InitRouter之前执行
		cache, err := caches.New(&localcache.Cache{
			Path: filepath.Join(env.RunPath, env.Config.Cache.Local_Path),
			Bind: env.Config.Cache.Local_Bind,
		})
		if err != nil {
			icl.Error(`驱动["local"]初始化失败: %v, 将禁用缓存功能`, err)
		}
		caches.UseCache = cache
		icl.Warn(`本地缓存绑定地址:%q,请确认其与实际访问地址相符`, env.Config.Cache.Local_Bind)
		// LocalCachePath = filepath.Join(runPath, env.Config.Cache.Local_Path)
		// UseCache = &localcache.Cache{
		// 	Path: LocalCachePath,
		// 	Addr: env.Config.Apis.BindAddr,
		// }
		// icl.Info(`使用本地缓存，文件路径 %q，绑定地址 %v`, LocalCachePath, env.Config.Apis.BindAddr)
	case `2`, `cloudreve`:
		icl.Fatal(`Cloudreve驱动暂未完善，未兼容新版调用方式，当前版本禁用`)
		// icl.Warn(`Cloudreve驱动暂未完善，使用非本机存储时存在兼容性问题，请谨慎使用`)
		// cs, err := cloudreve.NewSite(&cloudreve.Config{
		// 	SiteUrl:  env.Config.Cache.Cloud_Site,
		// 	Username: env.Config.Cache.Cloud_User,
		// 	Password: env.Config.Cache.Cloud_Pass,
		// 	Session:  env.Config.Cache.Cloud_Sess,
		// })
		// if err != nil {
		// 	icl.Error(`驱动["cloudreve"]初始化失败: %v, 将禁用缓存功能`, err)
		// }
		// UseCache = &crcache.Cache{
		// 	Cs:   cs,
		// 	Path: env.Config.Cache.Cloud_Path,
		// 	IsOk: err == nil,
		// }
	default:
		icl.Error(`未定义的缓存模式，请检查配置 [Cache].Mode，本次启动禁用缓存`)
	}
	icl.Free()

	// 初始化音乐源
	ise := env.Loger.NewGroup(`InitSource`)
	switch env.Config.Source.Mode {
	case `0`, `off`:
		break
	case `1`, `builtin`:
		sources.UseSource = &builtin.Source{}
	case `2`, `custom`:
		ise.Fatal(`暂未实现账号解析源`)
	default:
		ise.Error(`未定义的音乐源，请检查配置 [Source].Mode，本次启动禁用内置源`)
	}
	ise.Free()
}
