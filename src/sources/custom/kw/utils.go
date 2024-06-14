package kw

import (
	"bytes"
	"lx-source/src/sources"

	"github.com/ZxwyWebSite/ztool/x/bytesconv"
)

var (
	fileInfo = map[string]struct {
		E string // 扩展名
		H string // 专用音质
	}{
		sources.Q_128k: {
			E: sources.X_mp3,
			H: sources.Q_128k,
		},
		sources.Q_320k: {
			E: sources.X_mp3,
			H: sources.Q_320k,
		},
		sources.Q_flac: {
			E: sources.Q_flac,
			H: `2000k`,
		},
		sources.Q_fl24: {
			E: sources.Q_flac,
			H: `4000k`,
		},
	}
	// 注：这个还是有规律的，加上或去掉k即可直接比较
	// qualityMapReverse = map[string]string{
	// 	`128`:  sources.Q_128k,
	// 	`320`:  sources.Q_320k,
	// 	`2000`: sources.Q_flac,
	// 	`4000`: sources.Q_fl24,
	// }
	desheader = map[string]string{
		// `User-Agent`: `okhttp/3.10.0`,
	}
	bdheader = map[string]string{
		`channel`: `guanfang`,
		`plat`:    `ar`,
		`net`:     `wifi`,
		`ver`:     `3.1.4`,
		`api-ver`: `application/json`,

		`user-agent`: `Dart/2.18 (dart:io)`, //`Dalvik/2.1.0 (Linux; U; Android 7.1.1; OPPO R9sk Build/NMF26F)`,
	}
	// bdsreg = regexp.MustCompile(`[^a-zA-Z0-9]`)
)

func mkMap(data []byte) map[string]string {
	out := make(map[string]string)
	sep := bytes.Split(data, []byte{13, 10})
	for i, r := 0, len(sep); i < r; i++ {
		var s = sep[i]
		if p := bytes.IndexByte(s, '='); p != -1 {
			out[bytesconv.BytesToString(s[:p])] = bytesconv.BytesToString(s[p+1:])
			continue
		} else {
			out[`_`] += bytesconv.BytesToString(s) + `;`
		}
		/*pat := bytes.Split(sep[i], []byte{61})
		if len(pat) >= 2 {
			out[bytesconv.BytesToString(pat[0])] = bytesconv.BytesToString(pat[1])
			continue
		}
		out[`_`] += bytesconv.BytesToString(pat[0]) + `;`*/
	}
	return out
}

// 波点签名算法
/*func Bdsign(str string, m, m2 map[string]string) *strings.Builder {
	var b strings.Builder
	b.WriteString(`uid=`)
	b.WriteString(env.Config.Custom.Kw_Bd_Uid)
	b.WriteByte('&')
	b.WriteString(`token=`)
	b.WriteString(env.Config.Custom.Kw_Bd_Token)
	b.WriteByte('&')
	b.WriteString(`timestamp=`)
	b.WriteString(strconv.FormatInt(time.Now().UnixMilli(), 10))
	for k, v := range m2 {
		b.WriteByte('&')
		b.WriteString(k)
		b.WriteByte('=')
		b.WriteString(url.QueryEscape(v))
	}
	// 取 strings.Builder.buf []byte 地址
	pb := (*[]byte)(unsafe.Pointer(uintptr(unsafe.Pointer(&b)) + unsafe.Sizeof((*strings.Builder)(nil))))
	charArray := bdsreg.ReplaceAll(*pb, []byte{})
	slices.Sort(charArray)
	str3 := string(charArray)
	fmt.Println(str3)
	lowerCase := zcypt.MD5EncStr(`kuwotest` + str3 + `/api/play/music/v2/audioUrl`)
	b.WriteByte('&')
	b.WriteString(`sign=`)
	b.WriteString(lowerCase)
	return &b
}*/
