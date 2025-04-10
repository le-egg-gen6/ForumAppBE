package http

import "github.com/spf13/viper"

type HTTPServerConfig struct {
	APIVersion         string
	Port               int
	ReadTimeoutSec     int
	WriteTimeoutSec    int
	ShutdownTimeoutSec int
}

func LoadHTTPServerConfig() (*HTTPServerConfig, error) {
	viper.SetConfigName("server_config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	return &HTTPServerConfig{
		APIVersion:         viper.GetString("http_server.api_version"),
		Port:               viper.GetInt("http_server.port"),
		ReadTimeoutSec:     viper.GetInt("http_server.read_timeout_sec"),
		WriteTimeoutSec:    viper.GetInt("http_server.write_timeout_sec"),
		ShutdownTimeoutSec: viper.GetInt("http_server.shutdown_timeout_sec"),
	}, nil
}
