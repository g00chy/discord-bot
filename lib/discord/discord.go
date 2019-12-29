package discord

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"kokoro-bot/lib/dotenv"
	"log"
	"os"
)

var (
	Token   = "Bot"
	stopBot = make(chan bool)
)

func SendMessage(s *discordgo.Session, c *discordgo.Channel, msg string) {
	_, err := s.ChannelMessageSend(c.ID, msg)

	log.Println(">>> " + msg)
	if err != nil {
		log.Println("Error sending message: ", err)
	}
}

func GetChannel(s *discordgo.Session, m *discordgo.MessageCreate) *discordgo.Channel {
	c, err := s.State.Channel(m.ChannelID) //チャンネル取得
	if err != nil {
		log.Println("Error getting channel: ", err)
		return nil
	}
	return c
}

func StartDiscordBot(onMessageCreate func(s *discordgo.Session, m *discordgo.MessageCreate)) error {
	dotenv.EnvLoad()

	Token = Token + " " + os.Getenv("TOKEN")
	discord, err := discordgo.New()
	if err != nil {
		fmt.Println("Error logging in")
		fmt.Println(err)
		return err
	}
	discord.Token = Token

	discord.AddHandler(onMessageCreate)
	err = discord.Open()
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println("Listening...")
	<-stopBot
	return nil
}

func GetMemberId(u []*discordgo.User) *discordgo.User {
	return u[len(u)-1]
}
