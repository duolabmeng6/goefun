package E

import (
	"encoding/json"

	"github.com/tidwall/gjson"
)

//https://github.com/tidwall/gjson

//简化用法
func Json解析文本(data string, path string) string {
	return gjson.Get(data, path).String()
}
func Json解析(data string, path string) gjson.Result {
	return gjson.Get(data, path)
}

func E到Json(v interface{}) string {
	ujs, _ := json.Marshal(v)
	return string(ujs)
}
