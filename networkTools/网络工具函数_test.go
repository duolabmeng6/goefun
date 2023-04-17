package os

import "testing"

func TestE获取本机ip(t *testing.T) {
	t.Log("获取本机ip", E获取本机ip())
	t.Log("E获取ip信息", E获取ip信息(E获取本机ip()))

}
