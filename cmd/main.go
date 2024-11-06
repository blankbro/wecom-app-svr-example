package main

import (
	"context"
	"flag"
	"github.com/blankbro/wecom-app-svr"
	"github.com/langgenius/dify-sdk-go"
	"github.com/sirupsen/logrus"
	"wecom-app-svr-sample/configs"
	"wecom-app-svr-sample/utils/log_util"
)

var userConversationId = map[string]string{}

func init() {
	log_util.Init()
}

func main() {
	configPath := flag.String("config", "config.yml", "config file path")
	flag.Parse()

	configObj := config.LoadConfig(*configPath)

	difyClientConfig := &dify.ClientConfig{
		Host:             configObj.Dify.Host,
		DefaultAPISecret: configObj.Dify.ApiKey,
	}

	difyClient := dify.NewClientWithConfig(difyClientConfig)

	wecom_app_svr.Run(
		configObj.Server.Port, configObj.Wecom.Path,
		configObj.Wecom.Token, configObj.Wecom.AesKey, configObj.Wecom.CorpId,
		func(msgContent wecom_app_svr.MsgContent) {
			logrus.Infof("msgContent: %+v", msgContent)
			if msgContent.Content == "" {
				logrus.Error("msgContent.Content is empty")
				return
			}
			if msgContent.Content == "#clear" {
				userConversationId[msgContent.FromUsername] = ""
				return
			}
			conversationId := userConversationId[msgContent.FromUsername]
			resp, err := difyClient.Api().ChatMessages(context.Background(), &dify.ChatMessageRequest{
				Query:          msgContent.Content,
				User:           msgContent.FromUsername,
				ConversationID: conversationId,
			})
			if err != nil {
				logrus.Fatalf(err.Error())
			}
			logrus.Infof("resp: %+v", resp)
			if conversationId == "" {
				userConversationId[msgContent.FromUsername] = resp.ConversationID
			}
		},
	)
}
