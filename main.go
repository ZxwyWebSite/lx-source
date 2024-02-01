package main

import (
	"context"
	"flag"
	"io/fs"
	"lx-source/src/env"
	"lx-source/src/server"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"sync"
	"sync/atomic"
	"syscall"
	"time"

	"github.com/ZxwyWebSite/ztool"
	"github.com/ZxwyWebSite/ztool/logs"
	"github.com/gin-gonic/gin"
)

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
	etag := flag.String(`e`, ``, `扩展启动参数`)
	perm := flag.Uint(`p`, 0, `自定义文件权限(8进制前面加0)`)
	flag.Parse()
	if perm != nil && *perm != 0 {
		ztool.Fbj_DefPerm = fs.FileMode(*perm)
		fp := env.Loger.NewGroup(`FilePerm`)
		// if ztool.Fbj_DefPerm > 777 {
		// 	fp.Fatal(`请在实际权限前面加0`)
		// }
		fp.Info(`设置默认文件权限为 %o (%v)`, *perm, ztool.Fbj_DefPerm).Free()
	}
	parseEtag(etag)
	env.Cfg.MustInit(confPath)
	// fmt.Printf("%+v\n", env.Config)
	env.Loger.NewGroup(`ServHello`).Info(`欢迎使用 LX-SOURCE 洛雪音乐自定义源`).Free()
	if !env.Config.Main.Debug {
		gin.SetMode(gin.ReleaseMode)
	} else {
		logs.Levell = logs.LevelDebu // logs.Level = 3
		env.Loger.NewGroup(`DebugMode`).Debug(`已开启调试模式, 将输出更详细日志 (配置文件中 [Main].Debug 改为 false 关闭)`).Free()
	}
	genAuth()
	if env.Config.Main.SysLev {
		sl := env.Loger.NewGroup(`(beta)SysLev`)
		if err := ztool.Sys_SetPriorityLev(ztool.Sys_GetPid(), ztool.Sys_PriorityHighest); err != nil {
			sl.Error(`系统优先级设置失败: %v`, err)
		} else {
			sl.Warn(`成功设置较高优先级，此功能可能导致系统不稳定`)
		}
		sl.Free()
	}
	if env.Config.Main.Timeout != env.DefCfg.Main.Timeout {
		ztool.Net_client.Timeout = time.Second * time.Duration(env.Config.Main.Timeout) // 自定义请求超时
		env.Loger.NewGroup(`InitNet`).Info(`请求超时已设为 %s`, ztool.Net_client.Timeout).Free()
	}
}

func main() {
	defer env.Defer.Do()
	// 初始化基础功能
	initMain()

	// 载入必要模块
	env.Inits.Do()
	env.Loger.NewGroup(`ServInit`).Info(`服务端启动, 监听地址 %s`, strings.Join(env.Config.Main.Listen, `|`)).Free()
	loadFileLoger()
	env.Defer.Add(env.Tasker.Run(env.Loger)) // wait

	// 启动Http服务
	listenAndServe(server.InitRouter(), env.Config.Main.Listen)
}

// 监听多端口
func listenAndServe(handler http.Handler, addrs []string) {
	// 前置检测
	length := len(addrs)
	ss := env.Loger.NewGroup(`ServStart`)
	if length == 0 {
		ss.Fatal(`监听地址列表为空`)
	}
	// ss.Info(`服务端启动,请稍候...`)
	srvlist := make(map[int]*http.Server, length) // 伪数组，便于快速删除数据
	lock := new(sync.Mutex)
	var failnum int32
	length32 := int32(length)
	// 启动服务
	for i := 0; i < length; i++ {
		lock.Lock()
		srvlist[i] = &http.Server{Addr: addrs[i], Handler: handler}
		lock.Unlock()
		go func(n int) {
			server := srvlist[n]
			// ss.Info(`开始监听 %v`, server.Addr)
			if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				ss.Error(`监听%q失败: %s`, server.Addr, err) // 监听":1011"失败: http: Server closed
				fn := atomic.AddInt32(&failnum, 1)
				if fn == length32 {
					ss.Fatal(`所有地址监听失败，程序被迫退出`)
				}
				lock.Lock()
				delete(srvlist, n)
				lock.Unlock()
			}
		}(i)
	}
	// time.Sleep(time.Millisecond * 300)
	// if len(srvlist) == 0 {
	// 	ss.Fatal(`所有地址监听失败，程序被迫退出`)
	// }
	// ss.Free()
	// 安全退出
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGQUIT)
	<-quit
	sc := env.Loger.NewGroup(`ServClose`)
	sc.Info(`等待结束活动连接...`)
	// 停止服务
	var unsafenum int32
	wg := new(sync.WaitGroup)
	for i := range srvlist {
		wg.Add(1)
		go func(n int) {
			server := srvlist[n]
			ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
			if err := server.Shutdown(ctx); err != nil {
				sc.Warn(`连接%q未安全退出: %s`, server.Addr, err) // 连接":1011"未安全退出: timeout
				atomic.AddInt32(&unsafenum, 1)
			}
			cancel()
			wg.Done()
		}(i)
	}
	wg.Wait()
	if unsafenum != 0 {
		sc.Warn(`未安全退出 :(`)
	} else {
		sc.Info(`已安全退出 :)`)
	}
}
