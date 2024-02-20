package kg

import (
	"errors"
	"fmt"
	"lx-source/src/env"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
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
	// 检查用户配置文件，获取mixsongmid
	mixid := `582534238`
	if mixid == `auto` {
		mixid, err = randomMixSongMid()
		if err != nil {
			return
		}
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
	params := map[string]string{
		`userid`:     env.Config.Custom.Kg_userId,
		`token`:      env.Config.Custom.Kg_token,
		`appid`:      env.Config.Custom.Kg_Client_AppId,
		`clientver`:  env.Config.Custom.Kg_Client_Version,
		`clienttime`: strconv.FormatInt(time.Now().Unix(), 10),
		`mid`:        mid,
		`uuid`:       zcypt.HexToString(zcypt.RandomBytes(16)),
		`dfid`:       `-`,
	}

	// 发送请求
	var out refreshInfo
	err = signRequest(
		http.MethodPost,
		`https://gateway.kugou.com/v2/report/listen_song`,
		strings.NewReader(body),
		params, headers, &out,
	)
	if err != nil {
		return err
	}
	if out.Status != 1 {
		return errors.New(out.ErrorMsg)
	}
	fmt.Printf("%#v\n", out)
	env.Config.Custom.Kg_Lite_Interval = now + 86000

	return nil
}

func init() {
	env.Inits.Add(func() {
		if env.Config.Custom.Kg_Lite_Enable && false {
			if env.Config.Custom.Kg_Client_AppId == `3116` && env.Config.Custom.Kg_token != `` {
				env.Tasker.Add(`kg_refresh`, do_account_signin, 86000, true)
			}
		}
	})
}
