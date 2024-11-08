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

var Obj Config

func LoadConfig(configPath string) Config {
	logrus.Infof("Using config file: %s", configPath)

	bytes, err := os.ReadFile(configPath)
	if err != nil {
		logrus.Fatalf("read config file error: %s", err.Error())
	}

	Obj = Config{}
	err = yaml.Unmarshal(bytes, &Obj)
	if err != nil {
		logrus.Fatalf("unmarshal config file error: %s", err.Error())
	}

	logrus.Infof("config server → %+v", Obj.Server)
	logrus.Infof("config wecom → %+v", Obj.Wecom)
	logrus.Infof("config dify → %+v", Obj.Dify)
	return Obj
}
