// 全局变量
package env

import (
	"time"

	"github.com/ZxwyWebSite/ztool"
	"github.com/ZxwyWebSite/ztool/cache/memo"
	"github.com/ZxwyWebSite/ztool/conf"
	"github.com/ZxwyWebSite/ztool/logs"
	"github.com/ZxwyWebSite/ztool/task"
)

const (
	Version = `1.0.2-b0.9`
)

var (
	RunPath string
)

// 配置结构
/*
 注：Mode字段建议使用名称方式调用，序号可能频繁更改
 e.g. 0: off(), 1: builtin(), 2: custom()
     序号           名称                描述
*/
type (
	Conf_Main struct {
		Debug   bool   `comment:"调试模式"`
		Listen  string `comment:"监听地址"`
		Gzip    bool   `comment:"开启GZip (对已压缩的内容使用会产生反效果)"`
		LogPath string `comment:"文件日志路径，不填禁用"`
		Print   bool   `comment:"控制台输出"`
		SysLev  bool   `comment:"(实验性) 设置进程高优先级"`
		// FFConv  bool   `comment:"(实验性) 使用FFMpeg修复音频(本地缓存)"`
	}
	Conf_Apis struct {
		// 预留：后期可能会出一个WebUI，/webui，相关设置
		// BindAddr string `comment:"外部访问地址，用于生成文件链接"`
		// LxM_Auth string `comment:"验证Key，自动生成，填写null禁用"`
	}
	Conf_Auth struct {
		// ApiKey
		ApiKey_Enable bool   `comment:"是否开启Key验证"`
		ApiKey_Value  string `comment:"验证Key值，留空自动生成"`
		// 速率限制
		RateLimit_Enable bool   `comment:"是否开启速率限制"`
		RateLimit_Block  uint32 `comment:"检测范围，每分区为x秒"` // 每x秒一个分区
		RateLimit_Global uint32 `comment:"全局速率限制，单位次每x秒(暂未开放)"`
		RateLimit_Single uint32 `comment:"单IP速率限制，单位次每x秒"`
		RateLimit_BanNum uint32 `comment:"容忍限度，超出限制N次后封禁"`
		RateLimit_BanTim uint32 `comment:"封禁后每次延长时间"`
		// 黑白名单
		BanList_Mode  string   `comment:"名单模式 0: off(关闭), 1: white(白名单), 2: black(黑名单)"`
		BanList_White []string `comment:"host白名单"`
		BanList_Black []string `comment:"host黑名单"`
	}
	Conf_Source struct {
		Mode string `comment:"音乐来源 0: off(关闭 仅本地), 1: builtin(内置), 2: custom(登录账号 暂不支持)"`
		// 伪装IP
		FakeIP_Mode  string `comment:"伪装IP模式 0: off(关闭), 1: req(传入值), 2: val(静态)"`
		FakeIP_Value string `comment:"静态伪装IP"`
		// 代理
		Proxy_Enable  bool   `comment:"使用代理"`
		Proxy_Address string `comment:"代理地址 (支持http, socks)"`
		// 平台账号
		// ...(待实现)
	} // `comment:""`
	Conf_Custom struct {
		// wy (暂未实现)
		Wy_Enable bool `comment:"是否开启小芸源"`
		// Wy_Cookie string `comment:"账号cookie数据"`
		// mg (暂未实现)
		// Mg_Enable bool `comment:"是否开启小蜜源"`
		// kg (暂未实现)
		// Kg_Enable bool `comment:"是否开启小枸源"`
		// tx
		Tx_Enable bool   `comment:"是否开启小秋源"`
		Tx_Ukey   string `comment:"Cookie中/客户端的请求体中的（comm.authst）"`
		Tx_Uuin   string `comment:"key对应的QQ号"`
		// tx refresh_login
		Tx_Refresh_Enable   bool  `comment:"是否启动刷新登录"`
		Tx_Refresh_Interval int64 `comment:"刷新间隔 (由程序维护，非必要无需修改)"`
	}
	Conf_Script struct {
		Ver   string `comment:"自定义脚本版本" json:"ver"`
		Log   string `comment:"更新日志" json:"log"`
		Url   string `comment:"脚本下载地址 (public目录内文件名)" json:"url"`
		Force bool   `comment:"强制推送更新" json:"force"`
	}
	Conf_Cache struct {
		Mode     string `comment:"缓存模式 0: off(关闭), 1: local(本地), 2: cloudreve(云盘 未完善)"`
		LinkMode string `comment:"外链样式 1: static(永久链), 2: dynamic(临时链)"`
		// 本地
		Local_Path string `comment:"本地缓存保存路径"`
		Local_Bind string `comment:"本地缓存外部访问地址"`
		// 云盘
		Cloud_Site string `comment:"Cloudreve站点地址"`
		Cloud_User string `comment:"Cloudreve用户名"`
		Cloud_Pass string `comment:"Cloudreve密码"`
		Cloud_Sess string `comment:"Cloudreve会话"`
		Cloud_Path string `comment:"Cloudreve存储路径"`
	}
	Conf struct {
		Main   Conf_Main   `comment:"程序主配置"`
		Apis   Conf_Apis   `comment:"接口设置"`
		Auth   Conf_Auth   `comment:"访问控制"`
		Source Conf_Source `comment:"解析源配置"`
		Custom Conf_Custom `comment:"解析账号配置"`
		Script Conf_Script `comment:"自定义脚本更新"` // ini:",omitempty"
		Cache  Conf_Cache  `comment:"音乐缓存设置"`
	}
)

var (
	// 默认配置
	defCfg = Conf{
		Main: Conf_Main{
			Debug:   false,
			Listen:  `127.0.0.1:1011`,
			Gzip:    false,
			LogPath: `/data/logfile.log`,
			Print:   true,
			SysLev:  false,
		},
		Apis: Conf_Apis{
			// BindAddr: `http://192.168.10.22:1011/`,
		},
		Auth: Conf_Auth{
			ApiKey_Enable:    true,
			RateLimit_Enable: false,
			RateLimit_Block:  30,
			RateLimit_Global: 1,
			RateLimit_Single: 15,
			RateLimit_BanNum: 5,
			RateLimit_BanTim: 10,
			BanList_Mode:     `off`,
			BanList_White:    []string{`127.0.0.1`},
		},
		Source: Conf_Source{
			Mode:          `builtin`,
			FakeIP_Mode:   `0`,
			FakeIP_Value:  `192.168.10.2`,
			Proxy_Enable:  false,
			Proxy_Address: `{protocol}://({user}:{password})@{address}:{port}`,
		},
		Custom: Conf_Custom{
			Wy_Enable:           true,
			Tx_Enable:           false,
			Tx_Refresh_Enable:   false,
			Tx_Refresh_Interval: 86000,
		},
		Script: Conf_Script{
			Log: `发布更新 (请删除旧源后重新导入)：进行了部分优化，修复了部分Bug`, // 更新日志

			Ver:   `1.0.3`,               // 自定义脚本版本
			Url:   `lx-custom-source.js`, // 脚本下载地址
			Force: true,                  // 强制推送更新
		},
		Cache: Conf_Cache{
			Mode:       `local`, // 缓存模式
			LinkMode:   `1`,
			Local_Path: `data/cache`,
			Local_Bind: `http://127.0.0.1:1011/`,
			Cloud_Site: `https://cloudreveplus-demo.onrender.com/`,
			Cloud_User: `admin@cloudreve.org`,
			Cloud_Pass: `CloudrevePlusDemo`,
			Cloud_Sess: ``,
			Cloud_Path: `/Lx-Source/cache`,
		},
	}
	Config = defCfg
	// 通用对象
	Loger  = logs.NewLogger(`LX-SOURCE`)
	Cfg, _ = conf.New(&Config, &conf.Confg{
		AutoFormat: true,
		UseBuf:     true,
		UnPretty:   true,
		Loger:      Loger.NewGroup(`Config`),
	})
	Defer = new(ztool.Err_DeferList)
	Cache = memo.NewMemoStoreConf(Loger, 300) // 内存缓存 默认每5分钟进行一次GC //memo.NewMemoStore()

	Tasker = task.New(time.Hour, 2) // 定时任务 (暂时没有什么快速任务，默认每小时检测一次)
)

// func init() {

// }
