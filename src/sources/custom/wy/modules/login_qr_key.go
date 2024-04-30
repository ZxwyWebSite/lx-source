package wy

import "net/http"

// 二维码 key 生成接口
func LoginQrKey() (*ReqAnswer, error) {
	res, err := createRequest(
		http.MethodPost,
		`https://music.163.com/weapi/login/qrcode/unikey`,
		map[string]any{
			`type`: 1,
		},
		reqOptions{
			Crypto: `weapi`,
			Cookie: nil,
		},
	)
	return res, err
}
