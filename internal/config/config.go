package config

import (
	"shokHorizon/go_recipes/pkg/logger"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	IsDebug bool   `yaml:"is_debug" env-default:"false"`
	Port    string `yaml:"port"     env-default:"8000"`
	BindIp  string `yaml:"bind_ip"  env-default:"127.0.0.1"`
}

var instance *Config

func GetConfig() *Config {
	if instance == nil {
		l := logger.GetLogger()
		instance = &Config{}
		if err := cleanenv.ReadConfig("config.yaml", instance); err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			l.Err(help)
			l.Fatal(err)
		}
	}
	return instance
}
