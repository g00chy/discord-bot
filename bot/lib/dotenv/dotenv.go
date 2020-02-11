package dotenv

import (
	"github.com/joho/godotenv"
	"log"
)

func EnvLoad() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
