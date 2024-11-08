package main

import (
	"flag"
	"wecom-app-to-dify/internal/config"
	"wecom-app-to-dify/internal/log_util"
	"wecom-app-to-dify/internal/server"
)

func main() {
	logDir := flag.String("logdir", "logs", "log dir")
	confDir := flag.String("confdir", "configs", "conf dir")
	flag.Parse()

	log_util.Init(*logDir)
	config.LoadConfig(*confDir)

	server.Run()
}
