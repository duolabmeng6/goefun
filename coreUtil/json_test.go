package E

import (
	os "github.com/duolabmeng6/goefun/os/存取键值表"
	"testing"
)

func TestJson解析文本(t *testing.T) {
	jsonstr := `
{
	"title": "json在线解析（简版） -JSON在线解析",
	"json.url": "https://www.sojson.com/simple_json.html",
	"keywords": "json在线解析",
	"功能": ["JSON美化", "JSON数据类型显示", "JSON数组显示角标", "高亮显示", "错误提示", {
		"备注": ["www.sojson.com", "json.la"]
	}],
	"加入我们": {
		"qq群": "259217951"
	}
}
`
	t.Log("Json解析文本", Json解析文本(jsonstr, "title"))
	t.Log("Json解析文本", Json解析文本(jsonstr, `json\.url`)) //需要转义符号
	t.Log("Json解析文本", Json解析文本(jsonstr, "keywords"))
	t.Log("Json解析文本", Json解析文本(jsonstr, "功能.0"))
	t.Log("Json解析文本", Json解析文本(jsonstr, "功能.1"))
	t.Log("Json解析文本", Json解析文本(jsonstr, "功能"))
	t.Log("Json解析文本", Json解析文本(jsonstr, "功能.5.备注"))
	t.Log("Json解析文本", Json解析文本(jsonstr, "加入我们"))
	t.Log("Json解析文本", Json解析文本(jsonstr, "加入我们.qq群"))

	t.Log("Json解析文本", Json解析(jsonstr, "加入我们.qq群").Int())
	t.Log("Json解析文本", Json解析(jsonstr, "功能").Array())

}

func TestE到Json(t *testing.T) {
	t.Log("E到Json", E到Json(os.H{
		"aaa": 123,
		"bbb": 321,
		"ccc": os.H{
			"c1": 123,
			"c2": 321,
		},
	}))

}
