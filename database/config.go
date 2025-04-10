package database

import (
	"github.com/spf13/viper"
)

type DatabaseConfig struct {
	DBUsername string
	DBPassword string
	DBName     string
	DBHost     string
	DBPort     int
}

func LoadDatabaseConfig() (*DatabaseConfig, error) {
	viper.SetConfigName("server_config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	return &DatabaseConfig{
		DBUsername: viper.GetString("database.username"),
		DBPassword: viper.GetString("database.password"),
		DBName:     viper.GetString("database.name"),
		DBHost:     viper.GetString("database.host"),
		DBPort:     viper.GetInt("database.port"),
	}, nil
}
