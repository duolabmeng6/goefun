// Copyright 2023 The duolabmeng6 Authors. All rights reserved.
// license that can be found in the LICENSE file.

// Package ecore 提供数值转换相关功能
package ecore

import (
	"strconv"
)

// E取十六进制文本 返回一个字符串，代表指定数值的十六进制形式。
//
// 参数：
//   - n: 整数型，欲取进制文本的数值
//
// 返回值：
//   - string: 整数n的十六进制形式
//
// 示例：
//
//	E取十六进制文本(255) // 返回 "ff"
func E取十六进制文本(n int) string {
	return strconv.FormatInt(int64(n), 16)
}

// E取八进制文本 返回一个字符串，代表指定数值的八进制形式。
//
// 参数：
//   - n: 整数型，欲取进制文本的数值
//
// 返回值：
//   - string: 整数n的八进制形式
//
// 示例：
//
//	E取八进制文本(64) // 返回 "100"
func E取八进制文本(n int) string {
	return strconv.FormatInt(int64(n), 8)
}

// E十六进制 转整数，返回一个整数，代表指定十六进制文本常量对应的整数值。
//
// 参数：
//   - s: 文本型，十六进制文本常量（如 "0x1234" 或 "1234"）
//
// 返回值：
//   - int: 十六进制文本常量对应的整数值
//
// 示例：
//
//	E十六进制("0xff") // 返回 255
//	E十六进制("ff")   // 返回 0 (需要0x前缀)
func E十六进制(s string) int {
	n, _ := strconv.ParseInt(s, 0, 0)
	return int(n)
}

// E二进制 转整数，返回一个整数，代表指定二进制文本常量对应的整数值。
//
// 参数：
//   - s: 文本型，二进制文本常量（如 "1010"）
//
// 返回值：
//   - int: 二进制文本常量对应的整数值
//
// 示例：
//
//	E二进制("1010") // 返回 10
func E二进制(s string) int {
	n, _ := strconv.ParseInt(s, 2, 0)
	return int(n)
}
