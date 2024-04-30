package kg

import (
	"crypto/aes"
	"crypto/cipher"
	"errors"
	"lx-source/src/env"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/ZxwyWebSite/ztool"
	"github.com/ZxwyWebSite/ztool/logs"
	"github.com/ZxwyWebSite/ztool/x/bytesconv"
	"github.com/ZxwyWebSite/ztool/x/json"
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
		for i := 0; true; i++ {
			mixid, err = randomMixSongMid()
			if err != nil {
				loger.Error(`ReTry: %v, Err: %s`, i, err)
				if i >= 2 {
					return
				}
				time.Sleep(time.Second)
				continue
			}
			break
		}
		// mixid, err = randomMixSongMid()
		// if err != nil {
		// 	return
		// }
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
	if out.Status != 1 || out.ErrorCode != 0 {
		switch out.ErrorCode {
		case 130012:
			loger.Info(`今日已签到过，明天再来吧`)
		case 51002:
			panic(`登录过期啦！请重新获取账号Token`)
		default:
			return errors.New(strconv.Itoa(out.ErrorCode) + `: ` + out.ErrorMsg)
		}
	} else {
		loger.Info(`Lite签到成功`)
	}
	tomorrow := time.Date(tnow.Year(), tnow.Month(), tnow.Day()+1, 0, 0, 0, 0, tnow.Location())
	env.Config.Custom.Kg_Lite_Interval = tomorrow.Unix()

	return env.Cfg.Save(``)
}

// 刷新Token
func login_by_token(loger *logs.Logger, now int64) (err error) {
	// 前置到期检测
	if now < env.Config.Custom.Kg_Refresh_Interval {
		loger.Debug(`Key未过期，跳过...`)
		return
	}
	// 获取加密参数
	var aeskey []byte
	switch env.Config.Custom.Kg_Client_AppId {
	case `1005`:
		aeskey = []byte(`90b8382a1bb4ccdcf063102053fd75b8`)
	case `3116`:
		aeskey = []byte(`c24f74ca2820225badc01946dba4fdf7`)
	default:
		panic(`当前应用AppId暂不支持此功能`)
	}
	// 生成请求数据
	tnow := time.Now()
	pbyte, _ := json.Marshal(map[string]any{
		`clienttime`: tnow.Unix(),
		`token`:      env.Config.Custom.Kg_token,
	})
	block, _ := aes.NewCipher(aeskey)
	encrypter := cipher.NewCBCEncrypter(block, aeskey[block.BlockSize():])
	padata := zcypt.PKCS7Padding(pbyte, block.BlockSize())
	encrypted := make([]byte, len(padata))
	encrypter.CryptBlocks(encrypted, padata)
	encstr := zcypt.HexToString(encrypted)
	bodys, _ := json.Marshal(map[string]any{
		`t1`:            0,
		`t2`:            0,
		`p3`:            encstr,
		`userid`:        env.Config.Custom.Kg_userId,
		`clienttime_ms`: tnow.UnixMilli(),
	})
	params := map[string]string{
		`dfid`:       `-`,
		`mid`:        `20211008`,
		`clientver`:  env.Config.Custom.Kg_Client_Version,
		`clienttime`: strconv.FormatInt(tnow.Unix(), 10),
		`appid`:      env.Config.Custom.Kg_Client_AppId,
	}
	headers := map[string]string{
		`User-Agent`: `Android711-1070-10860-14-0-LOGIN-wifi`,
		`KG-THash`:   `7af653c`,
		`KG-Rec`:     `1`,
		`KG-RC`:      `1`,
	}
	// 请求对应接口
	var res loginInfo
	err = signRequest(
		http.MethodPost,
		`http://login.user.kugou.com/v4/login_by_token`,
		bytesconv.BytesToString(bodys),
		params, headers, &res,
	)
	if err != nil {
		return errors.New(`接口请求失败: ` + err.Error())
	}
	loger.Info(`获取数据成功`)
	loger.Debug(`Resp: %+v`, res)
	if res.ErrorCode != 0 {
		return errors.New(`刷新登录失败: ` + strconv.Itoa(res.ErrorCode))
	}
	env.Config.Custom.Kg_token = res.Data.Token
	env.Config.Custom.Kg_userId = strconv.Itoa(res.Data.Userid)
	next := time.Date(tnow.Year(), tnow.Month(), tnow.Day()+25, 0, 0, 0, 0, tnow.Location())
	env.Config.Custom.Kg_Refresh_Interval = next.Unix()
	loger.Info(`刷新登录成功`)
	return env.Cfg.Save(``)
}

func init() {
	env.Inits.Add(func() {
		if env.Config.Custom.Kg_token != `` {
			if env.Config.Custom.Kg_Lite_Enable && env.Config.Custom.Kg_Client_AppId == `3116` {
				env.Tasker.Add(`kg_litsign`, do_account_signin, 86000, true)
			}
			if env.Config.Custom.Kg_Refresh_Enable {
				env.Tasker.Add(`kg_refresh`, login_by_token, 86000, true)
			}
		}
		/*if env.Config.Custom.Kg_Lite_Enable {
			if env.Config.Custom.Kg_Client_AppId == `3116` && env.Config.Custom.Kg_token != `` {
				env.Tasker.Add(`kg_litsign`, do_account_signin, 86000, true)
			}
		}
		if env.Config.Custom.Kg_Refresh_Enable && env.Config.Custom.Kg_token != `` {
			env.Tasker.Add(`kg_refresh`, login_by_token, 86000, true)
		}*/
	})
}
