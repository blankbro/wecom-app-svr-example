package main

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	"github.com/blankbro/wecom-app-svr"
	"github.com/langgenius/dify-sdk-go"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"strings"
	"time"
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
		func(w http.ResponseWriter, msgContent wecom_app_svr.MsgContent) {
			logrus.Infof("msgContent: %+v", msgContent)
			if msgContent.Content == "" {
				logrus.Error("msgContent.Content is empty")
				return
			}
			if strings.HasPrefix(msgContent.Content, "#") {
				if msgContent.Content == "#clear" {
					userConversationId[msgContent.FromUsername] = ""
					timestamp := time.Now().Unix()
					responseMsgContent := wecom_app_svr.MsgContent{
						FromUsername: msgContent.ToUsername,
						ToUsername:   msgContent.FromUsername,
						AgentId:      msgContent.AgentId,
						CreateTime:   uint32(timestamp),
						MsgType:      "text",
						MsgId:        msgContent.MsgId,
						Content:      "已重置",
					}
					timestampStr := strconv.FormatInt(timestamp, 10)
					nonce := uuid.NewV4().String()
					encryptedResponseMsgContentBytes, encryptErr := wecom_app_svr.EncryptMsgContent(responseMsgContent, timestampStr, nonce)
					if encryptErr != nil {
						logrus.Errorf("encryptMsgContent err: %v", encryptErr)
						w.Write(bytes.NewBufferString(fmt.Sprintf("encryptMsgContent err: %v", encryptErr)).Bytes())
					} else {
						w.Write(encryptedResponseMsgContentBytes)
					}
					return
				}
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
