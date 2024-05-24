package kg

import (
	"lx-source/src/env"
	"lx-source/src/sources"
	"strings"

	"github.com/ZxwyWebSite/ztool"
	"github.com/ZxwyWebSite/ztool/x/slices"
	"github.com/ZxwyWebSite/ztool/zcypt"
)

var (
	// qualityHashMap = map[string]string{
	// 	sources.Q_128k: `hash_128`,
	// 	sources.Q_320k: `hash_320`,
	// 	sources.Q_flac: `hash_flac`,
	// 	sources.Q_fl24: `hash_high`,
	// }
	qualityMap = map[string]string{
		sources.Q_128k: `128`,
		sources.Q_320k: `320`,
		sources.Q_flac: sources.Q_flac,
		sources.Q_fl24: `high`,

		sources.Q_master: `viper_atmos`,
	}
)

const (
	// signkey   = `OIlwieks28dk2k092lksi2UIkp`
	// pidversec = `57ae12eb6890223e355ccfcb74edf70d`
	// clientver = `12029`
	url = `https://gateway.kugou.com/v5/url`
	// appid     = `1005`
	mid = `211008`
)

func sortDict(dictionary map[string]string) ([]string, int) {
	length := len(dictionary)
	var keys = make([]string, 0, length)
	for k := range dictionary {
		keys = append(keys, k)
	}
	slices.Sort(keys)
	return keys, length
}

// func sign(params map[string]string, body string) string {
// 	keys, lens := sortDict(params)
// 	var b strings.Builder
// 	for i := 0; i < lens; i++ {
// 		b.WriteString(keys[i])
// 		b.WriteByte('=')
// 		b.WriteString(params[keys[i]])
// 	}
// 	// b.WriteString(body)
// 	return zcypt.MD5EncStr(ztool.Str_FastConcat(signkey, b.String(), signkey))
// }

func signRequest(method string, url string, body string, params, headers map[string]string, out any) error {
	keys, lens := sortDict(params)
	// buildSignatureParams
	var b strings.Builder
	for i := 0; i < lens; i++ {
		b.WriteString(keys[i])
		b.WriteByte('=')
		b.WriteString(params[keys[i]])
	}
	b.WriteString(body)
	// buildRequestParams
	var c strings.Builder
	for j := 0; j < lens; j++ {
		c.WriteString(keys[j])
		c.WriteByte('=')
		c.WriteString(params[keys[j]])
		c.WriteByte('&')
	}
	c.WriteString(`signature`)
	c.WriteByte('=')
	c.WriteString(zcypt.MD5EncStr(ztool.Str_FastConcat(
		env.Config.Custom.Kg_Client_SignKey,
		b.String(), env.Config.Custom.Kg_Client_SignKey,
	)))

	url = ztool.Str_FastConcat(url, `?`, c.String())
	// ztool.Cmd_FastPrintln(url)
	return ztool.Net_Request(
		method, url, strings.NewReader(body),
		[]ztool.Net_ReqHandlerFunc{ztool.Net_ReqAddHeader(headers)},
		[]ztool.Net_ResHandlerFunc{ //func(res *http.Response) error {
			// body, err := io.ReadAll(res.Body)
			// fmt.Printf("%s, %s, %s\n", body, err, res.Status)
			// return ztool.Err_EsContinue
			// },
			ztool.Net_ResToStruct(out),
		},
	)
}

func getKey(hash_ string) string {
	return zcypt.MD5EncStr(ztool.Str_FastConcat(
		strings.ToLower(hash_), env.Config.Custom.Kg_Client_PidVerSec,
		env.Config.Custom.Kg_Client_AppId, mid, env.Config.Custom.Kg_userId,
	))
}

// 解析版权字段 (Copilot)
/*
// 定义一个函数，接受一个整数作为参数，返回一个布尔值的切片，表示每一位的状态
 依次为: 下载是否付费 \ 下载是否禁止 \ 播放是否付费 \ 播放是否禁止
https://open.kugou.com/docs/open-player/#/android-sdk?v=1&id=%e6%ad%8c%e6%9b%b2%e5%ad%97%e6%ae%b5%e8%a7%a3%e6%9e%90
*/
// func parsePrivilege(privilege int) []bool {
// 	result := make([]bool, 4) // 创建一个长度为4的切片
// 	for i := 0; i < 4; i++ {
// 		// 用位运算符&来判断每一位是否为1，如果是则将对应的切片元素设为true
// 		if privilege&(1<<i) != 0 {
// 			result[i] = true
// 		}
// 	}
// 	return result
// }
