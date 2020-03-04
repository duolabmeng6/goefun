package core

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func E取命令行() []string {
	return os.Args
}
//调用格式： 〈文本型〉 取运行目录 （） - 系统核心支持库->环境存取
//英文名称：GetRunPath
//取当前被执行的易程序文件所处的目录。本命令为初级命令。
//
//操作系统需求： Windows
func E取运行目录() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Println(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

//调用格式： 〈文本型〉 读环境变量 （文本型 环境变量名称） - 系统核心支持库->环境存取
//英文名称：GetEnv
//返回文本，它关连于一个操作系统环境变量。成功时返回所取得的值，失败则返回空文本。本命令为初级命令。
//参数<1>的名称为“环境变量名称”，类型为“文本型（text）”。
//
//操作系统需求： Windows、Linux
func E读环境变量(环境变量名称 string) string {
	return os.Getenv(环境变量名称)
}

//调用格式： 〈逻辑型〉 写环境变量 （文本型 环境变量名称，文本型 欲写入内容） - 系统核心支持库->环境存取
//英文名称：PutEnv
//修改或建立指定的操作系统环境变量。成功返回真，失败返回假。本命令为初级命令。
//参数<1>的名称为“环境变量名称”，类型为“文本型（text）”。
//参数<2>的名称为“欲写入内容”，类型为“文本型（text）”。
//
//操作系统需求： Windows、Linux

func E写环境变量(环境变量名称 string, 欲写入内容 string) bool {
	err := os.Setenv(环境变量名称, 欲写入内容)
	return err == nil
}
