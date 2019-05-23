package auth

import (
	"github.com/jinzhu/gorm"
	"github.com/metallurgical/go-echo-boilerplate/models"
)

var (
	db    *gorm.DB
	count int
	user  models.User
)