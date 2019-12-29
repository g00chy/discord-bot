package main

import (
	"discord-bot/lib/discord"
	"discord-bot/lib/dotenv"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
	"time"
)

func main() {
	dotenv.EnvLoad()
	token := os.Getenv("CLAIM_BOT_TOKEN")
	_ = discord.StartDiscordBot(onClaimMessageCreate, token)
}

func onClaimMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if discord.IsOwnMessage(s, m) {
		return
	}
	channel, err := s.State.Channel(m.ChannelID)

	fmt.Printf("%20s %20s %20s %20s %20s > %s\n", channel.ParentID, channel.Name, m.ChannelID, time.Now().Format(time.Stamp), m.Author.Username, m.Content)

	cCategory := os.Getenv("ADMIN_CATEGORY")
	cChannel := os.Getenv("ADMIN_CLAIM_CHANNEL")
	claimChannel, err := discord.GetFixChannel(s, m, cCategory, cChannel)
	if err != nil {
		discord.SendMessage(s, channel, "エラーだよ")
	}

	if m.Author.ID == s.State.User.ID {
		return
	}

	if err != nil {
		log.Println("Error getting channel: ", err)
		return
	}

	isDM, err := discord.ComesFromDM(s, m)
	if err != nil {
		log.Println("Error getting channel: ", err)
		return
	}

	if isDM {
		message := fmt.Sprintf("User: %s, UserId: %s, Avator: %s Message: %s", m.Author.Username, m.Author.ID, m.Author.AvatarURL("128"), m.Content)
		discord.SendMessage(s, claimChannel, message)
		discord.SendMessage(s, channel, "受け付けました。\r\n"+message)
	}
}
