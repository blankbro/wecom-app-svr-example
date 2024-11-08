package server

import (
	"context"
	"github.com/blankbro/wecom-app-svr"
	"github.com/langgenius/dify-sdk-go"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"strings"
	"time"
	"wecom-app-to-dify/config"
)

var difyClient *dify.Client

func initMsgHandler() {
	configObj := config.Obj

	difyClientConfig := &dify.ClientConfig{
		Host:             configObj.Dify.Host,
		DefaultAPISecret: configObj.Dify.ApiKey,
	}

	difyClient = dify.NewClientWithConfig(difyClientConfig)
}

func msgHandler(w http.ResponseWriter, msgContent wecom_app_svr.MsgContent) {
	logrus.Infof("msgContent: %+v", msgContent)
	if msgContent.Content == "" {
		logrus.Error("msgContent.Content is empty")
		return
	}
	if strings.HasPrefix(msgContent.Content, "#") {
		if msgContent.Content == "#clear" {
			clearConversationId(msgContent.FromUsername)
			replyText(w, msgContent, "已重置")
			return
		}
		if msgContent.Content == "#get" {
			conversationId, ok := getConversationId(msgContent.FromUsername)
			if !ok {
				replyText(w, msgContent, "当前没有任何会话")
			} else {
				replyText(w, msgContent, "当前会话ID为: "+conversationId)
			}
			return
		}
	}

	go func() {
		conversationId, ok := getConversationId(msgContent.FromUsername)
		resp, err := difyClient.Api().ChatMessages(context.Background(), &dify.ChatMessageRequest{
			Query:          msgContent.Content,
			User:           msgContent.FromUsername,
			ConversationID: conversationId,
		})
		if err != nil {
			logrus.Fatalf(err.Error())
		}
		logrus.Infof("resp: %+v", resp)
		if !ok {
			setConversationId(msgContent.FromUsername, resp.ConversationID)
		}
	}()
}

func replyText(w http.ResponseWriter, fromMsg wecom_app_svr.MsgContent, replyText string) {
	replyMsg := wecom_app_svr.MsgContent{
		FromUsername: fromMsg.ToUsername,
		ToUsername:   fromMsg.FromUsername,
		AgentId:      fromMsg.AgentId,
		CreateTime:   time.Now().Unix(),
		MsgType:      "text",
		MsgId:        fromMsg.MsgId,
		Content:      replyText,
	}
	replyBytes, encryptErr := wecom_app_svr.EncryptMsgContent(
		replyMsg,
		strconv.FormatInt(replyMsg.CreateTime, 10),
		uuid.NewV4().String(),
	)
	if encryptErr != nil {
		logrus.Errorf("encryptMsgContent err: %v", encryptErr)
		w.Write([]byte("我暂时遇到了一些问题，请您稍后重试~"))
	} else {
		w.Write(replyBytes)
	}
}
