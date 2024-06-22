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
	Version = `1.0.3.0622`
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
	// 系统
	Conf_Main struct {
		Debug   bool     `comment:"调试模式"`
		Listen  []string `comment:"监听地址 (多端口以','分隔)"`
		Gzip    bool     `comment:"开启GZip (对已压缩的内容使用会产生反效果)"`
		Cors    bool     `comment:"添加跨域响应头 (兼容前端请求)"`
		LogPath string   `comment:"文件日志路径，不填禁用"`
		Print   bool     `comment:"控制台输出 (影响io性能，后台使用建议关闭)"`
		SysLev  bool     `comment:"(实验性) 设置进程高优先级"`
		// FFConv  bool   `comment:"(实验性) 使用FFMpeg修复音频(本地缓存)"`
		NgProxy bool   `comment:"兼容反向代理(beta)"`
		Timeout int64  `comment:"网络请求超时(单位:秒,海外服务器可适当调大)"`
		Store   string `comment:"内存缓存持久化文件地址"`
		ErrMp3  string `comment:"获取失败默认音频"`
	}
	// 接口
	Conf_Apis struct {
		// 预留：后期可能会出一个WebUI，/webui，相关设置
		// WebUI
		// WebUI_Enable bool `comment:"是否开启WebUI相关接口"`
		// App
		// App_Enable bool `comment:"是否开启软件接口"`
		// App Lx-Music
		// App_LX_Enable bool `comment:"是否开启Lx-Music相关接口"`
		// App MusicFree
		// App_MF_Enable bool `comment:"是否开启MusicFree相关接口"`
	}
	// 验证
	Conf_Auth struct {
		// ApiKey
		ApiKey_Enable bool   `comment:"是否开启Key验证"`
		ApiKey_Value  string `comment:"验证Key值，留空自动生成"`
		// 速率限制
		RateLimit_Enable bool   `comment:"是否开启速率限制"`
		RateLimit_Block  uint32 `comment:"检测范围，每分区为x秒"` // 每x秒一个分区
		// RateLimit_Global uint32 `comment:"全局速率限制，单位次每x秒(暂未开放)"`
		RateLimit_Single uint32 `comment:"单IP速率限制，单位次每x秒"`
		RateLimit_BanNum uint32 `comment:"容忍限度，超出限制N次后封禁"`
		RateLimit_BanTim uint32 `comment:"封禁后每次延长时间"`
		// 黑白名单
		// BanList_Mode  string   `comment:"名单模式 0: off(关闭), 1: white(白名单), 2: black(黑名单)"`
		// BanList_White []string `comment:"host白名单"`
		// BanList_Black []string `comment:"host黑名单"`
	}
	// 来源
	Conf_Source struct {
		// Mode string `comment:"音乐来源 0: off(关闭 仅本地), 1: builtin(内置), 2: custom(登录账号 暂不支持)"`
		// 伪装IP
		FakeIP_Mode  string `comment:"伪装IP模式 0: off(关闭), 1: req(传入值), 2: val(静态)"`
		FakeIP_Value string `comment:"静态伪装IP"`
		// 代理
		Proxy_Enable  bool   `comment:"使用代理"`
		Proxy_Address string `comment:"代理地址 (支持http, socks)"`
		// 验证
		MusicIdVerify bool `comment:"(beta) 验证音乐ID可用性"`
		ForceFallback bool `comment:"忽略音质限制,强制返回链接(部分支持)"`
		// 总开关(解决部分源无法彻底禁用问题)?
		Enable_Wy bool `comment:"是否开启小芸源"`
		Enable_Mg bool `comment:"是否开启小蜜源"`
		Enable_Kw bool `comment:"是否开启小蜗源"`
		Enable_Kg bool `comment:"是否开启小枸源"`
		Enable_Tx bool `comment:"是否开启小秋源"`
		Enable_Lx bool `comment:"是否开启小洛源"`
	} // `comment:""`
	// 账号
	Conf_Custom struct {
		// wy
		Wy_Enable bool   `comment:"是否启用小芸源"`
		Wy_Mode   string `comment:"获取方式 0: builtin, 1: 163api"`
		// wy 163api
		Wy_Api_Type    string `comment:"调用方式 0: native(内置模块), 1: remote(指定地址)"`
		Wy_Api_Address string `comment:"NeteaseCloudMusicApi项目地址"`
		Wy_Api_Cookie  string `comment:"账号cookie数据"`
		// wy refresh
		Wy_Refresh_Enable   bool  `comment:"是否启用刷新登录"`
		Wy_Refresh_Interval int64 `comment:"下次刷新时间 (由程序维护)"`

		// mg
		Mg_Enable bool   `comment:"是否启用小蜜源"`
		Mg_Mode   string `comment:"获取方式 0: builtin, 1: custom"`
		// mg custom
		Mg_Usr_VerId string `comment:"field user.aversionid"`
		Mg_Usr_Token string `comment:"field user.token"`
		Mg_Usr_OSVer string `comment:"field user.osversion"`
		Mg_Usr_ReqUA string `comment:"field user.useragent"`
		// mg refresh
		Mg_Refresh_Enable   bool  `comment:"是否启用Cookie保活"`
		Mg_Refresh_Interval int64 `comment:"下次运行时间 (自动更新)"`

		// kw
		Kw_Enable bool   `comment:"是否启用小蜗源"`
		Kw_Mode   string `comment:"接口模式 0: bdapi(需验证), 1: kwdes"`
		// kw bdapi
		Kw_Bd_Uid   string `comment:"field user.uid"`
		Kw_Bd_Token string `comment:"field user.token"`
		Kw_Bd_DevId string `comment:"field user.device_id"`
		// kw kwdes
		Kw_Des_Type   string `comment:"返回格式 0: text, 1: json"`
		Kw_Des_Source string `comment:"query source"`
		Kw_Des_Header string `comment:"请求头 User-Agent"`

		// kg
		Kg_Enable bool `comment:"是否启用小枸源"`
		// kg client
		Kg_Client_AppId     string `comment:"酷狗音乐的appid，官方安卓为1005，官方PC为1001（client.appid）"`
		Kg_Client_SignKey   string `comment:"客户端signature采用的key值，需要与appid对应（client.signatureKey）"`
		Kg_Client_Version   string `comment:"客户端versioncode，pidversionsecret可能随此值而变化（client.clientver）"`
		Kg_Client_PidVerSec string `comment:"获取URL时所用的key值计算验证值（client.pidversionsecret）"`
		Kg_Client_Pid       string `comment:"field client.pid"`
		// kg user
		Kg_token  string `comment:"field user.token"`
		Kg_userId string `comment:"field user.userid"`
		// kg lite_sign_in
		Kg_Lite_Enable   bool   `comment:"是否启用概念版自动签到，仅在appid=3116时运行"`
		Kg_Lite_MixId    string `comment:"mix_songmid的获取方式, 默认auto, 可以改成一个数字手动"`
		Kg_Lite_Interval int64  `comment:"调用时间，自动刷新"`
		// kg refresh_login
		Kg_Refresh_Enable   bool  `comment:"是否启动刷新登录"`
		Kg_Refresh_Interval int64 `comment:""`

		// tx
		Tx_Enable bool   `comment:"是否启用小秋源"`
		Tx_Ukey   string `comment:"Cookie中/客户端的请求体中的（comm.authst）"`
		Tx_Uuin   string `comment:"key对应的QQ号"`
		Tx_CDNUrl string `comment:"指定音频CDN地址"`
		// tx refresh_login
		Tx_Refresh_Enable   bool  `comment:"是否启动刷新登录"`
		Tx_Refresh_Interval int64 `comment:"刷新间隔 (由程序维护，非必要无需修改)"`

		// lx (local)
		// Lx_Enable bool `comment:"是否启用小洛源"`
	}
	// 脚本
	Conf_Script_Update struct {
		Ver   string `comment:"自定义脚本版本" json:"ver"`
		Log   string `comment:"更新日志" json:"log"`
		Url   string `comment:"脚本下载地址 (public目录内文件名)" json:"url"`
		Force bool   `comment:"强制推送更新" json:"force"`
	}
	Conf_Script struct {
		Name     string `comment:"源的名字，建议不要过长，24个字符以内"`
		Descript string `comment:"源的描述，建议不要过长，36个字符以内，可不填"`
		Version  string `comment:"源的版本号，可不填"`
		Author   string `comment:"脚本作者名字，可不填"`
		Homepage string `comment:"脚本主页，可不填"`

		Update Conf_Script_Update `ini:"Script"`

		Auto int `comment:"自动填写配置(beta) 0: 关闭, 1: 仅api地址, 2: 包含密钥"`
	}
	// 缓存
	Conf_Cache struct {
		Mode     string `comment:"缓存模式 0: off(关闭), 1: local(本地), 2: cloudreve(云盘 未完善)"`
		LinkMode string `comment:"外链样式 1: static(永久链), 2: dynamic(临时链)"`
		// 本地
		Local_Path string `comment:"本地缓存保存路径"`
		Local_Bind string `comment:"本地缓存外部访问地址"`
		Local_Auto bool   `comment:"自适应缓存访问地址(beta)"`
		// 云盘
		Cloud_Site string `comment:"Cloudreve站点地址"`
		Cloud_User string `comment:"Cloudreve用户名"`
		Cloud_Pass string `comment:"Cloudreve密码"`
		Cloud_Sess string `comment:"Cloudreve会话"`
		Cloud_Path string `comment:"Cloudreve存储路径"`
	}
	// 结构
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
	DefCfg = Conf{
		Main: Conf_Main{
			Debug:   false,
			Listen:  []string{`127.0.0.1:1011`},
			Gzip:    false,
			LogPath: `/data/logfile.log`,
			Print:   true,
			SysLev:  false,
			Timeout: 30,
			Store:   `/data/memo.bin`,
			ErrMp3:  `https://r2eu.zxwy.link/gh/lx-source/static/error.mp3`,
		},
		Apis: Conf_Apis{
			// BindAddr: `http://192.168.10.22:1011/`,
		},
		Auth: Conf_Auth{
			ApiKey_Enable:    true,
			RateLimit_Enable: false,
			RateLimit_Block:  30,
			// RateLimit_Global: 1,
			RateLimit_Single: 15,
			RateLimit_BanNum: 5,
			RateLimit_BanTim: 10,
			// BanList_Mode:     `off`,
			// BanList_White:    []string{`127.0.0.1`},
		},
		Source: Conf_Source{
			// Mode:          `builtin`,
			FakeIP_Mode:   `0`,
			FakeIP_Value:  `192.168.10.2`,
			Proxy_Enable:  false,
			Proxy_Address: `{protocol}://({user}:{password})@{address}:{port}`,

			Enable_Wy: true,
			Enable_Mg: true,
			Enable_Kw: true,
			Enable_Kg: true,
			Enable_Tx: true,
			Enable_Lx: false,
		},
		Custom: Conf_Custom{
			Wy_Enable:           true,
			Wy_Mode:             `builtin`,
			Wy_Api_Type:         `native`,
			Wy_Refresh_Interval: 1633622400,

			Mg_Enable:    true,
			Mg_Mode:      `builtin`,
			Mg_Usr_OSVer: `10`,
			Mg_Usr_ReqUA: `Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.82 Safari/537.36`,

			Kw_Enable:     true,
			Kw_Mode:       `kwdes`,
			Kw_Des_Type:   `json`,
			Kw_Des_Header: `okhttp/3.10.0`,

			Kg_Enable:           false,
			Kg_Client_AppId:     `1005`,
			Kg_Client_SignKey:   `OIlwieks28dk2k092lksi2UIkp`,
			Kg_Client_Version:   `12029`,
			Kg_Client_PidVerSec: `57ae12eb6890223e355ccfcb74edf70d`,
			Kg_Client_Pid:       `2`,
			Kg_userId:           `0`,
			Kg_Lite_MixId:       `auto`,

			Tx_Enable:           false,
			Tx_CDNUrl:           `https://isure6.stream.qqmusic.qq.com/`,
			Tx_Refresh_Enable:   false,
			Tx_Refresh_Interval: 86000,
		},
		Script: Conf_Script{
			Name:     `Lx-Source-Script`,
			Descript: `洛雪音乐自定义源脚本`,
			Version:  `1.1.0`,
			Author:   `Zxwy`,
			Homepage: `https://github.com/ZxwyWebSite/lx-script`,

			Update: Conf_Script_Update{
				Log: `发布更新 (请删除旧源后重新导入)：进行了部分优化，修复了部分Bug`, // 更新日志

				Ver:   `1.0.3`, // 自定义脚本版本
				Force: true,    // 强制推送更新

				Url: `lx-custom-source.js`, // 脚本下载地址
			},
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
	Config = DefCfg
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
	Inits = new(ztool.Err_DeferList)

	Tasker = task.New(time.Hour, 2) // 定时任务 (暂时没有什么快速任务，默认每小时检测一次)
)

// func init() {
// }
