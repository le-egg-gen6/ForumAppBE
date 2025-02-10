package logger

import (
	"github.com/spf13/viper"
)

type LoggerConfig struct {
	LogLevel    string
	FilePattern string
	MaxSize     int
	BaseLogDir  string
}

func LoadLoggerConfig() (*LoggerConfig, error) {
	viper.SetConfigName("log_config")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	return &LoggerConfig{
		LogLevel:    viper.GetString("log_level"),
		FilePattern: viper.GetString("file_pattern"),
		MaxSize:     viper.GetInt("max_size"),
		BaseLogDir:  viper.GetString("base_log_dir"),
	}, nil
}
