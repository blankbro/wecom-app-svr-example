package main

import (
	"flag"
	"wecom-app-to-dify/internal/config"
	"wecom-app-to-dify/internal/log_util"
	"wecom-app-to-dify/internal/server"
)

func main() {
	logPath := flag.String("logpath", "logs", "log path")
	flag.Parse()

	log_util.Init(*logPath)
	config.LoadConfig(config.ConfigYamlPath)

	server.Run()
}
