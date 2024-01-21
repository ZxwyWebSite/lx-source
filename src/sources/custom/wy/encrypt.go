package wy

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"math/rand"
	_ "unsafe"

	"github.com/ZxwyWebSite/ztool"
	"github.com/ZxwyWebSite/ztool/x/bytesconv"
	"github.com/ZxwyWebSite/ztool/x/json"
	"github.com/ZxwyWebSite/ztool/zcypt"
)

var (
	ivKey       = bytesconv.StringToBytes(`0102030405060708`)
	presetKey   = bytesconv.StringToBytes(`0CoJUm6Qyw8W8jud`)
	linuxapiKey = bytesconv.StringToBytes(`rFgB&h#%2?^eDg:Q`)
	base62      = bytesconv.StringToBytes(`abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789`)
	publicKey   = bytesconv.StringToBytes(`-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDgtQn2JZ34ZC28NWYpAUd98iZ37BUrX/aKzmFbt7clFSs6sXqHauqKWqdtLkF2KexO40H1YTX8z2lSgBBOAxLsvaklV8k4cBFK9snQXE9/DDaFt6Rr7iVZMldczhC0JNgTz+SHXT6CBHuX3e9SdB1Ua44oncaTWz7OBGLbCiK45wIDAQAB
-----END PUBLIC KEY-----`)
	eapiKey = bytesconv.StringToBytes(`e82ckenh8dichen8`)
)

// func eapiEncrypt(url, text string) map[string][]string {
// 	digest := zcypt.CreateMD5(bytesconv.StringToBytes(ztool.Str_FastConcat(
// 		`nobody`, url, `use`, text, `md5forencrypt`,
// 	)))
// 	data := ztool.Str_FastConcat(
// 		url, `-36cd479b6b5-`, text, `-36cd479b6b5-`, digest,
// 	)
// 	// 注：JSON编码时会自动将[]byte转为string，这里省去一步转换
// 	return map[string][]string{
// 		`params`: {bytesconv.BytesToString(aesEncrypt(bytesconv.StringToBytes(data), eapiKey, false))},
// 	}
// }

// crypto.js

func aesEncrypt(text, key []byte, iv bool) []byte {
	pad := 16 - len(text)%16
	text = append(text, bytes.Repeat([]byte{byte(pad)}, pad)...)
	block, _ := aes.NewCipher(key)
	// if err != nil {
	// 	panic(err)
	// }
	var encryptor cipher.BlockMode
	if iv {
		encryptor = cipher.NewCBCEncrypter(block, ivKey)
	} else {
		encryptor = zcypt.NewECBEncrypter(block)
	}
	ciphertext := make([]byte, len(text))
	encryptor.CryptBlocks(ciphertext, text)
	if iv {
		return zcypt.Base64Encode(base64.StdEncoding, ciphertext)
	}
	return bytes.ToUpper(zcypt.HexEncode(ciphertext))
}

//go:linkname rsaEncryptNone crypto/rsa.encrypt
func rsaEncryptNone(*rsa.PublicKey, []byte) ([]byte, error)

func rsaEncrypt(data []byte) string {
	pblock, _ := pem.Decode(publicKey)
	pubKey, _ := x509.ParsePKIXPublicKey(pblock.Bytes)
	// 注：为实现NONE加密手动导出了标准库里的encrypt方法，若编译不过添加以下代码
	// /usr/local/go/src/crypto/rsa/rsa.go:478
	// ```
	// var Encrypt = encrypt // export
	// ```
	// 第二种方式：linkname调用，不用改库 https://www.jianshu.com/p/7b3638b47845
	encData, err := rsaEncryptNone(pubKey.(*rsa.PublicKey), data)
	if err != nil {
		panic(err)
	}
	return zcypt.HexToString(encData)
}

func weapi(object map[string]any) map[string][]string {
	text, err := json.Marshal(object)
	if err != nil {
		panic(err)
	}
	secretKey := make([]byte, 16)
	for i := 0; i < 16; i++ {
		secretKey[i] = base62[rand.Intn(62)]
	}
	return map[string][]string{
		`params`: {bytesconv.BytesToString(aesEncrypt(
			aesEncrypt(text, presetKey, true),
			secretKey,
			true,
		))},
		`encSecKey`: {rsaEncrypt(ztool.Sort_ReverseNew(secretKey))},
	}
}

func linuxapi(object map[string]any) map[string][]string {
	text, err := json.Marshal(object)
	if err != nil {
		panic(err)
	}
	return map[string][]string{
		`eparams`: {bytesconv.BytesToString(aesEncrypt(text, linuxapiKey, false))},
	}
}

func eapi(url string, object map[string]any) map[string][]string {
	text, err := json.Marshal(object)
	if err != nil {
		panic(err)
	}
	message := ztool.Str_FastConcat(
		`nobody`, url, `use`, bytesconv.BytesToString(text), `md5forencrypt`,
	)
	digest := zcypt.CreateMD5(bytesconv.StringToBytes(message))
	data := bytes.Join(
		[][]byte{
			bytesconv.StringToBytes(url),
			text,
			bytesconv.StringToBytes(digest),
		},
		[]byte{45, 51, 54, 99, 100, 52, 55, 57, 98, 54, 98, 53, 45},
	)
	return map[string][]string{
		`params`: {bytesconv.BytesToString(aesEncrypt(data, eapiKey, false))},
	}
}

func decrypt(data []byte) (out []byte) {
	dec, err := zcypt.HexDecode(data)
	if err == nil {
		out, err = zcypt.AesDecrypt(dec, eapiKey)
	}
	if err != nil {
		panic(err)
	}
	return out
}
