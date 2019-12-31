package db

import (
	"discord-bot/lib/dotenv"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"os"
)

func ConnectDb() *gorm.DB {
	dotenv.EnvLoad()
	var (
		dbHostName = os.Getenv("POSTGRES_HOST")
		dbPort     = os.Getenv("POSTGRES_PORT")
		dbName     = os.Getenv("POSTGRES_DB")
		dbUser     = os.Getenv("POSTGRES_USER")
		dbPassword = os.Getenv("POSTGRES_PASSWORD")
	)
	var dbAccessString = "host=" + dbHostName + " port=" + dbPort +
		" user=" + dbUser + " dbname=" + dbName + " password=" + dbPassword + " sslmode=disable"
	fmt.Printf("%20s", dbAccessString)
	db, err := gorm.Open("postgres", dbAccessString)

	if err != nil {
		panic("データベースへの接続に失敗しました")
	}

	// スキーマのマイグレーション
	db.AutoMigrate(&User{})
	db.AutoMigrate(&UserJoin{})

	return db
}
