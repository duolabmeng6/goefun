// Copyright 2023 The duolabmeng6 Authors. All rights reserved.
// license that can be found in the LICENSE file.

// Package ecore 提供网络工具函数
package ecore

import (
	"fmt"
	"math/rand"
	"time"
)

// E取随机ip 生成一个随机IP地址。
//
// 返回值：
//   - string: 随机生成的IP地址
//
// 示例：
//
//	ip := E取随机ip() // 返回类似 "192.168.1.100" 的IP
func E取随机ip() string {
	rand.Seed(time.Now().Unix())
	ip := fmt.Sprintf("%d.%d.%d.%d", E取随机数(50, 254), E取随机数(50, 254), E取随机数(50, 254), E取随机数(50, 254))
	return ip
}
