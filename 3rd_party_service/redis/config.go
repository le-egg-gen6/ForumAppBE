package redis

import (
	"github.com/spf13/viper"
)

type RedisConfig struct {
	RedisHost     string
	RedisPort     int
	RedisPassword string
	RedisDatabase int
}

func LoadRedisConfig() (*RedisConfig, error) {
	viper.SetConfigName("server_config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}
	return &RedisConfig{
		RedisHost:     viper.GetString("redis.host"),
		RedisPort:     viper.GetInt("redis.port"),
		RedisPassword: viper.GetString("redis.password"),
		RedisDatabase: viper.GetInt("redis.database"),
	}, nil
}
