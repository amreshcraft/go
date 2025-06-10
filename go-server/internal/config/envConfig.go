package config

import (
	"go-server/internal/logger"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var AppConfig *Config

type Config struct {
	Port string
}

func LoadConfig() {
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		logger.Log.Error("Error loading config file", zap.Error(err))
	}

	AppConfig = &Config{
		Port: viper.GetString("PORT"),
	}
}
