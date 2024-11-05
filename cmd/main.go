package main

import (
	"github.com/blankbro/wecom-app-svr"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"wecom-app-svr-sample/configs"
)

func main() {
	// 读取配置文件
	bytes, err := os.ReadFile("configs/config.yml")
	if err != nil {
		bytes, err = os.ReadFile("config.yml")
		if err != nil {
			log.Printf("读取配置文件失败: %v\n", err)
			return
		}
	}

	config := config.Config{}
	err = yaml.Unmarshal(bytes, &config)
	if err != nil {
		log.Println("解析 yaml 文件失败：", err)
		return
	}

	log.Printf("config server → %+v\n", config.Server)
	log.Printf("config wecom → %+v\n", config.Wecom)

	wecom_app_svr.Run(
		config.Server.Port, config.Wecom.Path,
		config.Wecom.Token, config.Wecom.AesKey, config.Wecom.CorpId,
		func(msgContent wecom_app_svr.MsgContent) {
			// 编写自己的逻辑
		},
	)
}
