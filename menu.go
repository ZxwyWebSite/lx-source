package main

import (
	"lx-source/src/env"
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
	default:
		loger.Fatal(`未知参数:%q`, *etag)
	}
	loger.Free()
}

// func menuMian() {
// 	app := menu.NewApp(`Lx-Source`)
// 	app.Data = menu.Data{
// 		`Main`: func(this *menu.App) string { return ` ` },
// 	}
// 	app.Run()
// 	os.Exit(0)
// }
