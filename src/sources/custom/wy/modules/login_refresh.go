package wy

import (
	"net/http"
	"strings"
)

// 登录刷新
func LoginRefresh(query ReqQuery) (*ReqAnswer, error) {
	res, err := createRequest(
		http.MethodPost,
		`https://music.163.com/weapi/login/token/refresh`,
		map[string]any{},
		reqOptions{
			Crypto: `weapi`,
			UA:     `pc`,
			Cookie: query.Cookie,
			RealIP: query.RealIP,
		},
	)
	if code, ok := res.Body[`code`].(int); ok && err == nil {
		if code == 200 {
			res.Body[`cookie`] = strings.Join(res.Cookie, `;`)
		}
	}
	return res, err
}
