package main

import (
	"lx-source/src/env"
	wm "lx-source/src/sources/custom/wy/modules"
	"strings"
	"time"

	qrcode "github.com/skip2/go-qrcode"
)

func parseEtag(etag *string) {
	if etag == nil {
		return
	}
	loger := env.Loger.NewGroup(`ParseEtag`)
	switch *etag {
	case ``:
		break
	case `menu`:
		loger.Fatal(`暂不支持交互菜单，敬请期待...`)
		// menuMian()
	case `wyqr`:
		wyQrLogin()
	default:
		loger.Fatal(`未知参数:%q`, *etag)
	}
	loger.Free()
}

// 网易云扫码登录
func wyQrLogin() {
	loger := env.Loger.NewGroup(`WyQrLogin`)
	defer loger.Free()
	loger.Info(`执行模块: 网易云扫码登录`)
	res, err := wm.LoginQrKey()
	if err != nil {
		loger.Fatal(`无法创建请求: %s`, err)
	}
	key := res.Body[`unikey`].(string)
	loger.Info(`创建请求成功: %v`, key)

	link := wm.LoginQrCreate(key)
	qr, err := qrcode.New(link, qrcode.Low)
	if err != nil {
		loger.Fatal(`无法生成二维码: %s`, err)
	}
	loger.Info("\n请使用网易云音乐手机APP扫描以下二维码授权登录:\n%v", qr.ToSmallString(false))

	for {
		time.Sleep(time.Second * 5)
		res, err = wm.LoginQrCheck(key)
		if err != nil {
			loger.Error(`检测状态失败: %s`, err)
			continue
		}
		msg := res.Body[`message`].(string)
		switch msg {
		case `等待扫码`:
			loger.Info(msg)
		case `授权中`:
			loger.Info(`扫码成功: %q, 请在手机上确认登录`, res.Body[`nickname`])
		case `授权登陆成功`:
			loger.Info(`授权成功`)
			env.Config.Custom.Wy_Enable = true
			env.Config.Custom.Wy_Mode = `163api`
			env.Config.Custom.Wy_Api_Cookie = strings.Join(res.Cookie, `; `)
			env.Config.Custom.Wy_Refresh_Enable = true
			if err := env.Cfg.Save(``); err != nil {
				loger.Error(`配置保存失败: %s`, err)
			} else {
				loger.Info(`配置保存成功`)
			}
			return
		case `二维码不存在或已过期`:
			loger.Fatal(`授权请求超时，请重试！`)
		default:
			loger.Fatal(`未知状态: %v`, msg)
		}
	}
}

// func menuMian() {
// 	app := menu.NewApp(`Lx-Source`)
// 	app.Data = menu.Data{
// 		`Main`: func(this *menu.App) string { return ` ` },
// 	}
// 	app.Run()
// 	os.Exit(0)
// }
