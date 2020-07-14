package meem

import (
	"discord-bot/lib/db"
	"discord-bot/lib/discord"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
	"strconv"
	"strings"
)

type methodType struct {
	methodType string
}

const methodTypeAdd = "add "
const methodTypeList = "list"
const methodTypeDelete = "delete "
const meemPerPage = 10

var botPrefix string

var (
	connection = db.ConnectDb()
)

func Main() {
	discord.AddHandlerMeem(onAddMeemMessageCreate, onGetMeemMessageCreate, onGetListMessageCreate, onGetListMessageDelete)
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

func getListPage(m *discordgo.MessageCreate) (int, error) {
	if strings.HasPrefix(m.Message.Content, botPrefix+methodTypeList) {
		page := strings.TrimSpace(m.Message.Content)
		page = strings.Replace(page, botPrefix+methodTypeList, "", 1)
		pageNum, err := strconv.Atoi(strings.TrimSpace(page))
		if err != nil {
			return 1, nil
		}
		if pageNum < 1 {
			return 1, nil
		}
		return pageNum, nil
	}
	return 1, fmt.Errorf("err %s", "not found page")
}

func getDeleteId(m *discordgo.MessageCreate) (int, error) {
	if strings.HasPrefix(m.Message.Content, botPrefix+methodTypeDelete) {
		id := strings.TrimSpace(m.Message.Content)
		id = strings.Replace(id, botPrefix+methodTypeDelete, "", 1)
		idNum, err := strconv.Atoi(strings.TrimSpace(id))
		if err != nil {
			return 0, nil
		}
		return idNum, nil
	}
	return 0, fmt.Errorf("err %s", "not found page")
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
func onGetMeemMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	c, err := s.State.Channel(m.ChannelID)
	if err != nil {
		log.Fatal("チャンネルを特定できませんでした")
		return
	}

	getMeem(s, m, c)
}
func onGetListMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if !isMeem(m, methodType{methodTypeList}) {
		return
	}
	c, err := s.State.Channel(m.ChannelID)
	if err != nil {
		log.Fatal("チャンネルを特定できませんでした")
		return
	}

	getList(s, m, c)
}
func onGetListMessageDelete(s *discordgo.Session, m *discordgo.MessageCreate) {
	if !isMeem(m, methodType{methodTypeDelete}) {
		return
	}
	c, err := s.State.Channel(m.ChannelID)
	if err != nil {
		log.Fatal("チャンネルを特定できませんでした")
		return
	}

	deleteMeem(s, m, c)
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

	var meem db.Meem
	connection.Select("Url, Keyword").Where("keyword like ?", "%"+strings.TrimSpace(m.Message.Content)+"%").Order("keyword asc").First(&meem)

	if meem.Keyword == "" {
		return
	}
	discord.SendMessage(s, c, meem.Keyword+" "+meem.Url)
}

//getDelete meemを削除

func deleteMeem(s *discordgo.Session, m *discordgo.MessageCreate, c *discordgo.Channel) {
	if discord.IsOwnMessage(s, m) {
		return
	}
	id, err := getDeleteId(m)
	if err != nil || id == 0 {
		discord.SendMessage(s, c, "[delete]idが指定されていません。")
		return
	}

	var meem db.Meem

	tx := connection.Model(&db.Meem{}).Where(db.Meem{UserID: m.Author.ID}).Where(id)
	var count int
	tx.Count(&count)
	tx.Find(&meem)
	if count > 0 {
		tx.Delete(meem)
	} else {
		discord.SendMessage(s, c, "[delete]削除対象のmeemはありません。")
	}

	embed := discord.NewEmbed().
		SetTitle("削除に成功しました").
		SetColor(0x00ff00).
		AddField(strconv.FormatUint(uint64(id), 10), meem.Url)
	discord.SendEmbedMessage(s, c, embed.MessageEmbed)
}

//getList: 登録しているmeem一覧の取得
func getList(s *discordgo.Session, m *discordgo.MessageCreate, c *discordgo.Channel) {
	if discord.IsOwnMessage(s, m) {
		return
	}
	page, err := getListPage(m)
	if err != nil {
		discord.SendMessage(s, c, "[list]キーワードが指定されてません。")
		return
	}

	var meems []*db.Meem
	var count int
	var totalPage int
	tx := connection.Model(&db.Meem{}).Where(db.Meem{UserID: m.Author.ID})
	tx.Count(&count)
	tx.Select("id, url, keyword").Offset((page - 1) * meemPerPage).Limit(meemPerPage).Find(&meems)
	log.Println(count, len(meems))
	if count > meemPerPage {
		totalPage = (count / meemPerPage) + 1
	} else {
		totalPage = 1
	}
	if len(meems) < 1 {
		discord.SendMessage(s, c, "[list]meemは登録されていないみたいよ")

		return
	}
	if totalPage < page {
		discord.SendMessage(s, c, "[list]ページ数が超過しています")
		return
	}

	embed := discord.NewEmbed().
		SetTitle("あなたの登録しているMeem").
		SetColor(0x00ff00)
	for _, meem := range meems {
		log.Print(meem.ID)
		embed.AddField(strconv.FormatUint(uint64(meem.ID), 10)+":"+meem.Keyword, meem.Url)
	}
	embed.SetFooter("page " + strconv.Itoa(page) + " / " + strconv.Itoa(totalPage))
	discord.SendEmbedMessage(s, c, embed.MessageEmbed)
}
