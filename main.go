package main

import (
	afk_bot "discord-bot/afk-bot"
	claim_bot "discord-bot/claim-bot"
	. "discord-bot/lib/discord"
	"discord-bot/lib/dotenv"
	nleave_ban_bot "discord-bot/nleave-ban-bot"
	"os"
)

func main() {
	dotenv.EnvLoad()
	token := os.Getenv("CLAIM_BOT_TOKEN")
	claim_bot.Main()
	afk_bot.Main()
	nleave_ban_bot.Main()

	SetUpDiscordBot(token)
}
