package etool

import (
	"github.com/duolabmeng6/goefun/ecore"
	"testing"
)

func TestEJson(t *testing.T) {
	json := New存取键值表()
	json.Set("a", "b")
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
	//json.E删除("data.0")
	ecore.E调试输出(json.ToJson(true))
}
