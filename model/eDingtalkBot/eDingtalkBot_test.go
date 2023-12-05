package eDingtalkBot

import (
	"context"
	"fmt"
	"github.com/duolabmeng6/goefun/ecore"
	"github.com/duolabmeng6/goefun/etool"
	"github.com/duolabmeng6/goefun/model/etts"
	"github.com/open-dingtalk/dingtalk-stream-sdk-go/chatbot"
	"strconv"
	"testing"
)

func Test_newEJson(t *testing.T) {
	ecore.E加载环境变量_从文件(ecore.E取运行源文件路径() + "/.env")
	println(ecore.E读环境变量("clientId"))
	bot := NewEDingtalkBot()
	clientId := ecore.E读环境变量("clientId")
	clientSecret := ecore.E读环境变量("clientSecret")
	tts := etts.NewETTS(ecore.E取运行源文件路径() + "/tmp")

	bot.E创建机器人(ecore.E读环境变量("clientId"), ecore.E读环境变量("clientSecret"),
		func(ctx context.Context, data *chatbot.BotCallbackDataModel) ([]byte, error) {
			chatbotReplier := chatbot.NewChatbotReplier()
			收到消息 := ""
			if data.Msgtype == "text" {
				收到消息 = ecore.E删首尾空(data.Text.Content)
			} else if data.Msgtype == "audio" {
				ejson := etool.NewJson()
				ejson.LoadFromJsonString(etool.E到Json(data.Content))
				收到消息 = ecore.E删首尾空(ejson.GetString("recognition"))
			}

			//发送文本消息
			chatbotReplier.SimpleReplyText(ctx, data.SessionWebhook, []byte("已收到"+收到消息))

			//生成语音消息
			filepath, err := tts.E文本转语音(收到消息)
			if err != nil {
				fmt.Printf("语音合成失败 %s", err.Error())
			}
			token := E钉钉取Token(clientId, clientSecret)
			mid, err := E钉钉上传文件(token, filepath)
			if err != nil {
				fmt.Printf("钉钉语音文件上传失败 %s", filepath)
				return []byte(""), nil
			}
			duration, _ := etts.E取MP3时间(filepath)
			requestBody := map[string]interface{}{
				"msgtype": "audio",
				"audio": map[string]interface{}{
					"mediaId":  mid,
					"duration": strconv.Itoa(duration),
				},
			}
			chatbotReplier.ReplyMessage(ctx, data.SessionWebhook, requestBody)

			return []byte(""), nil
		})
}
