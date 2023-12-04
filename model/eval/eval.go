package eval

import (
	"encoding/json"
	"github.com/gogf/gf/v2/util/gconv"
)

func E到字节集(value interface{}) []byte {
	return gconv.Bytes(value)
}
func E到字节(value interface{}) byte {
	return gconv.Byte(value)
}
func E到整数(value interface{}) int64 {
	return gconv.Int64(value)
}
func E到数值(value interface{}) float64 {
	return gconv.Float64(value)
}
func E到文本(value interface{}) string {
	return gconv.String(value)
}

func E到Json(v interface{}) string {
	ujs, _ := json.Marshal(v)
	return string(ujs)
}

func E到Json美化(v interface{}) string {
	ujs, err := json.MarshalIndent(v, "", "    ")
	if err != nil {
		return ""
	}
	return string(ujs)
}
