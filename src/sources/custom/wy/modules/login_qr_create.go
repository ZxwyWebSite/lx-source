package wy

// 二维码生成接口
func LoginQrCreate(key string) string {
	return `https://music.163.com/login?codekey=` + key
}
