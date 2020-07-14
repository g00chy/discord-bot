package nleaveban

import (
	"discord-bot/lib/db"
	"discord-bot/lib/discord"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/jinzhu/gorm"
	"log"
	"os"
	"strconv"
)

const eventTypeJoin = 1 // joinイベント

type eventType struct {
	eventType int
	user      discordgo.User
	GuildID   string
}

var (
	session                *discordgo.Session
	leaveMaxCount          int
	announceChannelDiscord *discordgo.Channel
	connection             = db.ConnectDb()
)

// Main main
func Main() {

	countStr := os.Getenv("LEAVE_MAX_COUNT")
	leaveMaxCount, _ = strconv.Atoi(countStr)
	log.Printf("count: %d", leaveMaxCount)
	discord.AddHandlerJoinAndLeave(onJoinMessageCreate, onLeaveMessageCreate)
}

func onJoinMessageCreate(s *discordgo.Session, m *discordgo.GuildMemberAdd) {
	e := eventType{1, *m.User, m.GuildID}
	discord.SetSendMessageChannel(s)
	count(e)
}
func onLeaveMessageCreate(s *discordgo.Session, m *discordgo.GuildMemberRemove) {
	e := eventType{2, *m.User, m.GuildID}
	discord.SetSendMessageChannel(s)
	count(e)
}

func count(event eventType) {

	if event.eventType == 1 {
		log.Printf("Join: ID:%s NAME:%s\r\n", event.user.ID, event.user.Username)
	} else {
		log.Printf("Remove: ID:%s NAME:%s\r\n", event.user.ID, event.user.Username)
	}

	var userJoin []*db.UserJoin
	connection.Where("user_id = ?", event.user.ID).First(&userJoin)

	if len(userJoin) == 0 {
		if event.eventType == eventTypeJoin {
			message := fmt.Sprintf("はじめまして。 %s 本サーバーは3回抜けるとBANになります。ご注意ください。",
				event.user.Mention())
			discord.SendMessage(session, discord.AnnounceChannelDiscord, message)
		}
		createUserJoin(event, connection)
		return
	}
	if event.eventType == eventTypeJoin {
		message := fmt.Sprintf("%s 今までのサーバー離脱回数:%d あと%d回サーバーから抜けるとBANになります。",
			event.user.Mention(),
			userJoin[0].LeaveCount, leaveMaxCount-userJoin[0].LeaveCount)
		discord.SendMessage(session, announceChannelDiscord, message)
		joinCount := userJoin[0].JoinCount
		connection.Model(userJoin).Update(db.UserJoin{JoinCount: joinCount + 1})
	} else {
		leaveCount := userJoin[0].JoinCount
		connection.Model(userJoin).Update(db.UserJoin{LeaveCount: leaveCount + 1})
		userJoin[0].LeaveCount = leaveCount + 1
	}

	if userJoin[0].LeaveCount >= leaveMaxCount {
		userBan(event)
	}
}

func createUserJoin(e eventType, connection *gorm.DB) {
	if e.eventType == eventTypeJoin {
		connection.Create(&db.UserJoin{UserID: e.user.ID, UserName: e.user.Username, JoinCount: 1})
	} else {
		connection.Create(&db.UserJoin{UserID: e.user.ID, UserName: e.user.Username, LeaveCount: 1})
	}
}

func userBan(e eventType) {
	error := session.GuildBanCreateWithReason(e.GuildID, e.user.ID, "サーバー上限離脱回数を超えました。", 0)
	if error != nil {
		discord.SendMessage(session, announceChannelDiscord, fmt.Sprintf("error %s", error))
	}
	discord.SendMessage(session, announceChannelDiscord,
		fmt.Sprintf("%s は　サーバー上限離脱回数%dを超えたため、BANとなりました。", e.user.Username, leaveMaxCount))
}
