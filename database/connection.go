package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/metallurgical/go-echo-boilerplate/config"
	"log"
)

type DatabaseProvider interface {}

type DatabaseProviderConnection struct {
	Db *gorm.DB
}

func ConnectMYSQL() DatabaseProvider {
	databaseConfig := config.DatabaseNew().(*config.DatabaseConfig);
	// Connecting to MYSQL database
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		databaseConfig.Mysql.DbUsername,
		databaseConfig.Mysql.DbPassword,
		databaseConfig.Mysql.DbHost,
		databaseConfig.Mysql.DbPort,
		databaseConfig.Mysql.DbDatabase,
	))
	if err != nil {
		log.Fatalf("Could not connect to database :%v", err)
	}

	return &DatabaseProviderConnection{
		Db: db,
	}
}