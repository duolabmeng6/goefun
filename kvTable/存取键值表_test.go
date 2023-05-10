package os

import (
	"testing"

	E "github.com/duolabmeng6/goefun/eCore"
)

func TestEJson(t *testing.T) {
	json := New存取键值表()
	json.SetArray("data", H{
		"name": "1",
		"link": "1",
	})
	json.SetArray("data", H{
		"name": "2",
		"link": "2",
	})
	json.SetArray("data", H{
		"name": "3",
		"link": "3",
	})
	json.E删除("data.0")
	E.E调试输出(json.ToJson(true))
}
