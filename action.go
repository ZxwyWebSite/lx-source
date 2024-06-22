//go:build ignore

package main

import (
	"archive/zip"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

const (
	// 运行参数
	args_name = `lx-source` // 程序名称
	args_path = `dist/`     // 输出目录
	args_zpak = true        // 打包文件
	// args_repo = `repo/`     // 源码目录
	args_home = `/home/runner` // 用户目录
)

var (
	workDir string
	homeDir string
)

func init() {
	if runtime.GOOS != `linux` {
		fmt.Println(`不兼容的运行环境:`, runtime.GOOS)
		os.Exit(0)
	}
	workDir, _ = os.Getwd()
	fmt.Println(`运行目录:`, workDir)
	homeDir = os.Getenv(`HOME`)
	if homeDir == `` {
		homeDir = args_home
	}
	fmt.Println(`用户目录:`, homeDir)
}

type (
	// 架构参数 [v2]
	list_vers map[string]struct {
		Tags string
	}
	// CGO参数 [nil]
	list_cgos struct {
		AR  string
		CC  string
		CXX string
	}
	// 架构列表 [amd64]
	list_arch map[string]struct {
		Cgos *list_cgos
		Vers list_vers
		Venv string // 覆盖架构参数名 'mipsle'->'GOMIPS'
	}
	// 目标系统 [linux]
	list_goos map[string]struct {
		Arch list_arch
	}
	// 编译环境 [go1.20.14]
	list_conf map[string]struct {
		Args []string
		GoOS list_goos
	}
)

// 构建参数
var def_args = []string{
	`-trimpath`, `-buildvcs=false`,
	`-ldflags`, `-s -w -linkmode external`,
}

type param struct {
	GoVer  string   // 环境 go1.20.14
	GoOS   string   // 系统 linux
	GoArch string   // 架构 amd64
	GoIns  string   // 指令 GOAMD64=v2
	Args   []string // 参数 ldflags
	Tag    string   // 标志 go_json
	Cgos   *list_cgos
	Venv   string
}

// 获取相对用户目录
func home(str string) string {
	return homeDir + `/` + str
}

// 检测环境是否存在
func chkenv(s ...string) (err error) {
	for _, f := range s {
		if _, e := exec.LookPath(f); e != nil && !errors.Is(e, exec.ErrDot) {
			err = fmt.Errorf(`未找到指定环境: %s`, e)
			break
		}
	}
	return
}

func main() {
	var def_list = list_conf{
		`go`: {
			Args: def_args,
			GoOS: list_goos{
				`linux`: {
					Arch: list_arch{
						`amd64`: {
							Cgos: &list_cgos{
								AR:  `x86_64-linux-gnu-ar`,
								CC:  `x86_64-linux-gnu-gcc`,
								CXX: `x86_64-linux-gnu-g++`,
							},
							Vers: list_vers{
								`v1`: {
									Tags: `go_json`,
								},
								`v2`: {
									Tags: `go_json`,
								},
								`v3`: {
									Tags: `sonic avx`,
								},
								`v4`: {
									Tags: `sonic avx`,
								},
							},
						},
						`arm`: {
							Cgos: &list_cgos{
								AR:  `arm-linux-gnueabihf-gcc-ar`,
								CC:  `arm-linux-gnueabihf-gcc`,
								CXX: `arm-linux-gnueabihf-cpp`,
							},
							Vers: list_vers{
								`5`: {
									Tags: `go_json`,
								},
								`6`: {
									Tags: `go_json`,
								},
								`7`: {
									Tags: `go_json`,
								},
							},
						},
						`arm64`: {
							Cgos: &list_cgos{
								AR:  `aarch64-linux-gnu-gcc-ar`,
								CC:  `aarch64-linux-gnu-gcc`,
								CXX: `aarch64-linux-gnu-cpp`,
							},
							Vers: list_vers{
								``: {
									Tags: `go_json`,
								},
							},
						},
					},
				},
				`windows`: {
					Arch: list_arch{
						`amd64`: {
							Cgos: &list_cgos{
								AR:  `x86_64-w64-mingw32-ar`,
								CC:  `x86_64-w64-mingw32-gcc`,
								CXX: `x86_64-w64-mingw32-cpp`,
							},
							Vers: list_vers{
								`v2`: {
									Tags: `go_json`,
								},
								`v3`: {
									Tags: `sonic avx`,
								},
								`v4`: {
									Tags: `sonic avx`,
								},
							},
						},
					},
				},
				`android`: {
					Arch: list_arch{
						`amd64`: {
							Cgos: &list_cgos{
								AR:  home(`android-ndk-r26b/toolchains/llvm/prebuilt/linux-x86_64/bin/llvm-ar`),
								CC:  home(`android-ndk-r26b/toolchains/llvm/prebuilt/linux-x86_64/bin/x86_64-linux-android24-clang`),
								CXX: home(`android-ndk-r26b/toolchains/llvm/prebuilt/linux-x86_64/bin/x86_64-linux-android24-clang++`),
							},
							Vers: list_vers{
								``: {
									Tags: `go_json`,
								},
							},
						},
						`arm64`: {
							Cgos: &list_cgos{
								AR:  home(`android-ndk-r26b/toolchains/llvm/prebuilt/linux-x86_64/bin/llvm-ar`),
								CC:  home(`android-ndk-r26b/toolchains/llvm/prebuilt/linux-x86_64/bin/aarch64-linux-android24-clang`),
								CXX: home(`android-ndk-r26b/toolchains/llvm/prebuilt/linux-x86_64/bin/aarch64-linux-android24-clang++`),
							},
							Vers: list_vers{
								``: {
									Tags: `go_json`,
								},
							},
						},
						`386`: {
							Cgos: &list_cgos{
								AR:  home(`android-ndk-r26b/toolchains/llvm/prebuilt/linux-x86_64/bin/llvm-ar`),
								CC:  home(`android-ndk-r26b/toolchains/llvm/prebuilt/linux-x86_64/bin/i686-linux-android24-clang`),
								CXX: home(`android-ndk-r26b/toolchains/llvm/prebuilt/linux-x86_64/bin/i686-linux-android24-clang++`),
							},
							Vers: list_vers{
								``: {
									Tags: `go_json`,
								},
							},
						},
						`arm`: {
							Cgos: &list_cgos{
								AR:  home(`android-ndk-r26b/toolchains/llvm/prebuilt/linux-x86_64/bin/llvm-ar`),
								CC:  home(`android-ndk-r26b/toolchains/llvm/prebuilt/linux-x86_64/bin/armv7a-linux-androideabi24-clang`),
								CXX: home(`android-ndk-r26b/toolchains/llvm/prebuilt/linux-x86_64/bin/armv7a-linux-androideabi24-clang++`),
							},
							Vers: list_vers{
								``: {
									Tags: `go_json`,
								},
							},
						},
					},
				},
			},
		},
		home(`go/bin/go1.20.14`): {
			Args: []string{
				`-trimpath`, `-buildvcs=false`,
				`-ldflags`, `-s -w -extldflags '-v -static'`,
			},
			GoOS: list_goos{
				`windows`: {
					Arch: list_arch{
						`amd64`: {
							Cgos: &list_cgos{
								AR:  `x86_64-w64-mingw32-ar`,
								CC:  `x86_64-w64-mingw32-gcc`,
								CXX: `x86_64-w64-mingw32-cpp`,
							},
							Vers: list_vers{
								`v1`: {
									Tags: `go_json`,
								},
								`v2`: {
									Tags: `go_json`,
								},
								`v3`: {
									Tags: `sonic avx`,
								},
							},
						},
					},
				},
				`linux`: {
					Arch: list_arch{
						`arm`: {
							Cgos: &list_cgos{
								AR:  `arm-linux-gnueabihf-gcc-ar`,
								CC:  `arm-linux-gnueabihf-gcc`,
								CXX: `arm-linux-gnueabihf-cpp`,
							},
							Vers: list_vers{
								`5`: {
									Tags: `go_json`,
								},
								`6`: {
									Tags: `go_json`,
								},
								`7`: {
									Tags: `go_json`,
								},
							},
						},
						`arm64`: {
							Cgos: &list_cgos{
								AR:  `aarch64-linux-gnu-gcc-ar`,
								CC:  `aarch64-linux-gnu-gcc`,
								CXX: `aarch64-linux-gnu-cpp`,
							},
							Vers: list_vers{
								``: {
									Tags: `go_json`,
								},
							},
						},
						// 针对部分OpenWrt路由器系统 暂不支持开启CGO
						`mips`: {
							Vers: list_vers{
								`hardfloat`: {
									Tags: `go_json`,
								},
								`softfloat`: {
									Tags: `go_json`,
								},
							},
						},
						`mipsle`: {
							Vers: list_vers{
								`hardfloat`: {
									Tags: `go_json`,
								},
								`softfloat`: {
									Tags: `go_json`,
								},
							},
							Venv: `MIPS`,
						},
						`mips64`: {
							Vers: list_vers{
								`hardfloat`: {
									Tags: `go_json`,
								},
								`softfloat`: {
									Tags: `go_json`,
								},
							},
						},
						`mips64le`: {
							Vers: list_vers{
								`hardfloat`: {
									Tags: `go_json`,
								},
								`softfloat`: {
									Tags: `go_json`,
								},
							},
							Venv: `MIPS64`,
						},
					},
				},
				// Mac OS
				`darwin`: {
					Arch: list_arch{
						`amd64`: {
							Vers: list_vers{
								`v2`: {
									Tags: `go_json`,
								},
								`v3`: {
									Tags: `sonic avx`,
								},
							},
						},
						`arm64`: {
							Vers: list_vers{
								``: {
									Tags: `go_json`,
								},
							},
						},
					},
				},
			},
		},
	}
	fmt.Printf(`
 ================================
 |  Action 一键编译脚本
 | 程序名称：%v
 | 输出目录：%v
 | 打包文件：%v
 ================================
`, args_name, args_path, args_zpak)
	// 解析配置文件
	for goVer, conf_list := range def_list {
		// 环境检测
		if err := chkenv(goVer); err != nil {
			fmt.Println(err, `跳过该环境`)
			continue
		}
		for goOS, goos_list := range conf_list.GoOS {
			for goArch, arch_list := range goos_list.Arch {
				// 工具链检测
				if arch_list.Cgos != nil {
					if err := chkenv(
						arch_list.Cgos.AR,
						arch_list.Cgos.CC,
						arch_list.Cgos.CXX,
					); err != nil {
						fmt.Println(err, `跳过该架构`)
						continue
					}
				}
				for goIns, vers_list := range arch_list.Vers {
					// 构建程序二进制
					if err := build(&param{
						GoVer:  goVer,
						GoOS:   goOS,
						GoArch: goArch,
						GoIns:  goIns,
						Args:   conf_list.Args,
						Tag:    vers_list.Tags,
						Cgos:   arch_list.Cgos,
						Venv:   arch_list.Venv,
					}); err != nil {
						fmt.Println(`err:`, err)
					}
				}
			}
		}
	}
	fmt.Println(`执行结束`)
}

func build(p *param) (err error) {
	// 拼接程序名称
	var b strings.Builder
	b.WriteString(args_name) // lx-source
	b.WriteByte('-')         // lx-source-
	b.WriteString(p.GoOS)    // lx-source-linux
	b.WriteByte('-')         // lx-source-linux-
	b.WriteString(p.GoArch)  // lx-source-linux-amd64
	/*var digit byte
	if p.Venv != `` {
		digit = p.Venv[len(p.Venv)-1]
	} else {
		digit = p.GoArch[len(p.GoArch)-1]
	}
	if !unicode.IsDigit(rune(digit)) {
		// 架构名结尾不是数字的再加一个连字符
		b.WriteByte('-') // lx-source-linux-mipsle-softfloat
	}*/
	b.WriteString(p.GoIns) // lx-source-linux-amd64v2
	if biname := filepath.Base(p.GoVer); biname != `go` {
		b.WriteByte('-')      // lx-source-linux-amd64v2-
		b.WriteString(biname) // lx-source-linux-amd64v2-go1.20.14
	}
	// 拼接输出名称
	oname := args_path + b.String() // dist/lx-source-linux-amd64v2
	if p.GoOS == `windows` {
		oname += `.exe` // dist/lx-source-linux-amd64v2.exe
	}
	fmt.Println(`开始编译:`, oname)
	fmt.Printf("编译参数: %+v\n", *p)
	// 填入参数并构建
	var args = []string{
		`build`, `-o`, oname,
		// `-asmflags=-trimpath="` + workDir + `"`,
		// `-gcflags=-trimpath="` + workDir + `"`,
		`-tags`, p.Tag,
	}
	cmd := exec.Command(
		p.GoVer,
		// append(append(args, p.Args...), args_repo)...,
		append(args, p.Args...)...,
	)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = cmd.Stdout
	cmd.Dir = workDir
	cmd.Env = append(os.Environ(), []string{
		`GOOS=` + p.GoOS,
		`GOARCH=` + p.GoArch,
	}...)
	if p.Cgos != nil {
		cmd.Env = append(cmd.Env, []string{
			`AR=` + p.Cgos.AR,
			`CC=` + p.Cgos.CC,
			`CXX=` + p.Cgos.CXX,
			`CGO_ENABLED=1`,
		}...)
	} /*else {
		cmd.Env = append(cmd.Env, `CGO_ENABLED=0`)
	}*/
	if p.GoIns != `` {
		if p.Venv != `` {
			cmd.Env = append(cmd.Env, `GO`+p.Venv+`=`+p.GoIns)
		} else {
			cmd.Env = append(cmd.Env, `GO`+strings.ToUpper(p.GoArch)+`=`+p.GoIns)
		}
	}

	if err = cmd.Start(); err == nil {
		err = cmd.Wait()
	}
	if err != nil || !args_zpak {
		return
	}
	// 打包输出文件
	/*apath := filepath.Join(args_path, `archieve`)
	if _, e := os.Stat(apath); e != nil {
		if os.IsNotExist(e) {
			err = os.MkdirAll(apath, os.ModePerm)
			if err != nil {
				return
			}
		}
	}*/
	zipname := filepath.Join(args_path, b.String()+`.zip`)
	fmt.Println(`打包文件:`, zipname)
	zipfile, err := os.Create(zipname)
	if err != nil {
		return err
	}
	archive := zip.NewWriter(zipfile)
	info, err := os.Lstat(oname)
	if err == nil {
		header, _ := zip.FileInfoHeader(info)
		header.Method = zip.Deflate
		header.Name = filepath.Base(oname)
		writer, err := archive.CreateHeader(header)
		if err == nil {
			file, err := os.Open(oname)
			if err == nil {
				_, err = io.Copy(writer, file)
				file.Close()
				if err == nil {
					err = os.Remove(oname)
				}
			}
		}
	}
	archive.Close()
	zipfile.Close()
	return err
}
