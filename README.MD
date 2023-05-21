# GoEfun

为golang提供强大且易用的中文函数库，易语言go函数库，完整的封装了易语言核心支持库所有功能，以及易语言中简单易用的函数

# 特性

- 完全兼容window linux mac运行
- 完全兼容golang运行环境
- 全中文函数名和注释
- 封装了超100+命令
- 全部命令经过测试 100% 测试 放心使用
- 所有命令均有测试用例可以查看后立即使用
- 封装了超丰富的文本处理函数,简单,易用


## 使用方法

### 引入模块

```go
package main

import (
	. "github.com/duolabmeng6/go-efun/ecore"
	. "github.com/duolabmeng6/go-efun/ehttp"
	. "github.com/duolabmeng6/go-efun/etool"
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