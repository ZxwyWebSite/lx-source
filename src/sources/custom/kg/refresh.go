package kg

import (
	"errors"
	"lx-source/src/env"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/ZxwyWebSite/ztool"
	"github.com/ZxwyWebSite/ztool/logs"
	"github.com/ZxwyWebSite/ztool/zcypt"
)

// 通过TOP500榜单获取随机歌曲的mixsongmid
func randomMixSongMid() (mid string, err error) {
	// 声明榜单url
	const rankUrl = `http://mobilecdnbj.kugou.com/api/v3/rank/song?version=9108&ranktype=1&plat=0&pagesize=100&area_code=1&page=1&rankid=8888&with_res_tag=0&show_portrait_mv=1`
	// 请求
	var res rankInfo
	err = ztool.Net_Request(
		http.MethodGet,
		rankUrl, nil,
		[]ztool.Net_ReqHandlerFunc{ztool.Net_ReqAddHeaders()},
		[]ztool.Net_ResHandlerFunc{ztool.Net_ResToStruct(&res)},
	)
	if err != nil {
		return
	}
	// fmt.Printf("%#v\n", res)
	if res.Status != 1 {
		err = errors.New(res.Error)
		return
	}

	// 随机选择一首歌曲
	randomSong := res.Data.Info[rand.Intn(len(res.Data.Info))]
	// fmt.Printf("%#v\n", randomSong)

	// 因为排行榜api不会返回mixsongmid
	// 所以需要进行一次搜索接口来获取
	var body searchInfo
	err = ztool.Net_Request(
		http.MethodGet,
		ztool.Str_FastConcat(
			`https://songsearch.kugou.com/song_search_v2?`,
			ztool.Net_Values(map[string]string{
				`keyword`:          randomSong.Filename,
				`area_code`:        `1`,
				`page`:             `1`,
				`pagesize`:         `1`,
				`userid`:           `0`,
				`clientver`:        ``,
				`platform`:         `WebFilter`,
				`filter`:           `2`,
				`iscorrection`:     `1`,
				`privilege_filter`: `0`,
			}),
		), nil,
		[]ztool.Net_ReqHandlerFunc{ztool.Net_ReqAddHeaders(map[string]string{
			`Referer`: `https://www.kugou.com`,
		})},
		[]ztool.Net_ResHandlerFunc{ztool.Net_ResToStruct(&body)},
	)
	if err != nil {
		return
	}
	// fmt.Printf("%#v\n", body)
	if body.Status != 1 {
		err = errors.New(body.ErrorMsg)
		return
	}
	if body.Data.Total == 0 || len(body.Data.Lists) == 0 {
		err = errors.New(`歌曲搜索失败`)
		return
	}
	mid = body.Data.Lists[0].MixSongID
	return
}

// 签到主函数，传入userinfo，响应None就是成功，报错即为不成功
func do_account_signin(loger *logs.Logger, now int64) (err error) {
	// 时间检测
	if now < env.Config.Custom.Kg_Lite_Interval {
		loger.Debug(`Key未过期，跳过...`)
		return nil
	}
	// 检查用户配置文件，获取mixsongmid
	mixid := env.Config.Custom.Kg_Lite_MixId //`582534238`
	if mixid == `auto` || mixid == `` {
		mixid, err = randomMixSongMid()
		if err != nil {
			return
		}
		loger.Info(`成功获取MixSongMid: ` + mixid)
	} else {
		loger.Info(`使用固定MixSongMid: ` + mixid)
	}

	// 声明变量
	headers := map[string]string{
		`User-Agent`: ztool.Str_FastConcat(
			`Android712-AndroidPhone-`,
			env.Config.Custom.Kg_Client_Version,
			`-18-0-NetMusic-wifi`,
		),
		`KG-THash`: `3e5ec6b`,
		`KG-Rec`:   `1`,
		`KG-RC`:    `1`,
		`x-router`: `youth.kugou.com`,
	}
	body := ztool.Str_FastConcat(
		`{"mixsongid":"`, mixid, `"}`,
	)
	tnow := time.Now()
	params := map[string]string{
		`userid`:     env.Config.Custom.Kg_userId,
		`token`:      env.Config.Custom.Kg_token,
		`appid`:      env.Config.Custom.Kg_Client_AppId,
		`clientver`:  env.Config.Custom.Kg_Client_Version,
		`clienttime`: strconv.FormatInt(tnow.Unix(), 10),
		`mid`:        mid,
		`uuid`:       zcypt.HexToString(zcypt.RandomBytes(16)),
		`dfid`:       `-`,
	}

	// 发送请求
	var out refreshInfo
	err = signRequest(
		http.MethodPost,
		`https://gateway.kugou.com/v2/report/listen_song`,
		body, params, headers, &out,
	)
	if err != nil {
		return err
	}
	loger.Debug(`Resp: %+v`, out)
	if out.Status != 1 {
		if out.ErrorCode == 130012 {
			loger.Info(`今日已签到过，明天再来吧`)
		} else {
			return errors.New(out.ErrorMsg)
		}
	} else {
		loger.Info(`Lite签到成功`)
	}
	tomorrow := time.Date(tnow.Year(), tnow.Month(), tnow.Day()+1, 0, 0, 0, 0, tnow.Location())
	env.Config.Custom.Kg_Lite_Interval = tomorrow.Unix()

	return env.Cfg.Save(``)
}

func init() {
	env.Inits.Add(func() {
		if env.Config.Custom.Kg_Lite_Enable {
			if env.Config.Custom.Kg_Client_AppId == `3116` && env.Config.Custom.Kg_token != `` {
				env.Tasker.Add(`kg_refresh`, do_account_signin, 86000, true)
			}
		}
	})
}
