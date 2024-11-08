package main

import (
	"flag"
	"wecom-app-to-dify/config"
	"wecom-app-to-dify/server"
	"wecom-app-to-dify/utils/log_util"
)

func main() {
	configFilepath := flag.String("config", "config.yml", "config file path")
	logPath := flag.String("logpath", "logs", "log file path")
	flag.Parse()

	log_util.Init(*logPath)
	config.LoadConfig(*configFilepath)

	server.Run()
}
