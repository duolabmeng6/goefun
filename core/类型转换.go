package core

import (
	"github.com/axgle/mahonia"
	"github.com/gogf/gf/util/gconv"
)

//E文本编码转换("测试一下","gbk","utf-8")
func E文本编码转换(src string, oldEncoder string, newEncoder string) string {
	srcDecoder := mahonia.NewDecoder(oldEncoder)
	desDecoder := mahonia.NewDecoder(newEncoder)
	resStr := srcDecoder.ConvertString(src)
	_, resBytes, _ := desDecoder.Translate([]byte(resStr), true)
	return string(resBytes)
}

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
func E到结构体(待转换的参数 interface{}, 结构体指针 interface{}) error {
	return gconv.Struct(待转换的参数, 结构体指针)
}
