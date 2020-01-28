package afk

import (
	"discord-bot/lib/db"
	"discord-bot/lib/discord"
	"discord-bot/lib/dotenv"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/jinzhu/gorm"
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
	if isValidRequest(m.Author, connection) {
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
	afkCircleMinuteStr := os.Getenv("AFK_CIRCLE_MINUTES")
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
