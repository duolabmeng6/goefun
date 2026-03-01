// Copyright 2023 The duolabmeng6 Authors. All rights reserved.
// license that can be found in the LICENSE file.

// Package ecore 提供时间统计功能
package ecore

import (
	"fmt"
	"time"
)

// E时间统计类 时间统计结构体，用于统计代码执行时间
type E时间统计类 struct {
	t time.Time
}

// New时间统计类 创建一个新的时间统计对象并自动开始计时。
//
// 返回值：
//   - *E时间统计类: 新创建的时间统计对象
//
// 示例：
//
//	计时器 := New时间统计类()
//	// ... 执行代码 ...
//	fmt.Println(计时器.E取毫秒()) // 输出执行时间（毫秒）
func New时间统计类() *E时间统计类 {
	t := new(E时间统计类)
	t.E开始()
	return t
}

// E开始 开始或重新开始计时。
func (this *E时间统计类) E开始() {
	this.t = time.Now()
}

// E取毫秒 获取从开始到现在经过的毫秒数。
//
// 返回值：
//   - string: 经过的毫秒数（字符串形式）
func (this *E时间统计类) E取毫秒() string {
	return fmt.Sprintf("%d", time.Since(this.t).Milliseconds())
}

// E取秒 获取从开始到现在经过的秒数。
//
// 返回值：
//   - string: 经过的秒数（字符串形式，保留3位小数）
func (this *E时间统计类) E取秒() string {
	return fmt.Sprintf("%.3f", float64(time.Since(this.t).Milliseconds())/float64(1000))
}
