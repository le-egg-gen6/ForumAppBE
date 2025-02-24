package mail_sender

import "github.com/spf13/viper"

type MailSenderConfig struct {
	SMTPHost     string
	SMTPPort     int
	SMTPUsername string
	SMTPPassword string
}

func LoadMailSenderConfig() (*MailSenderConfig, error) {
	viper.SetConfigName("mail_config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	return &MailSenderConfig{
		SMTPHost:     viper.GetString("smtp_host"),
		SMTPPort:     viper.GetInt("smtp_port"),
		SMTPUsername: viper.GetString("smtp_username"),
		SMTPPassword: viper.GetString("smtp_password"),
	}, nil
}
