package meem

import (
	"discord-bot/lib/db"
	"discord-bot/lib/discord"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
	"strings"
)

type methodType struct {
	methodType string
}

const methodTypeAdd = "add "
const methodTypeGet = "get "

var botPrefix string

var (
	connection = db.ConnectDb()
)

func Main() {
	discord.AddHandlerMeem(onAddMeemMessageCreate, onMeemGetMessageCreate)
	botPrefix = os.Getenv("MEEM_PREFIX") + " "
}

func getMessageKeyword(m *discordgo.MessageCreate, t methodType) (string, error) {

	if strings.HasPrefix(m.Message.Content, botPrefix+t.methodType) {
		keyword := strings.TrimSpace(m.Message.Content)
		keyword = strings.Replace(keyword, botPrefix+t.methodType, "", 1)
		keyword = strings.TrimSpace(keyword)
		if len(keyword) == 0 {
			return "", fmt.Errorf("not found keywrod")
		}
		return keyword, nil
	}
	return "", fmt.Errorf("err %s", "not found keyword")
}

func isMeem(m *discordgo.MessageCreate, t methodType) bool {
	if strings.HasPrefix(m.Message.Content, botPrefix+t.methodType) {
		return true
	}
	return false
}

func onAddMeemMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if !isMeem(m, methodType{methodTypeAdd}) {
		return
	}
	if len(m.Attachments) == 0 {
		return
	}

	c, err := s.State.Channel(m.ChannelID)
	if err != nil {
		log.Fatal("チャンネルを特定できませんでした")
		return
	}
	addMeem(s, m, c)
}
func onMeemGetMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if !isMeem(m, methodType{methodTypeGet}) {
		return
	}
	c, err := s.State.Channel(m.ChannelID)
	if err != nil {
		log.Fatal("チャンネルを特定できませんでした")
		return
	}

	getMeem(s, m, c)
}

//addMeem: meemを追加
func addMeem(s *discordgo.Session, m *discordgo.MessageCreate, c *discordgo.Channel) {
	if discord.IsOwnMessage(s, m) {
		return
	}

	keyword, err := getMessageKeyword(m, methodType{methodTypeAdd})
	if err != nil {
		discord.SendMessage(s, c, "[add]キーワードが指定されてません。")
		return
	}

	url, err := discord.GetImage(m)
	if err != nil {
		discord.SendMessage(s, c, "画像URLを取得できませんでした")
		return
	}
	var meem db.Meem
	connection.Where(db.Meem{ServerID: m.GuildID, UserID: m.Author.ID, Keyword: keyword}).Assign(db.Meem{Url: url, ChannelID: m.ChannelID}).FirstOrCreate(&meem)

	discord.SendMessage(s, c, "登録しました")
}

//getMeem: meemを取得
func getMeem(s *discordgo.Session, m *discordgo.MessageCreate, c *discordgo.Channel) {
	if discord.IsOwnMessage(s, m) {
		return
	}
	keyword, err := getMessageKeyword(m, methodType{methodTypeGet})
	if err != nil {
		discord.SendMessage(s, c, "[get]キーワードが指定されてません。")
		return
	}

	var meem []*db.Meem
	connection.Select("Url, Keyword").Where("keyword like ?", "%"+keyword+"%").First(&meem)

	if len(meem) < 1 {
		return
	}
	discord.SendMessage(s, c, meem[0].Keyword+" "+meem[0].Url)
}
