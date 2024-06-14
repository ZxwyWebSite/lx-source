package tx

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"lx-source/src/env"
	"math/rand"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strconv"
	"strings"
	"time"
	_ "unsafe"

	"github.com/ZxwyWebSite/ztool"
	"github.com/ZxwyWebSite/ztool/logs"
	"github.com/ZxwyWebSite/ztool/x/json"
	"github.com/google/uuid"
)

//go:linkname request github.com/ZxwyWebSite/ztool.request
func request(client *http.Client, method, url string, body io.Reader, reqh []ztool.Net_ReqHandlerFunc, resh []ztool.Net_ResHandlerFunc) error

// QQ快速登录 - 直接使用本机已登录账号
func Qlogin_graph(l *logs.Logger) error {
	// 参考文章: https://learnku.com/articles/33970

	jar, _ := cookiejar.New(nil)
	client := &http.Client{
		Timeout: time.Second * 10,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
		Jar: jar,
	}

	// Step0 - 获取本机QQ服务地址
	// (应该没有人会占用4301端口导致qq切换备用端口吧，先不写了...懒)
	err := request(
		client, http.MethodGet, `https://localhost.ptlogin2.qq.com:4301/pc_querystatus`, nil,
		[]ztool.Net_ReqHandlerFunc{ztool.Net_ReqAddHeaders()},
		[]ztool.Net_ResHandlerFunc{func(res *http.Response) error {
			_, err := io.Copy(io.Discard, res.Body)
			return err
		}},
	)
	if err != nil {
		return errors.New(`step0: 无法连接本机QQ服务`)
	}

	// Step1 - 获取 pt_local_token
	var pt_local_token string
	err = request(
		client, http.MethodGet,
		`https://xui.ptlogin2.qq.com/cgi-bin/xlogin?appid=716027609&daid=383&style=33&login_text=%E7%99%BB%E5%BD%95&hide_title_bar=1&hide_border=1&target=self&s_url=https://graph.qq.com/oauth2.0/login_jump&pt_3rd_aid=100497308&pt_feedback_link=https://support.qq.com/products/77942?customInfo=.appid100497308&theme=2&verify_theme=`,
		nil,
		[]ztool.Net_ReqHandlerFunc{ztool.Net_ReqAddHeaders(map[string]string{
			`Referer`: `https://graph.qq.com/`,
		})},
		[]ztool.Net_ResHandlerFunc{func(res *http.Response) error {
			for _, v := range res.Cookies() {
				if v.Name == `pt_local_token` {
					pt_local_token = v.Value
					break
				}
			}
			_, err := io.Copy(io.Discard, res.Body)
			return err
		}},
	)
	if err != nil {
		return err
	}
	l.Info(`pt_local_token: %v`, pt_local_token)

	// Step2 - 获取本机登录的 QQ 号
	var url2 strings.Builder
	url2.WriteString(`https://localhost.ptlogin2.qq.com:4301/pt_get_uins?callback=ptui_getuins_CB&r=`)
	url2.WriteString(strconv.FormatFloat(rand.Float64(), 'f', -1, 64))
	url2.WriteString(`&pt_local_tk=`)
	url2.WriteString(pt_local_token)
	var out2 []struct {
		Uin int `json:"uin"`
		// FaceIndex  int    `json:"face_index"`
		// Gender     int    `json:"gender"`
		Nickname string `json:"nickname"`
		// ClientType int    `json:"client_type"`
		// UinFlag    int    `json:"uin_flag"`
		// Account    int    `json:"account"`
	}
	var header2 = map[string]string{
		`Referer`: `https://xui.ptlogin2.qq.com/`,
		// `Cookie`:  `pt_local_token=` + pt_local_token,
	}
	err = request(
		client, http.MethodGet, url2.String(), nil,
		[]ztool.Net_ReqHandlerFunc{ztool.Net_ReqAddHeaders(header2)},
		[]ztool.Net_ResHandlerFunc{func(res *http.Response) error {
			data, err := io.ReadAll(res.Body)
			if err != nil {
				return err
			}
			sep_b := bytes.IndexByte(data, '[')
			sep_e := bytes.LastIndexByte(data, ']')
			if sep_b == -1 || sep_e == -1 {
				return errors.New(`step2: 无法解析返回数据`)
			}
			sep := data[sep_b : sep_e+1]
			return json.Unmarshal(sep, &out2)
		}},
	)
	if err != nil {
		return err
	}
	var uin string
	length := len(out2)
	switch length {
	case 0:
		return errors.New(`step2: 无可用账号`)
	case 1:
		uin = strconv.Itoa(out2[0].Uin)
	default:
		fmt.Println(`请选择要登录的账号:`)
		for i, v := range out2 {
			fmt.Println(i, v.Nickname, v.Uin)
		}
		for {
			fmt.Print(`输入序号: `)
			var input string
			fmt.Scanln(&input)
			i, err := strconv.Atoi(input)
			if err != nil {
				l.Error(`err: %v`, err)
				continue
			}
			if i >= length {
				l.Error(`err: 下标越界`)
				continue
			}
			uin = strconv.Itoa(out2[i].Uin)
			break
		}
	}
	l.Info(`uin: %v`, uin)

	// Step3 - 获取 clientkey
	var url3 strings.Builder
	url3.WriteString(`https://localhost.ptlogin2.qq.com:4301/pt_get_st?clientuin=`)
	url3.WriteString(uin)
	url3.WriteString(`&r=`)
	url3.WriteString(strconv.FormatFloat(rand.Float64(), 'f', -1, 64))
	url3.WriteString(`&pt_local_tk=`)
	url3.WriteString(pt_local_token)
	url3.WriteString(`&callback=__jp0`)
	// var clientkey string
	err = request(
		client, http.MethodGet, url3.String(), nil,
		[]ztool.Net_ReqHandlerFunc{ztool.Net_ReqAddHeaders(header2)},
		[]ztool.Net_ResHandlerFunc{func(res *http.Response) error {
			/*for _, v := range res.Cookies() {
				if v.Name == `clientkey` {
					clientkey = v.Value
					break
				}
			}*/
			_, err := io.Copy(io.Discard, res.Body)
			return err
		}},
	)
	if err != nil {
		return err
	}

	// Step4 - 获取 skey
	var url4 strings.Builder
	url4.WriteString(`https://ssl.ptlogin2.qq.com/jump?clientuin=`)
	url4.WriteString(uin)
	url4.WriteString(`&keyindex=9&pt_aid=716027609&daid=383&u1=https://graph.qq.com/oauth2.0/login_jump&pt_local_tk=`)
	url4.WriteString(pt_local_token)
	url4.WriteString(`&pt_3rd_aid=100497308&ptopt=1&style=40`)
	/*var cookie4 strings.Builder
	cookie4.WriteString(`pt_local_token=`)
	cookie4.WriteString(pt_local_token)
	cookie4.WriteByte(';')
	cookie4.WriteString(`clientuin=`)
	cookie4.WriteString(uin)
	cookie4.WriteByte(';')
	cookie4.WriteString(`clientkey=`)
	cookie4.WriteString(clientkey)*/
	var jurl string
	err = request(
		client, http.MethodGet, url4.String(), nil,
		[]ztool.Net_ReqHandlerFunc{ztool.Net_ReqAddHeaders( /*map[string]string{
				`Referer`: `https://xui.ptlogin2.qq.com/`,
				`Cookie`:  cookie4.String(),
			}*/header2)},
		[]ztool.Net_ResHandlerFunc{func(res *http.Response) error {
			data, err := io.ReadAll(res.Body)
			if err != nil {
				return err
			}
			sep_b := bytes.IndexByte(data, ',')
			sep_e := bytes.LastIndexByte(data, ' ')
			if sep_b == -1 || sep_e == -1 {
				return errors.New(`step4: 无法解析返回数据`)
			}
			jurl = string(data[sep_b+3 : sep_e-2])
			return nil
		}},
	)
	if err != nil {
		return err
	}

	// Step5 - 获取 p_skey
	var p_skey string
	err = request(
		client, http.MethodGet, jurl, nil,
		[]ztool.Net_ReqHandlerFunc{ztool.Net_ReqAddHeaders(header2)},
		[]ztool.Net_ResHandlerFunc{func(res *http.Response) error {
			for _, v := range res.Cookies() {
				if v.Name == `p_skey` && v.Value != `` {
					p_skey = v.Value
					break
				}
			}
			_, err := io.Copy(io.Discard, res.Body)
			return err
		}},
	)
	if err != nil {
		return err
	}

	// Step6 - 登录账号
	getGtk := func(skey string) string {
		var hash = 5381
		for _, v := range skey {
			hash += (hash << 5) + int(v)
		}
		return strconv.Itoa(hash & 0x7fffffff)
	}
	now := time.Now()
	var authcode string
	err = request(
		client, http.MethodPost,
		`https://graph.qq.com/oauth2.0/authorize`,
		strings.NewReader(ztool.Str_FastConcat(
			`response_type=code&client_id=100497308&redirect_uri=https%3A%2F%2Fy.qq.com%2Fportal%2Fwx_redirect.html%3Flogin_type%3D1%26surl%3Dhttps%3A%2F%2Fy.qq.com%2F&scope=get_user_info%2Cget_app_friends&state=state&switch=&from_ptlogin=1&src=1&update_auth=1&openapi=1010_1030`,
			`&g_tk=`, getGtk(p_skey),
			`&auth_time=`, strconv.FormatInt(now.UnixMilli(), 10),
			`&ui=`, strings.ToUpper(uuid.NewString()),
		)),
		[]ztool.Net_ReqHandlerFunc{ztool.Net_ReqAddHeaders(map[string]string{
			`Referer`:      `https://graph.qq.com/oauth2.0/show?which=Login&display=pc&response_type=code&client_id=100497308&redirect_uri=https%3A%2F%2Fy.qq.com%2Fportal%2Fwx_redirect.html%3Flogin_type%3D1%26surl%3Dhttps%3A%2F%2Fy.qq.com%2F&state=state&display=pc&scope=get_user_info%2Cget_app_friends`,
			`Content-Type`: `application/x-www-form-urlencoded`,
		})},
		[]ztool.Net_ResHandlerFunc{func(res *http.Response) error {
			/*if res.StatusCode != 302 {
				return errors.New(`step6: not redirect`)
			}*/
			location := res.Header[`Location`][0]
			l.Info(`loc: %v`, location)
			loc, err := url.Parse(location)
			if err != nil {
				return err
			}
			authcode = loc.Query()[`code`][0]
			return nil
		}},
	)
	if err != nil {
		return err
	}
	l.Info(`authcode: %v`, authcode)
	var out6 struct {
		Code int `json:"code"`
		// Ts      int64  `json:"ts"`
		// StartTs int64  `json:"start_ts"`
		// Traceid string `json:"traceid"`
		Req struct {
			Code int `json:"code"`
			Data struct {
				// Openid             string        `json:"openid"`
				RefreshToken string `json:"refresh_token"`
				AccessToken  string `json:"access_token"`
				ExpiredAt    int    `json:"expired_at"`
				// Musicid      int    `json:"musicid"`
				Musickey string `json:"musickey"`
				// MusickeyCreateTime int           `json:"musickeyCreateTime"`
				// FirstLogin         int           `json:"first_login"`
				// ErrMsg             string        `json:"errMsg"`
				// SessionKey         string        `json:"sessionKey"`
				// Unionid            string        `json:"unionid"`
				StrMusicid string `json:"str_musicid"`
				// Errtip             string        `json:"errtip"`
				// Nick               string        `json:"nick"`
				// Logo               string        `json:"logo"`
				// FeedbackURL        string        `json:"feedbackURL"`
				// EncryptUin         string        `json:"encryptUin"`
				// Userip             string        `json:"userip"`
				// LastLoginTime      int           `json:"lastLoginTime"`
				// KeyExpiresIn       int           `json:"keyExpiresIn"`
				// RefreshKey         string        `json:"refresh_key"`
				// LoginType          int           `json:"loginType"`
				// Prompt2Bind        int           `json:"prompt2bind"`
				// LogoffStatus       int           `json:"logoffStatus"`
				// OtherAccounts      []interface{} `json:"otherAccounts"`
				// OtherPhoneNo       string        `json:"otherPhoneNo"`
				// Token              string        `json:"token"`
				// IsPrized           int           `json:"isPrized"`
				// IsShowDevManage    int           `json:"isShowDevManage"`
				// ErrTip2            string        `json:"errTip2"`
				// Tip3               string        `json:"tip3"`
				// EncryptedPhoneNo   string        `json:"encryptedPhoneNo"`
				// PhoneNo            string        `json:"phoneNo"`
				// BindAccountType    int           `json:"bindAccountType"`
				// NeedRefreshKeyIn   int           `json:"needRefreshKeyIn"`
			} `json:"data"`
		} `json:"req"`
	}
	err = request(
		client, http.MethodPost,
		`https://u.y.qq.com/cgi-bin/musicu.fcg`,
		strings.NewReader(ztool.Str_FastConcat(
			`{"comm":{"g_tk":5381,"platform":"yqq","ct":24,"cv":0},"req":{"module":"QQConnectLogin.LoginServer","method":"QQLogin","param":{"code":"`, authcode, `"}}}`,
		)),
		[]ztool.Net_ReqHandlerFunc{ztool.Net_ReqAddHeaders(map[string]string{
			`Referer`:      `https://y.qq.com/`,
			`Content-Type`: `application/x-www-form-urlencoded`,
		})},
		[]ztool.Net_ResHandlerFunc{ztool.Net_ResToStruct(&out6)},
	)
	if err != nil {
		return err
	}
	l.Info(`res: %+v`, out6)
	l.Info(`登录成功`)

	env.Config.Custom.Tx_Enable = true
	env.Config.Custom.Tx_Uuin = out6.Req.Data.StrMusicid
	env.Config.Custom.Tx_Ukey = out6.Req.Data.Musickey
	env.Config.Custom.Tx_Refresh_Enable = false
	// env.Config.Custom.Tx_Refresh_Interval = time.Date(now.Year(), now.Month(), now.Day()+5, 0, 0, 0, 0, now.Location()).Unix()
	// env.Config.Custom.Tx_RefreshToken = out6.Req.Data.RefreshToken
	// env.Config.Custom.Tx_AccessToken = out6.Req.Data.AccessToken

	return env.Cfg.Save(``)
}

// QQ扫码登录(todo)
// func qlogin_qr_()
