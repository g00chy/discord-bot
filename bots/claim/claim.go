package claim

import (
	"discord-bot/lib/discord"
	"discord-bot/lib/dotenv"
	"fmt"
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
)

// Main Main処理
func Main() {
	dotenv.EnvLoad()
	discord.AddHandler(onMessageCreate)
}

func onMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if discord.IsOwnMessage(s, m) {
		return
	}
	channel, err := s.State.Channel(m.ChannelID)
	if err != nil {
		discord.SendMessage(s, channel, "エラーだよ")
	}

	cCategory := os.Getenv("ADMIN_CATEGORY")
	cChannel := os.Getenv("ADMIN_CLAIM_CHANNEL")
	claimChannel, err := discord.GetFixChannel(s, cCategory, cChannel)
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
