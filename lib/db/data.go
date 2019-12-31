package db

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	UserId   string `gorm:"not null"`
	UserName string `gorm:"not null"`
}

type UserJoin struct {
	gorm.Model
	UserId     string `gorm: "unique;not null"`
	UserName   string `gorm:"not null"`
	JoinCount  int    `gorm:"not null" sql:"DEFAULT:0"`
	LeaveCount int    `gorm:"not null" sql:"DEFAULT:0"`
}
