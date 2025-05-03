package socket_server

import "github.com/spf13/viper"

type SocketServerConfig struct {
	Port int
}

func LoadSocketServerConfig() (*SocketServerConfig, error) {
	viper.SetConfigName("server_config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	return &SocketServerConfig{
		Port: viper.GetInt("tcp_server.port"),
	}, nil
}
