package dify

import (
	"context"
	"github.com/langgenius/dify-sdk-go"
	"github.com/sirupsen/logrus"
	"strings"
	"testing"
	"wecom-app-to-dify/internal/config"
	"wecom-app-to-dify/internal/log_util"
)

var client *dify.Client

func init() {
	log_util.Init("logs")
	configObj := config.LoadConfig("../../configs/config.yml")
	var difyConfig = &dify.ClientConfig{Host: configObj.Dify.Host, DefaultAPISecret: configObj.Dify.ApiKey}
	client = dify.NewClientWithConfig(difyConfig)
}

func TestChatMessages(t *testing.T) {
	ctx := context.Background()
	resp, err := client.Api().ChatMessages(ctx, &dify.ChatMessageRequest{
		Query: "hello",
		User:  "blankbro",
	})
	if err != nil {
		logrus.Fatalf(err.Error())
	}
	logrus.Infof("resp: %+v", resp)
}

func TestChatMessagesStream(t *testing.T) {
	ctx := context.Background()
	ch, err := client.Api().ChatMessagesStream(ctx, &dify.ChatMessageRequest{
		Query: "hello",
		User:  "blankbro",
	})
	if err != nil {
		logrus.Fatalf(err.Error())
	}

	var (
		strBuilder strings.Builder
		cId        string
	)
	for {
		select {
		case <-ctx.Done():
			logrus.Infof("ctx.Done %s", strBuilder.String())
			return
		case r, isOpen := <-ch:
			if !isOpen {
				goto done
			}
			strBuilder.WriteString(r.Answer)
			cId = r.ConversationID
			logrus.Println("Answer2", r.Answer, r.ConversationID, cId, r.ID, r.TaskID)
		}
	}
done:
	logrus.Infof(strBuilder.String())
	logrus.Infof(cId)
}
