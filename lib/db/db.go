package db

import (
	"discord-bot/lib/dotenv"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

func ConnectDb() *gorm.DB {
	dotenv.EnvLoad()

	db, err := gorm.Open("sqlite3", "bot.db")
	if err != nil {
		panic("データベースへの接続に失敗しました")
	}

	// スキーマのマイグレーション
	db.AutoMigrate(&User{})
	db.AutoMigrate(&UserJoin{})

	return db
}
