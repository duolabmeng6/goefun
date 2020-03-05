package core

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

//调用格式： 〈无返回值〉 取命令行 （文本型变量数组 存放被取回命令行文本的数组变量） - 系统核心支持库->环境存取
//英文名称：GetCmdLine
//本命令可以取出在启动易程序时附加在其可执行文件名后面的所有以空格分隔的命令行文本段。本命令为初级命令。
//参数<1>的名称为“存放被取回命令行文本的数组变量”，类型为“文本型（text）”，提供参数数据时只能提供变量数组。在命令执行完毕后，本变量数组内被顺序填入在启动易程序时附加在其可执行文件名后面的以空格分隔的命令行文本段。变量数组内原有数据被全部销毁，变量数组的维数被自动调整为命令行文本段数。
//
//操作系统需求： Windows、Linux
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
func E读环境变量(环境变量名称 string, 默认值 ...interface{}) string {
	var def string
	if len(默认值) > 1 {
		def = E到文本(默认值[0])
	}
	e := os.Getenv(环境变量名称)
	if e == "" {
		return def
	} else {
		return e
	}
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
