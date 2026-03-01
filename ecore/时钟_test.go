// Copyright 2023 The duolabmeng6 Authors. All rights reserved.
// license that can be found in the LICENSE file.

package ecore

import (
	"testing"
)

// TestNew时钟 测试时钟功能（默认禁用，需要手动启用）
func TestNew时钟(t *testing.T) {
	return // 禁用测试，取消此行可启用
	时钟 := New时钟()
	时间统计 := New时间统计类()
	i := 0
	时钟.E创建(func() bool {
		i++
		E延时(1000)
		if i > 10 {
			E调试输出("10次停止了")
			return false
		}
		E调试输出(i, 时间统计.E取毫秒())
		return true
	}, 100)

	时钟.E创建执行一次(func() {
		E调试输出("执行一次", 时间统计.E取毫秒())
	}, 100)

	时钟周期函数 := 时钟.E时钟周期函数(func() {
		E调试输出("3秒调用多少次都只执行1次", 时间统计.E取毫秒())
	}, 1000*3)

	时钟.E创建(func() bool {
		时钟周期函数()
		return true
	}, 100)

	select {}
}
