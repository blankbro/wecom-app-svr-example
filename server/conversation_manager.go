package server

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

var userConversation map[string]string
var fileName = "user_conversation.json"

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

func initUserConversation() {
	_, statErr := os.Stat(fileName)
	if statErr != nil {
		if os.IsNotExist(statErr) {
			file, creatErr := os.Create(fileName)
			if creatErr != nil {
				logrus.Fatalf("Failed to create file: %s", creatErr.Error())
			}
			file.WriteString("{}")
			file.Close()
		} else {
			logrus.Fatal("Failed to stat file: %v", statErr)
			return
		}
	}

	// 打开文件
	file, err := os.Open(fileName)
	if err != nil {
		logrus.Errorf("Error opening file: %v", err)
		return
	}
	defer file.Close()

	// 读取文件内容
	jsonData, err := io.ReadAll(file)
	if err != nil {
		logrus.Errorf("Error reading file: %v", err)
		return
	}

	// 反序列化JSON为map
	err = json.Unmarshal(jsonData, &userConversation)
	if err != nil {
		logrus.Errorf("Error unmarshalling JSON to map: %v", err)
		return
	}

	// 打印map内容
	logrus.Infof("Map read from file:")
	for key, value := range userConversation {
		logrus.Infof("%s: %s", key, value)
	}
}

func flushUserConversation() {
	// 序列化map为JSON
	jsonData, err := json.Marshal(userConversation)
	if err != nil {
		logrus.Errorf("Error marshalling map to JSON: %v", err)
		return
	}

	// 创建文件
	file, err := os.Create(fileName)
	if err != nil {
		logrus.Errorf("Error creating file: %v", err)
		return
	}
	defer file.Close()

	// 写入JSON数据到文件
	_, err = file.Write(jsonData)
	if err != nil {
		logrus.Errorf("Error writing to file: %v", err)
		return
	}

	logrus.Infof("Map successfully written to file")
}
