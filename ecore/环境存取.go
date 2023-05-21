package ecore

import (
	"flag"
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"path/filepath"
	"strings"
)

// E取命令行  获取命令行参数，以字符串切片形式返回
//
// 返回值为字符串切片，包含命令行中的所有参数
func E取命令行() []string {
	return os.Args
}

// E取运行目录
//
// 取当前被执行的易程序文件所处的目录。
func E取运行目录() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Println(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

// E读环境变量
//
// 返回文本，它关连于一个操作系统环境变量。成功时返回所取得的值，失败则返回空文本。
//
// 参数<1>的名称为“环境变量名称”，类型为“文本型（text）”。
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

// E写环境变量
// 修改或建立指定的操作系统环境变量。成功返回真，失败返回假。
//
// 参数<1>的名称为“环境变量名称”，类型为“文本型（text）”。
//
// 参数<2>的名称为“欲写入内容”，类型为“文本型（text）”。
func E写环境变量(环境变量名称 string, 欲写入内容 string) bool {
	err := os.Setenv(环境变量名称, 欲写入内容)
	return err == nil
}

// env
func E加载环境变量_从文件(filepath string) bool {
	err := godotenv.Load(filepath)
	if err != nil {
		E调试输出("加载环境变量文件失败文件路径为:"+filepath, "错误信息为", err)
		return false
	}
	return true
}

func E加载环境变量_从内容到map(s string) (map[string]string, error) {
	myEnv, err := godotenv.Unmarshal(s)
	return myEnv, err
}

// 带有默认值的env读取函数
// E读环境变量 一样效果
func Env(环境变量名称 string, 默认值 string) string {
	v := E读环境变量(环境变量名称, 默认值)
	return v
}

func E设置命令行(name string, defaultvalue string, help string, value *string) {
	flag.StringVar(value, name, defaultvalue, help)
}

func E命令行解析() {
	flag.Parse()
}
