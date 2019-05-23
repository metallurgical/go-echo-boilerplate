package mail

import (
	"fmt"
	"github.com/metallurgical/go-echo-boilerplate/config"
	"net/smtp"
)

// Sending email when request to reset password.
func SendResetPasswordMail(email string) error {
	mailConfig := config.MailNew().(*config.MailConfig)
	appConfig := config.AppNew().(*config.AppConfig)
	from := mailConfig.MailFrom
	pass := mailConfig.MailPassword
	to := email

	msg := fmt.Sprintf(
		"From: %s \n"+"To: %s \n"+"Subject: [%s]:%s \n\n"+"%s",
		from,
		to,
		appConfig.AppName,
		"Request to Reset Password",
		"Hello there...",
	)

	err := smtp.SendMail(fmt.Sprintf("%s:%s", mailConfig.MailHost, mailConfig.MailPort),
		smtp.PlainAuth("", mailConfig.MailUsername, pass, mailConfig.MailHost),
		from, []string{to}, []byte(msg))

	if err != nil {
		return err
	}

	return nil
}
