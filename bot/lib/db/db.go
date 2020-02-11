package db

import (
	"discord-bot/lib/dotenv"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"os"
)

func ConnectDb() *gorm.DB {
	dotenv.EnvLoad()

	db, err := gorm.Open("sqlite3", "bot.db")
	if err != nil {
		panic("データベースへの接続に失敗しました")
	}

	var logMode = os.Getenv("SQL_DEBUG")
	var isLogged = false
	if logMode == "true" {
		isLogged = true
	}
	db.LogMode(isLogged)

	// スキーマのマイグレーション
	db.AutoMigrate(&User{})
	db.AutoMigrate(&UserJoin{})

	return db
}
