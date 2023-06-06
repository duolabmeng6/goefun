package dingding

import (
	"github.com/blinkbean/dingtalk"
)

type E钉钉机器人 struct {
	robot *dingtalk.DingTalk
}

func New钉钉机器人(token, secret string) *E钉钉机器人 {
	robot := dingtalk.InitDingTalkWithSecret(token, secret)
	return &E钉钉机器人{robot: robot}
}

func (d *E钉钉机器人) E发送文本消息(消息内容 string) {
	d.robot.SendTextMessage(消息内容)
}

func (d *E钉钉机器人) E发送链接消息(消息内容, 消息标题, 消息URL, 消息图片URL string) {
	d.robot.SendLinkMessage(消息标题, 消息内容, 消息图片URL, 消息URL)
}

func (d *E钉钉机器人) E发送markdown消息(消息内容 string) {
	d.robot.SendMarkDownMessage("markdown", 消息内容)
}

func (d *E钉钉机器人) E整体跳转消息类型(消息标题, 消息内容, 消息URL string) {
	按钮组 := []dingtalk.ActionCardMultiBtnModel{{
		Title:     "查看详情",
		ActionURL: 消息URL,
	}}
	d.robot.SendActionCardMessage(消息标题, 消息内容, dingtalk.WithCardBtns(按钮组))
}
