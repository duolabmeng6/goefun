// Copyright 2023 The duolabmeng6 Authors. All rights reserved.
// license that can be found in the LICENSE file.

// Package ecore 提供控制台操作相关功能
package ecore

import (
    "bufio"
    "fmt"
    "os"
)

// E标准输出 向控制台写入数据并自动换行。
//
// 参数：
//   - 欲写入的数据: 可变参数，欲写入控制台的数据
//
// 示例：
//
//    E标准输出("Hello", "World") // 输出: [Hello World]
func E标准输出(欲写入的数据 ...interface{}) {
    fmt.Print(欲写入的数据, "\n")
}

// E标准输入 从控制台读取一行数据。
//
// 返回值：
//   - string: 读取到的文本数据
//
// 示例：
//
//    文本 := E标准输入() // 等待用户输入
func E标准输入() string {
    input := bufio.NewScanner(os.Stdin)
    input.Scan()
    return input.Text()
}
