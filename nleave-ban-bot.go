package main

import (
	"discord-bot/lib/discord"
	"discord-bot/lib/dotenv"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"os"
	"time"
)

func main() {
	dotenv.EnvLoad()
	token := os.Getenv("CLAIM_BOT_TOKEN")
	_ = discord.StartDiscordBot(onLeaveMessageCreate, token)
}

func onLeaveMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if discord.IsOwnMessage(s, m) {
		return
	}
	channel, err := s.State.Channel(m.ChannelID)
	if err != nil {
		discord.SendMessage(s, channel, "エラーだよ")
	}

	fmt.Printf("%20s %20s %20s %20s %20s > %s\n", channel.ParentID, channel.Name, m.ChannelID,
		time.Now().Format(time.Stamp), m.Author.Username, m.Content)
}
