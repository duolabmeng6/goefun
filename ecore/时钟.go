// Copyright 2023 The duolabmeng6 Authors. All rights reserved.
// license that can be found in the LICENSE file.

// Package ecore 提供定时任务功能
//
// 提供周期性任务执行功能。使用方式类似于Linux下的Crontab或者Java里的Quartz。本对象线程安全。
package ecore

import (
	"time"

	"github.com/duolabmeng6/goefun/src/cable"
)

// E时钟 时钟结构体，提供定时任务相关功能
type E时钟 struct {
}

// New时钟 创建一个新的时钟对象。
//
// 返回值：
//   - *E时钟: 新创建的时钟对象
func New时钟() *E时钟 {
	this := new(E时钟)
	return this
}

// E创建 创建一个周期性执行的定时任务。
//
// 参数：
//   - fn: 回调函数，返回bool值控制是否继续执行
//   - 时钟周期: 执行周期（毫秒）
//
// 示例：
//
//	时钟.E创建(func() bool {
//	    fmt.Println("每100毫秒执行一次")
//	    return true // 返回false停止执行
//	}, 100)
func (this *E时钟) E创建(fn func() bool, 时钟周期 int64) {
	cable.SetInterval(fn, time.Duration(时钟周期)*time.Millisecond)
}

// E创建执行一次 创建一个只执行一次的定时任务。
//
// 参数：
//   - fn: 回调函数
//   - 时钟周期: 延迟时间（毫秒）
//
// 示例：
//
//	时钟.E创建执行一次(func() {
//	    fmt.Println("100毫秒后执行一次")
//	}, 100)
func (this *E时钟) E创建执行一次(fn func(), 时钟周期 int64) {
	cable.SetTimeout(fn, time.Duration(时钟周期)*time.Millisecond)
}

// E时钟周期函数 返回一个函数，无论调用多少次，它只会在指定的间隔内执行一次。
//
// 参数：
//   - fn: 回调函数
//   - 时钟周期: 执行间隔（毫秒）
//
// 返回值：
//   - func(): 返回一个节流函数
//
// 示例：
//
//	节流函数 := 时钟.E时钟周期函数(func() {
//	    fmt.Println("3秒内只执行一次")
//	}, 3000)
//	节流函数() // 多次调用也只在3秒内执行一次
func (this *E时钟) E时钟周期函数(fn func(), 时钟周期 int64) func() {
	return cable.Throttle(fn, time.Duration(时钟周期)*time.Millisecond, cable.ThrottleOptions{})
}

// E时钟_创建 创建一个周期性执行的定时任务（函数版本）。
//
// 参数：
//   - fn: 回调函数，返回bool值控制是否继续执行
//   - 时钟周期: 执行周期（毫秒）
func E时钟_创建(fn func() bool, 时钟周期 int64) {
	cable.SetInterval(fn, time.Duration(时钟周期)*time.Millisecond)
}

// E时钟_创建执行一次 创建一个只执行一次的定时任务（函数版本）。
//
// 参数：
//   - fn: 回调函数
//   - 时钟周期: 延迟时间（毫秒）
func E时钟_创建执行一次(fn func(), 时钟周期 int64) {
	cable.SetTimeout(fn, time.Duration(时钟周期)*time.Millisecond)
}

// E时钟_创建周期函数 返回一个节流函数（函数版本）。
//
// 参数：
//   - fn: 回调函数
//   - 时钟周期: 执行间隔（毫秒）
//
// 返回值：
//   - func(): 返回一个节流函数
func E时钟_创建周期函数(fn func(), 时钟周期 int64) func() {
	return cable.Throttle(fn, time.Duration(时钟周期)*time.Millisecond, cable.ThrottleOptions{})
}
