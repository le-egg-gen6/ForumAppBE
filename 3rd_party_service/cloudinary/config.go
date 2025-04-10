package cloudinary

import (
	"github.com/spf13/viper"
)

type CloudinaryConfig struct {
	CloudName    string
	APIKey       string
	APISecret    string
	UploadFolder string
}

func LoadCloudinaryConfig() (*CloudinaryConfig, error) {
	viper.SetConfigName("server_config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	return &CloudinaryConfig{
		CloudName:    viper.GetString("cloudinary.cloud_name"),
		APIKey:       viper.GetString("cloudinary.api_key"),
		APISecret:    viper.GetString("cloudinary.api_secret"),
		UploadFolder: viper.GetString("cloudinary.upload_folder"),
	}, nil
}
