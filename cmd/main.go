package main

import (
	"github.com/blankbro/wecom-app-svr"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"path/filepath"
)

type Config struct {
	Server Server
	Wecom  WeCom
}

type Server struct {
	Port string
}

type WeCom struct {
	Token  string
	AesKey string `yaml:"aes_key"`
	CorpId string `yaml:"corp_id"`
	Path   string
}

func main() {
	var configPath = os.Getenv("config_path")
	if configPath == "" {
		log.Println("config_path is empty, 从可执行文件所在目录读取配置文件")
		file, err := os.Executable()
		if err != nil {
			log.Println("无法获取可执行文件路径：", err)
			return
		}
		path, err := filepath.Abs(file)
		if err != nil {
			log.Println("无法获取文件绝对路径：", err)
			return
		}
		dir := filepath.Dir(path)
		configPath = filepath.Join(dir, "config.yaml")
	}
	log.Printf("configPath is %s", configPath)

	// 读取配置文件
	bytes, err := os.ReadFile(configPath)
	if err != nil {
		log.Printf("读取配置文件失败: %v\n", err)
		return
	}

	config := Config{}
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
