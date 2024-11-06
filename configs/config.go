package config

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	Server Server
	Wecom  WeCom
	Dify   Dify
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

type Dify struct {
	Host   string `yaml:"host"`
	ApiKey string `yaml:"api_key"`
}

func LoadConfig(configPath string) Config {
	logrus.Infof("Using config file: %s", configPath)

	bytes, err := os.ReadFile(configPath)
	if err != nil {
		logrus.Fatalf("read config file error: %s", err.Error())
	}

	configObj := Config{}
	err = yaml.Unmarshal(bytes, &configObj)
	if err != nil {
		logrus.Fatalf("unmarshal config file error: %s", err.Error())
	}

	logrus.Infof("config server → %+v", configObj.Server)
	logrus.Infof("config wecom → %+v", configObj.Wecom)
	logrus.Infof("config dify → %+v", configObj.Dify)
	return configObj
}
