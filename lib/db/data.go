package db

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	UserId   string `gorm:"not null"`
	UserName string `gorm:"not null"`
}
