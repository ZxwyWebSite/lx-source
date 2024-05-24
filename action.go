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

var workDir string

func init() {
	if runtime.GOOS != `linux` {
		fmt.Println(`不兼容的运行环境:`, runtime.GOOS)
		os.Exit(0)
	}
	workDir, _ = os.Getwd()
	fmt.Println(`运行目录:`, workDir)
}

type (
	list_vers map[string]struct {
		Tags string
	}
	list_arch map[string]struct {
		AR   string
		CC   string
		CXX  string
		Vers list_vers
	}
	list_goos map[string]struct {
		Arch list_arch
	}
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
	AR     string
	CC     string
	CXX    string
}

func main() {
	var def_list = list_conf{
		`go`: {
			Args: def_args,
			GoOS: list_goos{
				`linux`: {
					Arch: list_arch{
						`amd64`: {
							AR:  `x86_64-linux-gnu-ar`,
							CC:  `x86_64-linux-gnu-gcc`,
							CXX: `x86_64-linux-gnu-g++`,
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
							AR:  `arm-linux-gnueabihf-gcc-ar`,
							CC:  `arm-linux-gnueabihf-gcc`,
							CXX: `arm-linux-gnueabihf-cpp`,
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
							AR:  `aarch64-linux-gnu-gcc-ar`,
							CC:  `aarch64-linux-gnu-gcc`,
							CXX: `aarch64-linux-gnu-cpp`,
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
							AR:  `x86_64-w64-mingw32-ar`,
							CC:  `x86_64-w64-mingw32-gcc`,
							CXX: `x86_64-w64-mingw32-cpp`,
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
							AR:  args_home + `/android-ndk-r26b/toolchains/llvm/prebuilt/linux-x86_64/bin/llvm-ar`,
							CC:  args_home + `/android-ndk-r26b/toolchains/llvm/prebuilt/linux-x86_64/bin/x86_64-linux-android24-clang`,
							CXX: args_home + `/android-ndk-r26b/toolchains/llvm/prebuilt/linux-x86_64/bin/x86_64-linux-android24-clang++`,
							Vers: list_vers{
								``: {
									Tags: `go_json`,
								},
							},
						},
						`arm64`: {
							AR:  args_home + `/android-ndk-r26b/toolchains/llvm/prebuilt/linux-x86_64/bin/llvm-ar`,
							CC:  args_home + `/android-ndk-r26b/toolchains/llvm/prebuilt/linux-x86_64/bin/aarch64-linux-android24-clang`,
							CXX: args_home + `/android-ndk-r26b/toolchains/llvm/prebuilt/linux-x86_64/bin/aarch64-linux-android24-clang++`,
							Vers: list_vers{
								``: {
									Tags: `go_json`,
								},
							},
						},
						`386`: {
							AR:  args_home + `/android-ndk-r26b/toolchains/llvm/prebuilt/linux-x86_64/bin/llvm-ar`,
							CC:  args_home + `/android-ndk-r26b/toolchains/llvm/prebuilt/linux-x86_64/bin/i686-linux-android24-clang`,
							CXX: args_home + `/android-ndk-r26b/toolchains/llvm/prebuilt/linux-x86_64/bin/i686-linux-android24-clang++`,
							Vers: list_vers{
								``: {
									Tags: `go_json`,
								},
							},
						},
						`arm`: {
							AR:  args_home + `/android-ndk-r26b/toolchains/llvm/prebuilt/linux-x86_64/bin/llvm-ar`,
							CC:  args_home + `/android-ndk-r26b/toolchains/llvm/prebuilt/linux-x86_64/bin/armv7a-linux-androideabi24-clang`,
							CXX: args_home + `/android-ndk-r26b/toolchains/llvm/prebuilt/linux-x86_64/bin/armv7a-linux-androideabi24-clang++`,
							Vers: list_vers{
								``: {
									Tags: `go_json`,
								},
							},
						},
					},
				},
				// `darwin`: {
				// 	Arch: list_arch{
				// 		`amd64`: {
				// 			CC: ``,
				// 		},
				// 		`arm64`: {
				// 			CC: ``,
				// 		},
				// 	},
				// },
			},
		},
		args_home + `/go/bin/go1.20.14`: {
			Args: def_args,
			GoOS: list_goos{
				`windows`: {
					Arch: list_arch{
						`amd64`: {
							AR:  `x86_64-w64-mingw32-ar`,
							CC:  `x86_64-w64-mingw32-gcc`,
							CXX: `x86_64-w64-mingw32-cpp`,
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
		for goOS, goos_list := range conf_list.GoOS {
			for goArch, arch_list := range goos_list.Arch {
				for goIns, vers_list := range arch_list.Vers {
					// 构建程序二进制
					if err := build(&param{
						GoVer:  goVer,
						GoOS:   goOS,
						GoArch: goArch,
						GoIns:  goIns,
						Args:   conf_list.Args,
						Tag:    vers_list.Tags,
						AR:     arch_list.AR,
						CC:     arch_list.CC,
						CXX:    arch_list.CXX,
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
	// 检测必要环境
	for _, f := range []string{
		p.GoVer, p.AR, p.CC, p.CXX,
	} {
		if _, e := exec.LookPath(f); e != nil && !errors.Is(e, exec.ErrDot) {
			err = fmt.Errorf(`未找到指定环境: %s`, e)
			return
		}
	}
	// 拼接程序名称
	var b strings.Builder
	b.WriteString(args_name) // lx-source
	b.WriteByte('-')         // lx-source-
	b.WriteString(p.GoOS)    // lx-source-linux
	b.WriteByte('-')         // lx-source-linux-
	b.WriteString(p.GoArch)  // lx-source-linux-amd64
	b.WriteString(p.GoIns)   // lx-source-linux-amd64v2
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
		`AR=` + p.AR,
		`CC=` + p.CC,
		`CXX=` + p.CXX,
		`CGO_ENABLED=1`,
		`GO` + strings.ToUpper(p.GoArch) + `=` + p.GoIns,
	}...)

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
