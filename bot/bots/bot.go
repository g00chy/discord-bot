package main

import (
	"discord-bot/bots/afk"
	"discord-bot/bots/claim"
	"discord-bot/bots/logging"
	"discord-bot/bots/meem"
	"discord-bot/lib/discord"
	"discord-bot/lib/dotenv"
	"os"
)

func main() {
	dotenv.EnvLoad()
	token := os.Getenv("CLAIM_BOT_TOKEN")
	claim.Main()
	afk.Main()
	//nleaveban.Main()
	logging.Main()
	meem.Main()

	_ = discord.SetUpDiscordBot(token)
}
