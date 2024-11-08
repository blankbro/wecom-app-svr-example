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

func LoadConfig(configFilepath string) Config {
	logrus.Infof("config file: %s", configFilepath)
	bytes, err := os.ReadFile(configFilepath)
	if err != nil {
		logrus.Fatalf("read config file error, %s", err)
	}

	Obj = Config{}
	err = yaml.Unmarshal(bytes, &Obj)
	if err != nil {
		logrus.Fatalf("unmarshal config file error, %s", err)
	}

	logrus.Infof("server config → %+v", Obj.Server)
	logrus.Infof("wecom config → %+v", Obj.Wecom)
	logrus.Infof("dify config → %+v", Obj.Dify)
	return Obj
}
