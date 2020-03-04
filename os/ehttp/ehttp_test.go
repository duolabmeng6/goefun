package ehttp

import (
	"testing"
)

func TestNewHttp(t *testing.T) {
	http := NewHttp()
	ret, flag := http.Get("https://www.taobao.com/help/getip.php")
	if flag{
		t.Log("访问失败", ret)
	}else{
		t.Log(flag, ret)
	}

	ret, flag = http.Get("http://ip.taobao.com/service/getIpInfo.php?ip=myip")
	if flag{
		t.Log("访问失败", ret)
	}else{
		t.Log(flag, ret)
	}

}
