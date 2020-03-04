# GoEfun

为golang提供强大且易用的中文函数库，易语言go函数库，完整的封装了易语言核心支持库所有功能，以及易语言中简单易用的函数


## 使用方法

### 引入模块

```go
package main

import (
	. "github.com/duolabmeng6/goefun/core"
	. "github.com/duolabmeng6/goefun/coreUtil"
	. "github.com/duolabmeng6/goefun/os/ehttp"
	. "github.com/duolabmeng6/goefun/os/存取键值表"
)

func main() {
	E调试输出("欢迎使用go-efun")
	http := NewHttp()
	ret, flag := http.Get("https://www.taobao.com/help/getip.php")
	if flag {
		E调试输出("访问失败", ret)
	} else {
		E调试输出(flag, ret)
	}

	ret, flag = http.Get("http://ip.taobao.com/service/getIpInfo.php?ip=myip")
	if flag {
		E调试输出("访问失败", ret)
	} else {
		E调试输出(flag, ret)
	}

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

## 核心支持库
https://godoc.org/github.com/duolabmeng6/goefun/core

## 核心支持库工具函数库
https://godoc.org/github.com/duolabmeng6/goefun/coreUtil

## 存取键值表
https://godoc.org/github.com/duolabmeng6/goefun/os/存取键值表

## ehttp
https://godoc.org/github.com/duolabmeng6/goefun/os/ehttp





