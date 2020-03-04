package ehttp

import (
	"testing"
)

func TestNewHttp(t *testing.T) {
	http := NewHttp()
	ret, flag := http.Get("http://ip.taobao.com/service/getIpInfo.php?ip=myip", "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.86 Safari/537.36")

	t.Log(flag, ret)
}
