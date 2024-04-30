package wy

import "net/http"

// 二维码检测扫码状态接口
func LoginQrCheck(key string) (*ReqAnswer, error) {
	res, err := createRequest(
		http.MethodPost,
		`https://music.163.com/weapi/login/qrcode/client/login`,
		map[string]any{
			`key`:  key,
			`type`: 1,
		},
		reqOptions{
			Crypto: `weapi`,
		},
	)
	return res, err
}
