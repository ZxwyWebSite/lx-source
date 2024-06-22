//go:build ignore

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// 获取版本号
func version() string {
	fenv, _ := os.Open(`src/env/env.go`)
	benv := bufio.NewReader(fenv)
	var ever string
	for {
		line, _, _ := benv.ReadLine()
		length := len(line)
		if length == 0 {
			continue
		}
		sline := string(line)
		if strings.HasPrefix(sline, `	Version`) {
			ever = `v` + sline[12:length-1]
			break
		}
	}
	fenv.Close()
	if ever == `` {
		panic(`No Version`)
	} else {
		return ever
	}
}

// 生成更新日志
func changelog(ever string) string {
	fupd, _ := os.Open(`update.md`)
	bupd := bufio.NewReader(fupd)
	var eupd strings.Builder
	eupd.WriteString(`### 更新内容：`)
	eupd.WriteByte('\n')
	for {
		line, _, _ := bupd.ReadLine()
		length := len(line)
		if length == 0 {
			continue
		}
		if strings.Contains(string(line), ever) {
			for {
				lline, _, _ := bupd.ReadLine()
				length := len(lline)
				if length == 0 {
					break
				}
				eupd.WriteString(string(lline))
				eupd.WriteByte('\n')
			}
			break
		}
	}
	fupd.Close()
	eupd.WriteByte('\n')
	eupd.WriteString(`### CDN加速下载：`)
	eupd.WriteByte('\n')
	for _, v := range []string{
		`lx-source-android-arm.zip`,
		`lx-source-android-arm64.zip`,
		`lx-source-linux-amd64v2.zip`,
		`lx-source-linux-amd64v3.zip`,
		`lx-source-linux-arm7.zip`,
		`lx-source-linux-arm64.zip`,
		`lx-source-windows-amd64v2.zip`,
		`lx-source-windows-amd64v2-go1.20.14.zip`,
		`lx-source-windows-amd64v3.zip`,
	} {
		eupd.WriteByte('+')
		eupd.WriteByte(' ')

		eupd.WriteByte('[')
		eupd.WriteString(v)
		eupd.WriteByte(']')
		eupd.WriteByte('(')
		eupd.WriteString(`https://r2eu.zxwy.link/gh/lx-source/`)
		eupd.WriteString(ever)
		eupd.WriteByte('/')
		eupd.WriteString(v)
		eupd.WriteByte(')')

		eupd.WriteByte('\n')
	}
	return eupd.String()
}

func main() {
	ever := version()
	fmt.Println(ever)

	eupd := changelog(ever)
	file, err := os.Create(`changelog.md`)
	if err != nil {
		panic(err)
	}
	file.WriteString(eupd)
	file.Close()
}
