package dingding

import (
	"os"
	"testing"
)

func TestNew钉钉机器人(t *testing.T) {
	token := os.Getenv("DINGDING_TOKEN")
	secret := os.Getenv("DINGDING_SECRET")
	bot := New钉钉机器人(token, secret)

	bot.E发送文本消息("hello world", 3, "")
	bot.E发送markdown消息("# hello world", "标题", 3, "")
	//bot.E发送链接消息("# hello world", "标题", "https://www.baidu.com", "https://www.baidu.com/img/PCtm_d9c8750bed0b3c7d089fa7d55720d6cf.png")
	//bot.E整体跳转消息类型("标题", "内容", "https://www.baidu.com")

}
