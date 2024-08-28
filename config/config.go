package config

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"os"
)

var config *Config

type Config struct {
	service.ServiceConf
	ListenOn      string
	ElasticSearch struct {
		Username string
		Password string
		Addr     []string
	}
}

func Init() {
	config = new(Config)
	path := os.Getenv("CONFIG_PATH")
	if path == "" {
		path = "etc/config.yaml"
	}
	err := conf.Load(path, config)
	if err != nil {
		panic(err)
	}
	err = config.SetUp()
	if err != nil {
		panic(err)
	}
}

func Get() *Config {
	if config == nil {
		fmt.Println("config is nil")
	}
	return config
}
