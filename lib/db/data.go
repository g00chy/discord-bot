package db

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	UserId   string `gorm:"not null"`
	UserName string `gorm:"not null"`
}

type UserJoin struct {
	gorm.Model
	UserId string `gorm: "not null"`
	Count  int    `gorm:"not null"`
}
