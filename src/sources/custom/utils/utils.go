package utils

import (
	"github.com/ZxwyWebSite/ztool"
)

// func SizeFormat(size int) string {
// 	if size < 1024 {
// 		return ztool.Str_FastConcat(strconv.Itoa(size), `B`)
// 	}
// 	size64 := float64(size)
// 	if size64 < math.Pow(size64, 2) {

// 	}
// 	return ``
// }

// 删除?号后尾随内容
func DelQuery(str string) string {
	return ztool.Str_Before(str, `?`)
}
