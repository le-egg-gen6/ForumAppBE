package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	API_VERSION    string
	PORT           int
	TCP_PORT       int
	UDP_PORT       int
	DB_USERNAME    string
	DB_PASSWORD    string
	DB_NAME        string
	DB_HOST        string
	DB_PORT        int
	REDIS_HOST     string
	REDIS_PORT     int
	REDIS_PASSWORD string
	REDIS_DATABASE int
}

func LoadConfig() *Config {
	viper.SetConfigName("server_config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		panic("Error reading config file" + err.Error())
	}

	return &Config{
		API_VERSION:    viper.GetString("server.api_version"),
		PORT:           viper.GetInt("server.port"),
		TCP_PORT:       viper.GetInt("server.tcp_port"),
		UDP_PORT:       viper.GetInt("server.udp_port"),
		DB_USERNAME:    viper.GetString("database.username"),
		DB_PASSWORD:    viper.GetString("database.password"),
		DB_NAME:        viper.GetString("database.name"),
		DB_HOST:        viper.GetString("database.host"),
		DB_PORT:        viper.GetInt("database.port"),
		REDIS_HOST:     viper.GetString("redis.host"),
		REDIS_PORT:     viper.GetInt("redis.port"),
		REDIS_PASSWORD: viper.GetString("redis.password"),
		REDIS_DATABASE: viper.GetInt("redis.db"),
	}
}
