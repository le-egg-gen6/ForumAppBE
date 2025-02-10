package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	API_VERSION string
	PORT        string
	DB_USERNAME string
	DB_PASSWORD string
	DB_NAME     string
	DB_HOST     string
	DB_PORT     string
}

func LoadConfig() *Config {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		panic("Error reading config file" + err.Error())
	}

	return &Config{
		API_VERSION: viper.GetString("API_VERSION"),
		PORT:        viper.GetString("PORT"),
		DB_USERNAME: viper.GetString("DB_USERNAME"),
		DB_PASSWORD: viper.GetString("DB_PASSWORD"),
		DB_NAME:     viper.GetString("DB_NAME"),
		DB_HOST:     viper.GetString("DB_HOST"),
		DB_PORT:     viper.GetString("DB_PORT"),
	}
}
