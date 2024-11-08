package server

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"wecom-app-to-dify/internal/config"
)

var userConversation map[string]string
var fileName string

func setConversationId(username string, conversationId string) {
	userConversation[username] = conversationId
	flushUserConversation()
}

func getConversationId(username string) (string, bool) {
	conversationId, ok := userConversation[username]
	if !ok {
		return "", false
	}
	return conversationId, true
}

func clearConversationId(username string) {
	userConversation[username] = ""
	flushUserConversation()
}

func loadUserConversation() {
	fileName = config.Dir + "/user_conversation.json"
	logrus.Infof("loading user conversation, from file: %s", fileName)

	_, statErr := os.Stat(fileName)
	if statErr != nil {
		if os.IsNotExist(statErr) {
			file, creatErr := os.Create(fileName)
			if creatErr != nil {
				logrus.Fatalf("failed to create file, %v", creatErr)
			}
			file.WriteString("{}")
			file.Close()
		} else {
			logrus.Fatalf("failed to stat file, %v", statErr)
			return
		}
	}

	// 打开文件
	file, err := os.Open(fileName)
	if err != nil {
		logrus.Fatalf("opening file error, %v", err)
	}
	defer file.Close()

	// 读取文件内容
	jsonData, err := io.ReadAll(file)
	if err != nil {
		logrus.Fatalf("reading file error, %v", err)
	}

	// 反序列化JSON为map
	err = json.Unmarshal(jsonData, &userConversation)
	if err != nil {
		logrus.Fatalf("unmarshalling JSON to map error, %v", err)
	}

	// 打印map长度
	logrus.Infof("user conversation size: %d", len(userConversation))
}

func flushUserConversation() {
	// 序列化map为JSON
	jsonData, err := json.Marshal(userConversation)
	if err != nil {
		logrus.Errorf("marshalling map to JSON error, %v", err)
		return
	}

	// 创建文件
	file, err := os.Create(fileName)
	if err != nil {
		logrus.Errorf("creating file error, %v", err)
		return
	}
	defer file.Close()

	// 写入JSON数据到文件
	_, err = file.Write(jsonData)
	if err != nil {
		logrus.Errorf("writing to file error, %v", err)
		return
	}

	logrus.Infof("user conversation successfully written to file")
}
