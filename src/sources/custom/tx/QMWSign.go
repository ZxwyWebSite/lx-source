package tx

import (
	"crypto/md5"
	"encoding/hex"
	"regexp"
	"strings"

	"github.com/ZxwyWebSite/ztool"
	"github.com/ZxwyWebSite/ztool/x/bytesconv"
)

func v(b string) string {
	res := []byte{}
	p := [...]int{21, 4, 9, 26, 16, 20, 27, 30}
	for _, x := range p {
		res = append(res, b[x])
	}
	return bytesconv.BytesToString(res) //string(res)
}

func c(b string) string {
	res := []byte{}
	p := [...]int{18, 11, 3, 2, 1, 7, 6, 25}
	for _, x := range p {
		res = append(res, b[x])
	}
	return bytesconv.BytesToString(res) //string(res)
}

func y(a, b, c int) (e []int) {
	// e := []int{}
	r25 := a >> 2
	if b != 0 && c != 0 {
		r26 := a & 3
		r26_2 := r26 << 4
		r26_3 := b >> 4
		r26_4 := r26_2 | r26_3
		r27 := b & 15
		r27_2 := r27 << 2
		r27_3 := r27_2 | (c >> 6)
		r28 := c & 63
		e = append(e, r25)
		e = append(e, r26_4)
		e = append(e, r27_3)
		e = append(e, r28)
	} else {
		r10 := a >> 2
		r11 := a & 3
		r11_2 := r11 << 4
		e = append(e, r10)
		e = append(e, r11_2)
	}
	return //e
}

func n(ls []int) string {
	e := []int{}
	for i, r := 0, len(ls); i < r; i += 3 {
		if i < r-2 {
			e = append(e, y(ls[i], ls[i+1], ls[i+2])...)
		} else {
			e = append(e, y(ls[i], 0, 0)...)
		}
	}
	res := []byte{}
	b64all := `ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/=`
	for _, i := range e {
		res = append(res, b64all[i])
	}
	return bytesconv.BytesToString(res)
}

func t(b string) (res []int) {
	zd := map[string]int{
		`0`: 0,
		`1`: 1,
		`2`: 2,
		`3`: 3,
		`4`: 4,
		`5`: 5,
		`6`: 6,
		`7`: 7,
		`8`: 8,
		`9`: 9,
		`A`: 10,
		`B`: 11,
		`C`: 12,
		`D`: 13,
		`E`: 14,
		`F`: 15,
	}
	ol := [...]int{212, 45, 80, 68, 195, 163, 163, 203, 157, 220, 254, 91, 204, 79, 104, 6}
	// res := []int{}
	j := 0
	for i, r := 0, len(b); i < r; i += 2 {
		one := zd[string(b[i])]
		two := zd[string(b[i+1])]
		r := one*16 ^ two
		// if j >= 16 {
		// 	break
		// }
		res = append(res, r^ol[j])
		j++
	}
	return //res
}

func createMD5(s []byte) string {
	hash := md5.New()
	hash.Write(s)
	return hex.EncodeToString(hash.Sum(nil))
}

func sign(params []byte) string {
	md5Str := strings.ToUpper(createMD5(params))
	h := v(md5Str)
	e := c(md5Str)
	ls := t(md5Str)
	m := n(ls)
	res := ztool.Str_FastConcat(`zzb`, h, m, e) //`zzb` + h + m + e
	res = strings.ToLower(res)
	r := regexp.MustCompile(`[\/+]`)
	res = r.ReplaceAllString(res, ``)
	return res
}
