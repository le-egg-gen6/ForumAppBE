package mail_sender

import (
	"fmt"
	"net/mail"
	"net/smtp"
	"regexp"
	"strconv"
)

type MailSender struct {
	Config *MailSenderConfig
}

var Instance *MailSender

func InitializeMailSender() {
	cfg, err := LoadMailSenderConfig()
	if err != nil {
		panic("Mail sender configuration not found")
	}
	Instance = &MailSender{
		Config: cfg,
	}
}

func GetMailSenderInstance() *MailSender {
	return Instance
}

const MailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

func (ms *MailSender) ValidateMail(email string) bool {
	_, err := mail.ParseAddress(email)
	if err != nil {
		return false
	}
	matched, _ := regexp.MatchString(MailRegex, email)
	return matched
}

func (ms *MailSender) SendMail(to string, subject string, body string) error {
	auth := smtp.PlainAuth("", ms.Config.SMTPUsername, ms.Config.SMTPPassword, ms.Config.SMTPHost)

	message := []byte(
		"From: " + ms.Config.SMTPUsername + "\r\n" +
			"To: " + to + "\r\n" +
			"Subject: " + subject + "\r\n\r\n" +
			body + "\r\n")

	addr := fmt.Sprintf("%s:%d", ms.Config.SMTPHost, ms.Config.SMTPPort)

	err := smtp.SendMail(addr, auth, ms.Config.SMTPUsername, []string{to}, message)
	if err != nil {
		return err
	}
	return nil
}

func SendValidateMail(to string, username string, code uint64) error {
	body := fmt.Sprintf(Instance.Config.ValidateMailPattern, username, strconv.FormatUint(code, 10))
	err := Instance.SendMail(to, "Validate your account", body)
	if err != nil {
		return err
	}
	return nil
}
