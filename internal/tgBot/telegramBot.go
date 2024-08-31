package tgBot

import (
	"fmt"
	"time"

	"stalcraftBot/internal/getData"
	jSon "stalcraftBot/internal/jSon"
	"stalcraftBot/internal/timeRes"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const chatIDsFile = "/var/tmp/chat_ids.json"

// memory chats users
var ChatIDs []int64

// set your telegram bot token from @BotFather
var telegramToken string = "tgToken"

// make bot
var bot, _ = tgbotapi.NewBotAPI(telegramToken)

func SendMessageTG(s string) {

	// print bot info and send message
	botUser, err := bot.GetMe()
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Printf("\nBot user: %v\n", botUser)
	// read id from json file to ChatIDs slice
	LoadChatID()

	// Send message
	for _, id := range ChatIDs {
		msg := tgbotapi.NewMessage(id, s)
		bot.Send(msg)
	}
}

func BotChating() {
	// update chats and save id to json file
	url := "https://eapi.stalcraft.net/ru/emission"
	token := "stalcraftToken"
	clientID := "id"

	go func() {
		LoadChatID()
		u := tgbotapi.NewUpdate(0)
		u.Timeout = 30
		updates := bot.GetUpdatesChan(u)

		for update := range updates {
			// receive emm info and send message for user
			resp, err := getData.RequestReceiveing(url, clientID, token)
			if err != nil {
				fmt.Println(err)
				time.Sleep(30 * time.Second)
				continue
			}
			data, err := jSon.EncodingJson(resp)
			if err != nil {
				fmt.Print(err)
				time.Sleep(15 * time.Second)
				continue
			}

			if update.Message != nil {
				if update.Message.Text == "/last_emission" {
					lastEmm, err := timeRes.TimeResult(data)
					if err != nil {
						fmt.Println(err)
						time.Sleep(10 * time.Second)
						continue
					}
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, lastEmm)
					bot.Send(msg)
					fmt.Println(update.Message.Chat.UserName)
				}
				if update.Message.Text == "/start" {
					lastEmm, err := timeRes.TimeResult(data)
					if err != nil {
						fmt.Println(err)
						time.Sleep(10 * time.Second)
						continue
					}
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Здорово, мужик! Ты подписался на оповещение о выбросах!\n"+lastEmm)
					bot.Send(msg)
					fmt.Println(update.Message.Chat.UserName)
				}
				if update.Message.Text == "/promocodes" {
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, getData.ParseFunc())
					bot.Send(msg)
					fmt.Println(update.Message.Chat.UserName)
				}
			}
			if !searchID(update.Message.Chat.ID) {
				ChatIDs = append(ChatIDs, update.Message.Chat.ID)
				SaveChatID()
				fmt.Println("Find New ID: ", update.Message.Chat.ID, update.Message.Chat.UserName)
			}
		}

	}()

}
