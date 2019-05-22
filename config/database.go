package config

import "os"

type Database interface{}

type MysqlDbConnection struct {
	DbHost     string
	DbPort     string
	DbDatabase string
	DbUsername string
	DbPassword string
}

type DatabaseConfig struct {
	Mysql MysqlDbConnection
}

func DatabaseNew() Database {
	return &DatabaseConfig{
		Mysql: MysqlDbConnection{
			DbHost:     os.Getenv("DB_HOST"),
			DbPort:     os.Getenv("DB_PORT"),
			DbDatabase: os.Getenv("DB_DATABASE"),
			DbUsername: os.Getenv("DB_USERNAME"),
			DbPassword: os.Getenv("DB_PASSWORD"),
		},
	}
}
