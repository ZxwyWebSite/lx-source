package main

import (
	"context"
	"encoding/base64"
	"flag"
	"lx-source/src/caches"
	"lx-source/src/caches/localcache"
	"lx-source/src/env"
	"lx-source/src/router"
	"lx-source/src/sources"
	"lx-source/src/sources/builtin"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"github.com/ZxwyWebSite/ztool"
	"github.com/ZxwyWebSite/ztool/logs"
	"github.com/gin-gonic/gin"
)

func genAuth() {
	ga := env.Loger.NewGroup(`LxM-Auth`)
	// 检测Key是否存在, 否则生成并保存
	if env.Config.Auth.ApiKey_Value == `` {
		randomBytes := func(size int) []byte {
			buf := make([]byte, size)
			for i := 0; i < size; i++ {
				buf[i] = byte(rand.New(rand.NewSource(time.Now().UnixNano() + int64(i*i+rand.Intn(256)))).Intn(256))
			}
			return buf
		}
		pass := base64.StdEncoding.EncodeToString(randomBytes(4 * 4))
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
}

// 加载文件日志 (请在初始化配置文件后调用)
func loadFileLoger() {
	// 最后加载FileLoger保证必要日志已输出 (Debug模式强制在控制台输出日志)
	if env.Config.Main.LogPath != `` {
		lg := env.Loger.NewGroup(`FileLoger`)
		printout := env.Config.Main.Print // || env.Config.Main.Debug
		_, do, err := env.Loger.SetOutFile(ztool.Str_FastConcat(env.RunPath, env.Config.Main.LogPath), printout)
		if err == nil {
			env.Defer.Add(do)
			gin.DefaultWriter = env.Loger.GetOutput()
			gin.ForceConsoleColor()
			// lg.Info(`文件日志初始化成功`)
		} else {
			lg.Error(`文件日志初始化失败：%v`, err)
		}
	}
}

// 初始化
func init() {
	ztool.Cmd_FastPrint(ztool.Str_FastConcat(`
     __      __  __      ______  ______  __  __  ____    ______  ______
    / /     / / / /     / ____/ / __  / / / / / / __ \  / ____/ / ____/
   / /     / /_/ / __  / /___  / / / / / / / / / /_/ / / /     / /___  
  / /      \_\ \  /_/ /___  / / / / / / / / / /  ___/ / /     / ____/  
 / /___  / / / /     ____/ / / /_/ / / /_/ / / / \   / /___  / /___    
/_____/ /_/ /_/     /_____/ /_____/ /_____/ /_/ \_\ /_____/ /_____/    
=======================================================================
  Version: `, env.Version, `  Github: https://github.com/ZxwyWebSite/lx-source
`, "\n"))
	env.RunPath, _ = os.Getwd()
	var confPath string
	flag.StringVar(&confPath, `c`, ztool.Str_FastConcat(env.RunPath, `/data/conf.ini`), `指定配置文件路径`)
	flag.Parse()
	// fileLoger() // 注：记录日志必然会影响性能，自行选择是否开启
	// logs.DefLogger(`LX-SOURCE`, logs.LevelDebu)
	// logs.Main = `LX-SOURCE`
	env.Cfg.MustInit(confPath) //conf.InitConfig(confPath)
	// fmt.Printf("%+v\n", env.Config)
	env.Loger.NewGroup(`ServHello`).Info(`欢迎使用 LX-SOURCE 洛雪音乐自定义源`)
	if !env.Config.Main.Debug {
		gin.SetMode(gin.ReleaseMode)
	} else {
		logs.Levell = logs.LevelDebu // logs.Level = 3
		env.Loger.NewGroup(`DebugMode`).Debug(`已开启调试模式, 将输出更详细日志 (配置文件中 [Main].Debug 改为 false 关闭)`)
	}
	genAuth()
	if env.Config.Main.SysLev {
		sl := env.Loger.NewGroup(`(beta)SysLev`)
		if err := ztool.Sys_SetPriorityLev(ztool.Sys_GetPid(), ztool.Sys_PriorityHighest); err != nil {
			sl.Error(`系统优先级设置失败: %v`, err)
		} else {
			sl.Warn(`成功设置较高优先级，此功能可能导致系统不稳定`)
		}
	}
}

func main() {
	defer env.Defer.Do()

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
		icl.Warn(`本地缓存绑定地址: %q, 请确认其与实际访问地址相符`, env.Config.Cache.Local_Bind)
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

	// 载入必要模块
	env.Inits.Do()
	env.Loger.NewGroup(`ServStart`).Info(`服务端启动, 监听地址 %s`, env.Config.Main.Listen)
	loadFileLoger()
	env.Defer.Add(env.Tasker.Run(env.Loger)) // wait

	// 启动Http服务
	r := router.InitRouter() //InitRouter()
	server := &http.Server{
		Addr:    env.Config.Main.Listen,
		Handler: r,
	}
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			env.Loger.NewGroup(`InitRouter().Run`).Fatal(`%s`, err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGQUIT)
	<-quit
	sc := env.Loger.NewGroup(`ServClose`)
	sc.Info(`等待结束活动连接...`) // Shutdown Server ...

	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		sc.Fatal(`未安全退出: %s`, err) // Server Shutdown
	}
	sc.Info(`已安全退出 :)`) // Server exited

	// if err := InitRouter().Run(env.Config.Main.Listen); err != nil {
	// 	env.Loger.NewGroup(`InitRouter().Run`).Fatal(`%s`, err)
	// }
}
