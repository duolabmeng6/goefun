// Copyright 2023 The duolabmeng6 Authors. All rights reserved.
// license that can be found in the LICENSE file.

// Package ecore 提供程序调试功能
package ecore

import (
	"fmt"

	"github.com/kr/pretty"
)

// E调试输出 调试输出函数，以美观格式打印变量。
//
// 参数：
//   - a: 可变参数，要输出的内容
//
// 示例：
//
//	E调试输出("用户信息", 用户对象)
//	E调试输出("列表内容", []string{"a", "b", "c"})
func E调试输出(a ...interface{}) {
	pretty.Println(a...)
}

// E调试输出格式化 格式化输出调试信息。
//
// 参数：
//   - s: 格式化字符串
//   - a: 可变参数，格式化参数
//
// 示例：
//
//	E调试输出格式化("用户名: %s, 年龄: %d", "张三", 25)
func E调试输出格式化(s string, a ...interface{}) {
	fmt.Printf(s, a...)
}
