package eDingtalkBot

import (
	"context"
	"github.com/open-dingtalk/dingtalk-stream-sdk-go/chatbot"
	"github.com/open-dingtalk/dingtalk-stream-sdk-go/client"
	"github.com/open-dingtalk/dingtalk-stream-sdk-go/logger"
)

type EDingtalkBotI interface {
	E创建机器人(clientId string, clientSecret string, onChatBotMessageReceived func(ctx context.Context, data *chatbot.BotCallbackDataModel) ([]byte, error)) error
}

type EDingtalkBot struct {
	clientId     string
	clientSecret string
	voiceMode    bool
}

func NewEDingtalkBot() *EDingtalkBot {
	return &EDingtalkBot{}
}

func (a *EDingtalkBot) E创建机器人(clientId string, clientSecret string, onChatBotMessageReceived func(ctx context.Context, data *chatbot.BotCallbackDataModel) ([]byte, error)) error {
	a.clientId = clientId
	a.clientSecret = clientSecret

	logger.SetLogger(logger.NewStdTestLoggerWithDebug())
	cli := client.NewStreamClient(client.WithAppCredential(client.NewAppCredentialConfig(a.clientId, a.clientSecret)))
	cli.RegisterChatBotCallbackRouter(onChatBotMessageReceived)
	err := cli.Start(context.Background())
	if err != nil {
		return err
	}
	defer cli.Close()
	select {}
}
