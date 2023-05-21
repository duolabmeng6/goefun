package etool

import (
	"encoding/json"

	"github.com/tidwall/gjson"
)

//https://github.com/tidwall/gjson

// Json解析文本 直接返回解析文本
func Json解析文本(data string, path string) string {
	return gjson.Get(data, path).String()
}

// Json解析 解析对象
func Json解析(data string, path string) gjson.Result {
	return gjson.Get(data, path)
}

// E到Json
func E到Json(v interface{}) string {
	ujs, _ := json.Marshal(v)
	return string(ujs)
}
