package main

import (
	"github.com/blankbro/wecom-app-svr"
	"log"
	"wecom-app-svr-sample/configs"
)

func main() {
	configObj := config.LoadConfig()
	log.Printf("config server → %+v\n", configObj.Server)
	log.Printf("config wecom → %+v\n", configObj.Wecom)

	wecom_app_svr.Run(
		configObj.Server.Port, configObj.Wecom.Path,
		configObj.Wecom.Token, configObj.Wecom.AesKey, configObj.Wecom.CorpId,
		func(msgContent wecom_app_svr.MsgContent) {
			// 编写自己的逻辑
		},
	)
}
