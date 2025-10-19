// Package ejs js引擎 用于运行js代码
package ejs

import (
    "github.com/robertkrimen/otto"
)

// RunJs 运行一段 JavaScript 代码并调用给定的函数名，返回其字符串结果。
// content 为待执行的 JS 源码，functionName 为要调用的函数名，args 为可选的函数参数。
func RunJs(content string, functionName string, args ...interface{}) string {
    var err error
    vm := otto.New()
    _, err = vm.Run(content)
    if err != nil {
        panic(err)
    }
    value, err := vm.Call(functionName, nil, args...)
    if err != nil {
        panic(err)
    }
    return value.String()
}
