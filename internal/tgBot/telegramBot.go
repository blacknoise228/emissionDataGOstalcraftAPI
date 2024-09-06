package tgBot

import (
	"fmt"
	"time"

	"stalcraftBot/internal/jSon"
	"stalcraftBot/internal/logs"
	"stalcraftBot/pkg/getData"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/spf13/viper"
)

// memory chats users
var ChatIDs []int64

func MakeBot() *tgbotapi.BotAPI {
	// set your telegram bot token from @BotFather
	var telegramToken string = viper.GetString("stalcraft_tg_token")
	var bot, _ = tgbotapi.NewBotAPI(telegramToken)
	logs.Logger.Debug().Msg(fmt.Sprintln("Make bot is done"))
	return bot
}

// make bot

func SendMessageTG(s string) {
	var bot = MakeBot()
	// print bot info and send message
	botUser, err := bot.GetMe()
	if err != nil {
		logs.Logger.Fatal().Err(err).Msg("Send message bot error ")
	}
	fmt.Printf("\nBot user: %v\n", botUser)
	// read id from json file to ChatIDs slice
	jSon.LoadChatID()

	// Send message
	for _, id := range ChatIDs {
		msg := tgbotapi.NewMessage(id, s)
		bot.Send(msg)
	}
}

func BotChating() {
	var bot = MakeBot()
	time.Sleep(1 * time.Second)
	jSon.LoadChatID()
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 30
	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		// receive emm info and send message for user

		if update.Message != nil {
			if update.Message.Text == "/last_emission" {
				lastEmm := jSon.LoadEmData()
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, lastEmm)
				bot.Send(msg)
				logs.Logger.Info().Msg(fmt.Sprint("User: ", update.Message.Chat.UserName))
				logs.Logger.Debug().Msg("/last_emission msg send done")
			}
			if update.Message.Text == "/start" {
				lastEmm := jSon.LoadEmData()
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Здорово, мужик! Ты подписался на оповещение о выбросах!\n"+lastEmm)
				bot.Send(msg)
				logs.Logger.Info().Msg(fmt.Sprint("User: ", update.Message.Chat.UserName))
				logs.Logger.Debug().Msg("/last_emission msg send done")
			}
			if update.Message.Text == "/promocodes" {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, getData.ParseFunc())
				bot.Send(msg)
				logs.Logger.Info().Msg(fmt.Sprint("User: ", update.Message.Chat.UserName))
				logs.Logger.Debug().Msg("/last_emission msg send done")
			}
		}
		if !jSon.SearchID(update.Message.Chat.ID) {

			newUser := jSon.User{
				ID:     len(jSon.Users) + 1,
				UserID: update.Message.Chat.ID,
				Name:   update.Message.Chat.UserName,
			}
			jSon.Users = append(jSon.Users, newUser)
			jSon.SaveChatID()
			logs.Logger.Info().Msg(fmt.Sprint("Find New ID: ", update.Message.Chat.ID, " ", update.Message.Chat.UserName))
		}
	}
}
