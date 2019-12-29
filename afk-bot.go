package main

import (
	"discord-bot/lib/db"
	"discord-bot/lib/discord"
	"discord-bot/lib/dotenv"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/jinzhu/gorm"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	dotenv.EnvLoad()
	token := os.Getenv("AFK_BOT_TOKEN")
	_ = discord.StartDiscordBot(onAfkMessageCreate, token)
}

func onAfkMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if discord.IsOwnMessage(s, m) {
		return
	}

	c, err := s.State.Channel(m.ChannelID)

	if err != nil {
		log.Println("Error getting channel: ", err)
		return
	}

	var channel = discord.GetChannel(s, m)
	fmt.Printf("%20s %20s %20s %20s %20s > %s\n", channel.ParentID, c.Name, m.ChannelID, time.Now().Format(time.Stamp), m.Author.Username, m.Content)

	if strings.HasPrefix(m.Content, fmt.Sprintf("%s", "!afk")) {
		afk(s, m)
	}
}

func afk(s *discordgo.Session, m *discordgo.MessageCreate) {
	connection := db.ConnectDb()

	if !isExistMentions(m.Mentions) {
		discord.SendMessage(s, discord.GetChannel(s, m), "コマンドが違うぞ。")
		return
	}
	var member = discord.GetMemberId(m.Mentions)
	fmt.Printf("member %s", member.ID)
	if isValidRequest(member, connection) {
		goAfk(s, m, member)
	} else {
		discord.SendMessage(s, discord.GetChannel(s, m), os.Getenv("ERROR_IMG"))
	}

}

func isExistMentions(u []*discordgo.User) bool {
	if len(u) == 0 || len(u) > 1 {
		return false
	}
	return true
}

func isValidRequest(s *discordgo.User, connection *gorm.DB) bool {
	var user []*db.User
	t := time.Now()
	afkCircleMinuteStr := os.Getenv("AFK_CICLE_MINUTES")
	afkCircleMinute, _ := strconv.Atoi(afkCircleMinuteStr)
	afkCircleTime := t.Add(time.Duration(-1*afkCircleMinute) * time.Minute)

	connection.Where("user_id = ? AND created_at >= ?", s.ID, afkCircleTime).Find(&user)
	countStr := os.Getenv("AFK_MAX_COUNT")
	count, _ := strconv.Atoi(countStr)
	fmt.Printf("count: %d", count)
	if len(user) >= count {
		return false
	}
	createRequestRecord(connection, s)

	var deleteUser db.User
	connection.Where("user_id = ? AND created_at < ?", s.ID, afkCircleTime).Unscoped().Delete(&deleteUser)
	return true
}

func createRequestRecord(connection *gorm.DB, user *discordgo.User) {
	fmt.Printf("%s, %s", user.ID, user.Username)
	connection.Create(&db.User{UserId: user.ID, UserName: user.Username})
}

func goAfk(s *discordgo.Session, m *discordgo.MessageCreate, member *discordgo.User) {
	guild, _ := s.Guild(m.GuildID)
	afkChannelID := guild.AfkChannelID
	_ = s.GuildMemberMove(guild.ID, member.ID, afkChannelID)
}
