package handler

import "discord-bot/lib/db"

var (
	connection = db.ConnectDb()
)
