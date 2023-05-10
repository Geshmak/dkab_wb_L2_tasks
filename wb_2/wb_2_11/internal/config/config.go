package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	HTTP
}

type HTTP struct {
	Port string
}

func NewConfig() (*Config, error) {
	cfg := &Config{}
	viper.AddConfigPath("/Program Files/Go/src/wb_2_11")
	viper.SetConfigName("config")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln(err)
	}
	cfg.Port = viper.GetString("port")
	return cfg, nil
}
