// Copyright 2023 The duolabmeng6 Authors. All rights reserved.
// license that can be found in the LICENSE file.

// Package ecore 提供流程控制功能
package ecore

import "os"

// E结束 本命令结束当前易程序的运行。
//
// 参数：
//   - 无
//
// 示例：
//
//	E结束() // 立即退出程序
func E结束() {
	os.Exit(0)
}
