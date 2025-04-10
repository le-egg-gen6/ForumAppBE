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
	viper.SetConfigName("server_config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	return &LoggerConfig{
		LogLevel:    viper.GetString("log.log_level"),
		FilePattern: viper.GetString("log.file_pattern"),
		MaxSize:     viper.GetInt("log.max_size"),
		BaseLogDir:  viper.GetString("log.base_log_dir"),
	}, nil
}
