package ecore

import (
    "github.com/gogf/gf/v2/util/gconv"
)

// E到字节集 将任意类型转换为字节集
//
// 参数：
//   value interface{} - 欲转换的值
//
// 返回值：
//   []byte - 转换后的字节集
func E到字节集(value interface{}) []byte {
    return gconv.Bytes(value)
}

// E到字节 将任意类型转换为字节
//
// 参数：
//   value interface{} - 欲转换的值
//
// 返回值：
//   byte - 转换后的字节
func E到字节(value interface{}) byte {
    return gconv.Byte(value)
}

// E到整数 将任意类型转换为整数
//
// 参数：
//   value interface{} - 欲转换的值
//
// 返回值：
//   int64 - 转换后的整数值
func E到整数(value interface{}) int64 {
    return gconv.Int64(value)
}

// E到数值 将任意类型转换为双精度小数
//
// 参数：
//   value interface{} - 欲转换的值
//
// 返回值：
//   float64 - 转换后的双精度小数值
func E到数值(value interface{}) float64 {
    return gconv.Float64(value)
}

// E到文本 将任意类型转换为文本
//
// 参数：
//   value interface{} - 欲转换的值
//
// 返回值：
//   string - 转换后的文本
func E到文本(value interface{}) string {
    return gconv.String(value)
}

// E到结构体 将map或其他类型转换为结构体
//
// 参数：
//   待转换的参数 interface{} - 欲转换的原始数据
//   结构体指针 interface{} - 目标结构体指针
//
// 返回值：
//   error - 转换过程中的错误信息
func E到结构体(待转换的参数 interface{}, 结构体指针 interface{}) error {
    return gconv.Struct(待转换的参数, 结构体指针)
}
