package etts

import (
	"testing"
)

func Test_newEJson(t *testing.T) {
	etts := NewETTS("./tmp")
	文本转语音, err := etts.E文本转语音("你好")
	if err != nil {
		return
	}
	println(文本转语音)
	毫秒, err := E取MP3时间(文本转语音)
	if err != nil {
		return
	}
	println(毫秒)

}
