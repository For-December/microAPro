package config

import (
	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
	"log"
)

//var EnvCfg envConfig

var EnvCfg = struct {
}{}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Can not read env from file system, please check the right this program owned.")
	}

	//EnvCfg = envConfig{}

	if err := env.Parse(&EnvCfg); err != nil {
		panic("Can not parse env from file system, please check the env.")
	}
}
