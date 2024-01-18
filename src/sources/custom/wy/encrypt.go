package wy

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"

	"github.com/ZxwyWebSite/ztool"
	"github.com/ZxwyWebSite/ztool/x/bytesconv"
	"github.com/ZxwyWebSite/ztool/x/json"
	"github.com/ZxwyWebSite/ztool/zcypt"
)

var (
	// __all__ = []string{`weEncrypt`, `linuxEncrypt`, `eEncrypt`}
	// MODULUS = ztool.Str_FastConcat(
	// 	`00e0b509f6259df8642dbc35662901477df22677ec152b5ff68ace615bb7`,
	// 	`b725152b3ab17a876aea8a5aa76d2e417629ec4ee341f56135fccf695280`,
	// 	`104e0312ecbda92557c93870114af6c9d05c4f7f0c3685b7a46bee255932`,
	// 	`575cce10b424d813cfe4875d3e82047b97ddef52741d546b8e289dc6935b`,
	// 	`3ece0462db0a22b8e7`,
	// )
	// PUBKEY   = `010001`
	// NONCE    = bytesconv.StringToBytes(`0CoJUm6Qyw8W8jud`)
	// LINUXKEY = bytesconv.StringToBytes(`rFgB&h#%2?^eDg:Q`)
	eapiKey = bytesconv.StringToBytes(`e82ckenh8dichen8`)
	ivKey   = bytesconv.StringToBytes(`0102030405060708`)
)

func eapiEncrypt(url, text string) map[string][]string {
	digest := zcypt.CreateMD5(bytesconv.StringToBytes(ztool.Str_FastConcat(
		`nobody`, url, `use`, text, `md5forencrypt`,
	)))
	data := ztool.Str_FastConcat(
		url, `-36cd479b6b5-`, text, `-36cd479b6b5-`, digest,
	)
	// 注：JSON编码时会自动将[]byte转为string，这里省去一步转换
	return map[string][]string{
		`params`: {bytesconv.BytesToString(aesEncrypt(bytesconv.StringToBytes(data), eapiKey, false))},
	}
}

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
