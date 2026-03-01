// Copyright 2023 The duolabmeng6 Authors. All rights reserved.
// license that can be found in the LICENSE file.

// Package ecore 提供类型转换功能
package ecore

import (
	"github.com/gogf/gf/v2/util/gconv"
)

// E到字节集 将任意类型转换为字节切片。
//
// 参数：
//   - value: 任意类型的值
//
// 返回值：
//   - []byte: 转换后的字节切片
//
// 示例：
//
//	E到字节集("hello") // 返回 []byte("hello")
//	E到字节集(123)     // 返回 []byte{123}
func E到字节集(value interface{}) []byte {
	return gconv.Bytes(value)
}

// E到字节 将任意类型转换为字节。
//
// 参数：
//   - value: 任意类型的值
//
// 返回值：
//   - byte: 转换后的字节值
//
// 示例：
//
//	E到字节(65)     // 返回 byte(65)
//	E到字节("A")    // 返回 byte(65)
func E到字节(value interface{}) byte {
	return gconv.Byte(value)
}

// E到整数 将任意类型转换为int64整数。
//
// 参数：
//   - value: 任意类型的值
//
// 返回值：
//   - int64: 转换后的整数值
//
// 示例：
//
//	E到整数("123")   // 返回 int64(123)
//	E到整数(45.67)   // 返回 int64(45)
func E到整数(value interface{}) int64 {
	return gconv.Int64(value)
}

// E到数值 将任意类型转换为float64数值。
//
// 参数：
//   - value: 任意类型的值
//
// 返回值：
//   - float64: 转换后的浮点数值
//
// 示例：
//
//	E到数值("123.45") // 返回 float64(123.45)
//	E到数值(123)      // 返回 float64(123)
func E到数值(value interface{}) float64 {
	return gconv.Float64(value)
}

// E到文本 将任意类型转换为字符串。
//
// 参数：
//   - value: 任意类型的值
//
// 返回值：
//   - string: 转换后的字符串
//
// 示例：
//
//	E到文本(123)      // 返回 "123"
//	E到文本(45.67)    // 返回 "45.67"
//	E到文本(true)     // 返回 "true"
func E到文本(value interface{}) string {
	return gconv.String(value)
}

// E到结构体 将任意类型转换为指定的结构体。
//
// 参数：
//   - 待转换的参数: 任意类型的值（通常是map或JSON字符串）
//   - 结构体指针: 目标结构体的指针
//
// 返回值：
//   - error: 转换错误信息
//
// 示例：
//
//	type User struct {
//	    Name string `json:"name"`
//	    Age  int    `json:"age"`
//	}
//	var user User
//	E到结构体(`{"name":"张三","age":25}`, &user)
func E到结构体(待转换的参数 interface{}, 结构体指针 interface{}) error {
	return gconv.Struct(待转换的参数, 结构体指针)
}
