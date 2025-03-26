package server_http

import "github.com/spf13/viper"

type HTTPServerConfig struct {
	APIVersion string
	Port       int
}

func LoadHTTPServerConfig() (*HTTPServerConfig, error) {
	viper.SetConfigName("server_config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	return &HTTPServerConfig{
		APIVersion: viper.GetString("http_server.api_version"),
		Port:       viper.GetInt("http_server.port"),
	}, nil
}
