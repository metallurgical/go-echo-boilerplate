package config

import "os"

type App interface {}

type AppConfig struct {
	AppName string
}

func AppNew() App {
	return AppConfig{
		AppName : os.Getenv("APP_NAME"),
	}
}