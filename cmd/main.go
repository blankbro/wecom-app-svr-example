package main

import (
	"flag"
	"wecom-app-to-dify/config"
	"wecom-app-to-dify/server"
	"wecom-app-to-dify/utils/log_util"
)

func init() {
	log_util.Init()
}

func main() {
	configPath := flag.String("config", "config.yml", "config file path")
	flag.Parse()

	config.LoadConfig(*configPath)

	server.Run()
}
