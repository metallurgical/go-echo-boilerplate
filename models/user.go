package models

import "time"

type User struct {
	ID        uint   `gorm:"primary_key"`
	Email     string `gorm:"type:varchar(100);unique_index"`
	Name      string
	Password  string `gorm:"-"` // ignore this field
	CreatedAt time.Time
	UpdatedAt time.Time
}
