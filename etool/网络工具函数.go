package etool

import (
	"github.com/duolabmeng6/goefun/ecore"
	"github.com/duolabmeng6/goefun/ehttp"
)

// E获取本机ip
func E获取本机ip() string {
	http := ehttp.NewHttp()
	ret, flag := http.Get("https://www.taobao.com/help/getip.php")
	if flag {
		return ""
	}
	return ecore.StrCut(ret, `{ip:"$"}`)
}

// E获取本机ip
func E获取ip信息(ip string) string {
	http := ehttp.NewHttp()
	ret, flag := http.Get("http://ip.taobao.com/service/getIpInfo.php?ip=" + ip)
	if flag {
		return ""
	}
	json := New存取键值表()
	json.LoadFromJsonString(ret)
	return ecore.E格式化文本("%s %s %s %s",
		json.GetString("data.country"),
		json.GetString("data.region"),
		json.GetString("data.city"),
		json.GetString("data.isp"),
	)
}
