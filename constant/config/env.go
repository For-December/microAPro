package config

import (
	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
	"log"
)

//var EnvCfg envConfig

var EnvCfg = struct {
	DouBaoApiKey   string `env:"DOU_BAO_API_KEY"`
	BaiDuAppId     string `env:"BAI_DU_APP_ID"`
	BaiDuSecretKey string `env:"BAI_DU_SECRET_KEY"`
	BaiDuSalt      string `env:"BAI_DU_SALT"`

	AdminAccounts []int64 `env:"ADMIN_ACCOUNTS"`

	BotAccounts  []int64  `env:"BOT_ACCOUNTS"`
	BotEndpoints []string `env:"BOT_ENDPOINTS"`

	DouBaoAccessKey string `env:"DOU_BAO_ACCESS_KEY"`
	DouBaoSecretKey string `env:"DOU_BAO_SECRET_KEY"`
}{}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Can not read env from file system, please check the right this program owned.")
	}

	//EnvCfg = envConfig{}

	if err := env.Parse(&EnvCfg); err != nil {
		panic("Can not parse env from file system, please check the env.")
	}
	if len(EnvCfg.BotAccounts) == 0 {
		panic("BotAccounts cannot be empty!")
	}
	if len(EnvCfg.BotEndpoints) == 0 {
		panic("BotEndpoints cannot be empty!")
	}
	if len(EnvCfg.BotAccounts) != len(EnvCfg.BotEndpoints) {
		panic("BotAccounts cannot match BotEndpoints!")
	}
}
