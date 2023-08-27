# GoEfun 易函数

[![github stars](https://img.shields.io/github/stars/duolabmeng6/pyefun?style=social)](https://github.com/duolabmeng6/pyefun)

goefun 是用于golang的中文函数库，它提供了强大而易于使用的函数，它完整封装了易语言核心支持库的所有功能，同时提供简单易用的函数。

# 特性

- 可在Windows、Linux和Mac系统上完美运行
- 兼容Go开发环境
- 所有函数名和注释都使用中文
- 封装了100+命令
- 所有命令都经过100%测试，可放心使用
- 提供了测试用例以方便查看和使用
- 使用方便封装了大量文本处理函数


## 使用方法

### 引入模块

```go
package main

import (
	. "github.com/duolabmeng6/goefun/ecore"
	. "github.com/duolabmeng6/goefun/ehttp"
	. "github.com/duolabmeng6/goefun/etool"
)

func main() {
	E调试输出("欢迎使用 go-efun")
	http := NewHttp()
	返回文本, 访问失败 := http.Get("https://www.taobao.com/help/getip.php")
	if 访问失败 {
		E调试输出("访问失败", 返回文本)
	}
	E调试输出(返回文本)

	data := New存取键值表()
	data.Set("aaa", "111")
	data.Set("bbb", "222")
	data.Set("ccc", "333")
	data.SetArray("list", "a")
	data.SetArray("list", "b")
	data.SetArray("list", "c")

	E调试输出(data.ToJson(true))

	E调试输出(E取md5从文本("1234567"))

	E调试输出(E取现行时间().E时间到文本("Y-m-d H:i:s"))

}


```

# 文档

https://pkg.go.dev/github.com/duolabmeng6/goefun#section-sourcefiles
