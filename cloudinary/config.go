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
	viper.SetConfigName("cloudinary_config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	return &CloudinaryConfig{
		CloudName:    viper.GetString("cloud_name"),
		APIKey:       viper.GetString("api_key"),
		APISecret:    viper.GetString("api_secret"),
		UploadFolder: viper.GetString("upload_folder"),
	}, nil
}
