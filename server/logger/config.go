package logger

import (
	"github.com/spf13/viper"
)

type LoggerConfig struct {
	LOG_LEVEL    string
	FILE_PATTERN string
	MAX_SIZE     int
	BASE_LOG_DIR string
}

func LoadLoggerConfig() (*LoggerConfig, error) {
	viper.SetConfigName("log_config")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	return &LoggerConfig{
		LOG_LEVEL:    viper.GetString("log_level"),
		FILE_PATTERN: viper.GetString("file_pattern"),
		MAX_SIZE:     viper.GetInt("max_size"),
		BASE_LOG_DIR: viper.GetString("base_log_dir"),
	}, nil
}
