package logging

import (
	"discord-bot/lib/discord"
	"discord-bot/lib/dotenv"
	"log"
	"time"

	"github.com/bwmarrin/discordgo"
)

// Main main
func Main() {
	dotenv.EnvLoad()
	discord.AddHandler(onMessageCreate)
}

func onMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if discord.IsOwnMessage(s, m) {
		return
	}

	c, err := s.State.Channel(m.ChannelID)

	if err != nil {
		log.Println("Error getting channel: ", err)
		return
	}

	var channel = discord.GetChannel(s, m)
	log.Printf("%20s %20s %20s %20s %20s > %s\n", channel.ParentID, c.Name, m.ChannelID, time.Now().Format(time.Stamp), m.Author.Username, m.Content)
}
