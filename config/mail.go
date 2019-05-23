package config

import "os"

type Mail interface{}

type MailConfig struct {
	MailDriver     string
	MailHost       string
	MailPort       string
	MailUsername   string
	MailPassword   string
	MailEncryption string
	MailFrom       string
}

func MailNew() Mail {
	return &MailConfig{
		MailDriver:     os.Getenv("MAIL_DRIVER"),
		MailHost:       os.Getenv("MAIL_HOST"),
		MailPort:       os.Getenv("MAIL_PORT"),
		MailUsername:   os.Getenv("MAIL_USERNAME"),
		MailPassword:   os.Getenv("MAIL_PASSWORD"),
		MailEncryption: os.Getenv("MAIL_ENCRYPTION"),
		MailFrom:       os.Getenv("MAIL_FORM"),
	}
}
