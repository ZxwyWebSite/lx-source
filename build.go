//go:build ignore

// 一键编译脚本 `go run build.go`
package main

import (
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

	"github.com/ZxwyWebSite/ztool"
)

// 使用前请设置一下编译参数
var (
	// 系统-架构-C编译工具
	list_os_arch_cc = map[string]map[string]string{
		"linux": {
			"amd64": `x86_64-linux-gnu-gcc`,
			"arm":   `arm-linux-gnueabihf-gcc`,
		},
		"windows": {
			"amd64": `/usr/local/x86_64-w64-mingw32-cross/bin/x86_64-w64-mingw32-gcc`,
		},
	}
	// 架构-版本
	list_arch_ver = map[string][]string{
		"amd64": {`v2`}, //{`v1`, `v2`, `v3`},
		"arm":   {`7`},  //{`5`, `6`, `7`},
	}
)

const (
	// 运行参数
	args_name = `lx-source` // 程序名称
	args_path = `bin/`      // 输出目录
	args_zpak = true        // 打包文件
)

var workDir string

// 编译
func doCompile(v_os, v_arch, v_archv, v_cc string) error {
	//  构建               |    目标系统  |       目标架构  |           优化等级            |    不包含调试信息  |  使用外部链接器  |  输出详细操作 | 静态编译 |  JSON解释器
	// `go build -o bin/$1-$(go env GOOS)-$(go env GOARCH)$(go env GOAMD64)$(go env GOARM) -ldflags "-s -w -linkmode external -extldflags '-v -static'" -tags=jsoniter`
	fname := func() string {
		name := strings.Join([]string{args_name, v_os, v_arch}, `-`)
		var wexe string
		if v_os == `windows` {
			wexe = `.exe`
		}
		return ztool.Str_FastConcat(name, v_archv, wexe)
	}()
	pname := filepath.Clean(ztool.Str_FastConcat(args_path, fname))
	cmd := ztool.Str_FastConcat(
		`go build -o `, pname,
		` -gcflags=-trimpath="`, workDir, `" -asmflags=-trimpath="`, workDir, `" -trimpath -buildvcs=false`,
		` -ldflags "-s -w -linkmode external" -tags "go_json"`, // go_json | json(std) | jsoniter | sonic
	)
	// 输出要执行的命令
	ztool.Cmd_FastPrintln(ztool.Str_FastConcat(`执行命令：`, cmd))
	// 设置环境&执行编译
	envmap := map[string]string{
		`GOOS`:        v_os,
		`GOARCH`:      v_arch,
		`AR`:          `llvm-ar`, // 脚本默认使用Clang的Archiver, 没装llvm请注释掉以使用系统默认值
		`CC`:          v_cc,
		`CGO_ENABLED`: `1`,
		ztool.Str_FastConcat(`GO`, strings.ToUpper(v_arch)): v_archv, // GO{ARCH} Eg: GOARM, GOAMD64
	}
	setenv := func(env map[string]string) error {
		var handler ztool.Err_HandleList
		for k, v := range env {
			handler.Do(func() error {
				return os.Setenv(k, v)
			})
		}
		return handler.Err
	}
	if err := setenv(envmap); err != nil {
		return err
	}
	if err := ztool.Cmd_aSyncExec(cmd); err != nil {
		return err
	}
	// 打包文件
	if args_zpak { // DoSomeThing...
		if !ztool.Fbj_IsExists(`archieve`) {
			os.MkdirAll(filepath.Join(args_path, `archieve`), 0755)
		}
		if err := ztool.Pak_ZipFile(
			pname,
			filepath.Join(args_path, `archieve`, ztool.Str_LastBefore(fname, `.`))+`.zip`,
			ztool.Pak_ZipConfig{UnPath: true},
		); err != nil {
			ztool.Cmd_FastPrintln(ztool.Str_FastConcat(`打包["`, pname, `"]出错：`, err.Error()))
		} else {
			ztool.Cmd_FastPrintln(ztool.Str_FastConcat(`打包["`, pname, `"]完成`))
		}
	}
	return nil
}

func init() {
	if runtime.GOOS != `linux` {
		ztool.Cmd_FastPrintln("简易脚本，未对Linux以外系统做适配，请复制执行以下命令编译：\ngo build -ldflags \"-s -w\" -tags \"go_json\"\n如无报错则会在本目录生成名为lx-source的可执行文件。")
		os.Exit(1)
	}
	workDir, _ = os.Getwd()
	ztool.Cmd_FastPrintln(ztool.Str_FastConcat(`
 ================================
 |  Golang 一键编译脚本
 | 程序名称：`, args_name, `
 | 输出目录：`, args_path, `
 | 打包文件：`, strconv.FormatBool(args_zpak), `
 ================================
`))
}

func main() {
	var handler = ztool.Err_NewDefHandleList()
	handler.Do(func() error {
		// 检测入口函数是否存在
		if !ztool.Fbj_IsExists(`main.go`) {
			ztool.Cmd_FastPrintln(`入口函数不存在，请在源码根目录运行此程序！`)
			return ztool.Err_EsContinue
		}
		// 检测输出目录是否存在 (已在zTool中增加相关检测)
		// if !ztool.Fbj_IsExists(args_path) {
		// 	ztool.Cmd_FastPrintln(ztool.Str_FastConcat(`输出目录 "`, args_path, `" 不存在，尝试创建`))
		// 	return os.MkdirAll(args_path, 0755)
		// }
		return nil
	})
	for v_os, v_arch_cc := range list_os_arch_cc {
		for v_arch, v_cc := range v_arch_cc {
			// 检测CC是否存在
			o, e := ztool.Cmd_aWaitExec(ztool.Str_FastConcat(`which `, v_cc))
			if !ztool.Fbj_IsExists(v_cc) && (e != nil || o == ``) {
				ztool.Cmd_FastPrintln(ztool.Str_FastConcat(`编译工具 ["`, v_cc, `"] 不存在，跳过 `, v_arch, ` 架构`))
				continue
			}
			// 继续编译
			for _, v_arch_ver := range list_arch_ver[v_arch] {
				// handler.Do(func() error { return tool.ErrContinue })
				handler.Do(func() error {
					// (测试) 快速输出编译参数
					ztool.Cmd_FastPrintln(ztool.Str_FastConcat(`开始编译：`, v_os, `/`, v_arch, `/`, v_arch_ver, `/`, `[`, v_cc, `], 任务编号 `, handler.NumStr()))
					// 编译对应文件
					// return nil
					return doCompile(v_os, v_arch, v_arch_ver, v_cc)
				}) // handler.Do(func() error { return doCompile(v_os, v_arch, v_arch_ver, v_cc) })
			}
		}
	}
	if res := handler.Result(); res != nil {
		ztool.Cmd_FastPrintln(ztool.Str_FastConcat(`发生错误：`, res.Errors()))
		return
	}
	ztool.Cmd_FastPrintln(`恭喜！所有任务成功完成`)
}
