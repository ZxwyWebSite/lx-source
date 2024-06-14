package wy

import (
	"errors"
	"io"
	"math/rand"
	"net/http"
	stdurl "net/url"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/ZxwyWebSite/ztool"
	"github.com/ZxwyWebSite/ztool/x/json"
	"github.com/ZxwyWebSite/ztool/zcypt"
)

// request.js

const anonymous_token = `1f5fa7b6a6a9f81a11886e5186fde7fb98e25cf0036d7afd055b980b2261f5464b7f5273fc3921d1262bfec66a19a30c41d8da00c3685f5ace96f0d5a48b6db334d974731083682e3324751bcc9aaf44c3061cd1`

var (
	wapiReg = regexp.MustCompile(`\w*api`)
	csrfReg = regexp.MustCompile(`_csrf=([^(;|$)]+)`)
	domaReg = regexp.MustCompile(`\s*Domain=[^(;|$)]+;*`)
)

var userAgentMap = map[string]string{
	`mobile`: `Mozilla/5.0 (iPhone; CPU iPhone OS 17_2_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.2 Mobile/15E148 Safari/604.1`,
	`pc`:     `Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36 Edg/120.0.0.0`,
}

type (
	ReqQuery struct {
		Cookie map[string]string
		RealIP string
		Ids    string
		Br     string
		Level  string
	}
	reqOptions struct {
		Headers map[string]string
		UA      string
		RealIP  string
		IP      string
		Crypto  string
		Url     string
		Cookie  interface{}
	}
	ReqAnswer struct {
		Status int
		Body   map[string]any
		Cookie []string
	}
)

func createRequest(method, url string, data map[string]any, options reqOptions) (*ReqAnswer, error) {
	if options.Headers == nil {
		options.Headers = make(map[string]string)
	}
	options.Headers[`User-Agent`] = userAgentMap[options.UA]
	if method == http.MethodPost {
		options.Headers[`Content-Type`] = `application/x-www-form-urlencoded`
	}
	if strings.Contains(url, `music.163.com`) {
		options.Headers[`Referer`] = `https://music.163.com`
	}
	ip := ztool.Str_Select(options.RealIP, options.IP, ``)
	if ip != `` {
		options.Headers[`X-Real-IP`] = ip
		options.Headers[`X-Forwarded-For`] = ip
	}
	if obj, ok := options.Cookie.(map[string]string); ok {
		obj[`__remember_me`] = `true`
		obj[`_ntes_nuid`] = zcypt.HexToString(zcypt.RandomBytes(16))
		if !strings.Contains(url, `login`) {
			obj[`NMTID`] = zcypt.HexToString(zcypt.RandomBytes(16))
		}
		if _, ok := obj[`MUSIC_U`]; !ok {
			// 游客
			if _, ok := obj[`MUSIC_A`]; !ok {
				obj[`MUSIC_A`] = anonymous_token
				if obj[`os`] == `` {
					obj[`os`] = `ios`
				}
				if obj[`appver`] == `` {
					obj[`appver`] = `8.20.21`
				}
			}
		}
		keys, i := make([]string, len(obj)), 0
		for k, v := range obj {
			keys[i] = ztool.Str_FastConcat(
				stdurl.QueryEscape(k), `=`, stdurl.QueryEscape(v),
			)
			i++
		}
		options.Headers[`Cookie`] = strings.Join(keys, `; `)
		options.Cookie = obj
	} else if str, ok := options.Cookie.(string); ok && str != `` {
		options.Headers[`Cookie`] = str
	} else {
		options.Headers[`Cookie`] = `__remember_me=true; NMTID=xxx`
	}
	var form stdurl.Values
	switch options.Crypto {
	case `weapi`:
		options.Headers[`User-Agent`] = `Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/116.0.0.0 Safari/537.36 Edg/116.0.1938.69`
		csrfToken := csrfReg.FindStringSubmatch(options.Headers[`Cookie`])
		if len(csrfToken) > 1 {
			data[`csrf_token`] = csrfToken[1]
		} else {
			data[`csrf_token`] = ``
		}
		form = weapi(data)
		// fmt.Println(form.Encode())
		url = wapiReg.ReplaceAllString(url, `weapi`)
	case `linuxapi`:
		form = linuxapi(map[string]any{
			`method`: method,
			`url`:    wapiReg.ReplaceAllString(url, `weapi`),
			`params`: data,
		})
		options.Headers[`User-Agent`] = `Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/60.0.3112.90 Safari/537.36`
		url = `https://music.163.com/api/linux/forward`
	case `eapi`:
		cookie, ok := options.Cookie.(map[string]string)
		if !ok {
			cookie = make(map[string]string)
		}
		// csrfToken := cookie[`__csrf`]
		now := time.Now()
		reqid := strconv.Itoa(rand.Intn(1000))
		header := map[string]string{
			`osver`:       ztool.Str_Select(cookie[`osver`], `17,1,2`),    //系统版本
			`deviceId`:    cookie[`deviceId`],                             //zcypt.Base64ToString(base64.StdEncoding, bytesconv.StringToBytes(imei+"'\t02:00:00:00:00:00\t5106025eb79a5247\t70ffbaac7'"))
			`appver`:      ztool.Str_Select(cookie[`appver`], `8.9.70`),   // app版本
			`versioncode`: ztool.Str_Select(cookie[`versioncode`], `140`), //版本号
			`mobilename`:  cookie[`mobilename`],                           //设备model
			`buildver`:    ztool.Str_Select(cookie[`buildver`], strconv.FormatInt(now.Unix(), 10)),
			`resolution`:  ztool.Str_Select(cookie[`resolution`], `1920x1080`), //设备分辨率
			`__csrf`:      cookie[`__csrf`],                                    //csrfToken,
			`os`:          ztool.Str_Select(cookie[`os`], `ios`),
			`channel`:     cookie[`channel`],
			`requestId`: ztool.Str_FastConcat(
				strconv.FormatInt(now.UnixMilli(), 10), `_`,
				strings.Repeat(`0`, 4-len(reqid)), reqid,
			),
		}
		if cookie[`MUSIC_U`] != `` {
			header[`MUSIC_U`] = cookie[`MUSIC_U`]
		}
		if cookie[`MUSIC_A`] != `` {
			header[`MUSIC_A`] = cookie[`MUSIC_A`]
		}
		keys, i := make([]string, len(header)), 0
		for k, v := range header {
			keys[i] = ztool.Str_FastConcat(
				stdurl.QueryEscape(k), `=`, stdurl.QueryEscape(v),
			)
			i++
		}
		options.Headers[`Cookie`] = strings.Join(keys, `; `)
		out, err := json.Marshal(header)
		if err != nil {
			panic(err)
		}
		data[`header`] = out //bytesconv.BytesToString(out)
		form = eapi(options.Url, data)
		url = wapiReg.ReplaceAllString(url, `eapi`)
	default:
		return nil, errors.New(`not support`)
	}
	// values := stdurl.Values{}
	// for k, v := range data {
	// 	values.Add(k, v)
	// }
	answer := ReqAnswer{Status: 500, Body: map[string]any{} /*, Cookie: []string{}*/}
	err := ztool.Net_Request(method, url,
		strings.NewReader(form.Encode()),
		[]ztool.Net_ReqHandlerFunc{
			ztool.Net_ReqAddHeader(options.Headers),
		},
		[]ztool.Net_ResHandlerFunc{
			func(res *http.Response) error {
				body, err := io.ReadAll(res.Body)
				if err == nil {
					// fmt.Println(`body:`, string(body), "\nstr:", body)
					if len(body) == 0 {
						err = errors.New(`nil Body`)
					}
					if err == nil {
						answer.Cookie = res.Header[`Set-Cookie`] //res.Header.Values(`set-cookie`)
						for i, v := range answer.Cookie {
							answer.Cookie[i] = domaReg.ReplaceAllString(v, ``)
						}
						if options.Crypto == `eapi` && body[0] != '{' {
							err = json.Unmarshal(decrypt(body), &answer.Body)
						} else {
							err = json.Unmarshal(body, &answer.Body)
						}
						if code, ok := answer.Body[`code`].(string); ok {
							answer.Body[`code`], err = strconv.Atoi(code)
						} else {
							answer.Body[`code`] = res.StatusCode
						}
						if err == nil {
							if code, ok := answer.Body[`code`].(int); ok {
								if !ztool.Chk_IsMatchInt(code, 201, 302, 400, 502, 800, 801, 802, 803) {
									// 特殊状态码
									answer.Status = 200
								}
							}
							if answer.Status < 100 || answer.Status >= 600 {
								answer.Status = 400
							}
							if answer.Status != 200 {
								err = errors.New(strconv.Itoa(answer.Status))
							}
						}
					}
				}
				if err != nil {
					answer.Status = 502
					answer.Body = map[string]any{`code`: 502, `msg`: err.Error()}
					// return err
				}
				return err // nil
			},
		},
	)
	return &answer, err
}
