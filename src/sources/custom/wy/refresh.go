package wy

import (
	"lx-source/src/env"
	wy "lx-source/src/sources/custom/wy/modules"
	"maps"

	// "time"

	"github.com/ZxwyWebSite/ztool/logs"
	"github.com/ZxwyWebSite/ztool/x/cookie"
)

/*
 刷新登录模块 (来自 NeteaseCloudMusicApi)
 逻辑：
  检测返回结果中是否含有"MUSIC_U":
    如果有则为正常刷新，延时30天
    否则延时1天
 注：
  原代码未提供详细描述，无法确定有效结果判断条件，暂时先这么写
 2024-02-15:
  MUSIC_U 改变 则 6天 后 继续执行
  MUSIC_U 不变 则 1天 后 继续执行
*/

func refresh(loger *logs.Logger, now int64) error {
	// 前置检测
	// now := time.Now().Unix() //(执行时间已改为从参数获取)
	if now < env.Config.Custom.Wy_Refresh_Interval {
		loger.Debug(`Key未过期，跳过...`)
		return nil
	}
	// 刷新逻辑
	cookies := cookie.ToMap(cookie.Parse(env.Config.Custom.Wy_Api_Cookie))
	res, err := wy.LoginRefresh(wy.ReqQuery{
		Cookie: cookies,
	})
	loger.Debug(`Resp: %+v`, res)
	if err == nil {
		if out, ok := res.Body[`cookie`].(string); ok {
			loger.Info(`获取数据成功`)
			cmap := cookie.ToMap(cookie.Parse(out))
			maps.Copy(cookies, cmap)
			env.Config.Custom.Wy_Api_Cookie = cookie.Marshal(cookies)
			loger.Debug(`Cookie: %#v`, cookies)
			if _, ok := cmap[`MUSIC_U`]; ok {
				// MUSIC_U 改变 则 6天 后 继续执行
				env.Config.Custom.Wy_Refresh_Interval = now + 518400 //2147483647 - 86000
				loger.Debug(`MUSIC_U 改变, 6天 后 继续执行`)
			} else {
				// MUSIC_U 不变 则 1天 后 继续执行
				env.Config.Custom.Wy_Refresh_Interval = now + 86000
				loger.Debug(`MUSIC_U 不变, 1天 后 继续执行`) //`未发现有效结果，将在下次检测时再次尝试`
			}
			err = env.Cfg.Save(``)
			if err == nil {
				loger.Info(`配置更新成功`)
			}
		}
	}
	return err
}

func init() {
	env.Inits.Add(func() {
		if env.Config.Custom.Wy_Refresh_Enable && env.Config.Custom.Wy_Api_Cookie != `` {
			env.Tasker.Add(`wy_refresh`, refresh, 86000, true)
		}
	})
}
