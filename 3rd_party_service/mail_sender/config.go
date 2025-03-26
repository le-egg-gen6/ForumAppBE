package mail_sender

import "github.com/spf13/viper"

type MailSenderConfig struct {
	SMTPHost     string
	SMTPPort     int
	SMTPUsername string
	SMTPPassword string
}

func LoadMailSenderConfig() (*MailSenderConfig, error) {
	viper.SetConfigName("server_config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	return &MailSenderConfig{
		SMTPHost:     viper.GetString("mail.smtp_host"),
		SMTPPort:     viper.GetInt("mail.smtp_port"),
		SMTPUsername: viper.GetString("mail.smtp_username"),
		SMTPPassword: viper.GetString("mail.smtp_password"),
	}, nil
}
