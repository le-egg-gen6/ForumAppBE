package mail_sender

import (
	"github.com/spf13/viper"
	"os"
)

type MailSenderConfig struct {
	SMTPHost            string
	SMTPPort            int
	SMTPUsername        string
	SMTPPassword        string
	ValidateMailPattern string
}

func LoadMailSenderConfig() (*MailSenderConfig, error) {
	viper.SetConfigName("server_config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	validateMailPatternDir := viper.GetString("mail.validate_mail_pattern_dir")

	validateMailPattern, err := os.ReadFile(validateMailPatternDir)
	if err != nil {
		return nil, err
	}

	return &MailSenderConfig{
		SMTPHost:            viper.GetString("mail.smtp_host"),
		SMTPPort:            viper.GetInt("mail.smtp_port"),
		SMTPUsername:        viper.GetString("mail.smtp_username"),
		SMTPPassword:        viper.GetString("mail.smtp_password"),
		ValidateMailPattern: string(validateMailPattern),
	}, nil
}
