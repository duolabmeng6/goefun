package coreUtil

import (
	"fmt"
	. "github.com/duolabmeng6/goefun/core"
	"math/rand"
	"time"
)

//
//import . "github.com/duolabmeng6/goefun/os/ehttp"
//
//func E获取本机ip() string {
//	http := NewHttp()
//	ret, flag := http.Get("http://ip.taobao.com/service/getIpInfo.php?ip=myip")
//	if flag == false {
//		return ""
//	}
//	json := New存取键值表()
//	json.LoadFromJsonString(ret)
//	return json.GetString("data.ip")
//}
//
//func E获取ip信息(ip string) string {
//	http := NewHttp()
//	ret, flag := http.Get("http://ip.taobao.com/service/getIpInfo.php?ip=" + ip)
//	if flag == false {
//		return ""
//	}
//	json := New存取键值表()
//	json.LoadFromJsonString(ret)
//
//	return E格式化文本("%s %s %s %s",
//		json.GetString("data.country"),
//		json.GetString("data.region"),
//		json.GetString("data.city"),
//		json.GetString("data.isp"),
//	)
//}

func E取随机ip() string {
	rand.Seed(time.Now().Unix())
	ip := fmt.Sprintf("%d.%d.%d.%d", E取随机数(50, 254), E取随机数(50, 254), E取随机数(50, 254), E取随机数(50, 254))
	return ip
}
