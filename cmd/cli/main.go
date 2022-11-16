package main

import (
	_ "github.com/gowebspider/renderpool/configs"
	configs "github.com/gowebspider/renderpool/configs/dev"
	"github.com/spf13/viper"
)

type Cfg struct {
	configs.RenderConfig
}

var cfg Cfg

func main() {
	// load configuration file
	if err := viper.Unmarshal(&cfg); err != nil {
		panic(err)
	}

	// link redis backend
	if err := cfg.RenderBackend.RedisConfig.Init(); err != nil {
		panic(err)
	}

	// link mongo backend
	if err := cfg.RenderBackend.MongoConfig.Init(); err != nil {
		panic(err)
	}
}
