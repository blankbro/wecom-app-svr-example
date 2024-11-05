package config

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"strings"
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

func LoadConfig() Config {
	configDirs := []string{"configs/config.yml", "config.yml", "../../configs/config.yml"}
	var bytes []byte
	var err error
	for _, configDir := range configDirs {
		bytes, err = os.ReadFile(configDir)
		if err == nil {
			break
		}
	}
	if err != nil {
		log.Fatalf("未加载到配置文件, 加载顺序: %s", strings.Join(configDirs, ", "))
	}

	configObj := Config{}
	err = yaml.Unmarshal(bytes, &configObj)
	if err != nil {
		log.Fatalf("解析 yml 文件失败: %v", err)
	}
	return configObj
}
