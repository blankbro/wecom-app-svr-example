package server

import (
	"github.com/blankbro/wecom-app-svr"
	"wecom-app-to-dify/internal/config"
)

func Run() {

	initDifyClient()
	loadUserConversation()

	configObj := config.Obj
	wecom_app_svr.Run(
		configObj.Server.Port, configObj.Wecom.Path,
		configObj.Wecom.Token, configObj.Wecom.AesKey, configObj.Wecom.CorpId,
		msgHandler,
	)
}
