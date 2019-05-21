package config

import "os"

type App interface {}

type app struct {
	AppName string
}

func AppNew() App {
	return app{
		AppName : os.Getenv("APP_NAME"),
	}
}