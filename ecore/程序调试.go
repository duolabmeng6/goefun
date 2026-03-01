package ecore

import (
    "fmt"
    "github.com/kr/pretty"
)

// E调试输出 格式化输出调试信息到控制台
//
// 参数：
//   a ...interface{} - 要输出的内容
func E调试输出(a ...interface{}) {
    pretty.Println(a...)
}

// E调试输出格式化 使用格式化字符串输出调试信息到控制台
//
// 参数：
//   s string - 格式化字符串
//   a ...interface{} - 格式化参数
func E调试输出格式化(s string, a ...interface{}) {
    fmt.Printf(s, a...)
}
