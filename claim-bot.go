package main

import (
	"github.com/bwmarrin/discordgo"
	"kokoro-bot/lib/discord"
	"log"
)

func main() {
	_ = discord.StartDiscordBot(onClaimMessageCreate)
}

func onClaimMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	_, err := s.State.Channel(m.ChannelID)

	if err != nil {
		log.Println("Error getting channel: ", err)
		return
	}
}
