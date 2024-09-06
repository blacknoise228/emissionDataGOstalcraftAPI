package tgBot

import (
	"fmt"
	"stalcraftBot/internal/emissionInfo"
	"stalcraftBot/internal/jsWorker"
	"stalcraftBot/internal/logs"
	"stalcraftBot/pkg/getData"

	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/spf13/viper"
)

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
	jsWorker.LoadChatID()

	// Send message
	for _, user := range jsWorker.Users {
		msg := tgbotapi.NewMessage(user.UserID, s)
		bot.Send(msg)
	}
}

func BotChating() {
	var bot = MakeBot()
	time.Sleep(1 * time.Second)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 30
	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		// send message for user
		if jsWorker.SearchToBlackList(update.Message.Chat.ID) {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Администратор заблокировал вас.")
			bot.Send(msg)
			logs.Logger.Info().Msg(fmt.Sprint("Blocked User: ", update.Message.Chat.UserName))
			logs.Logger.Debug().Msg("Blocked user write message")
			continue
		}
		if update.Message != nil {
			if update.Message.Text == "/last_emission" {
				lastEmm := jsWorker.LoadEmData(emissionInfo.EmissionDataFile)
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, lastEmm)
				bot.Send(msg)
				logs.Logger.Info().Msg(fmt.Sprint("User: ", update.Message.Chat.UserName))
				logs.Logger.Debug().Msg("/last_emission msg send done")
			}
			if update.Message.Text == "/start" {
				lastEmm := jsWorker.LoadEmData(emissionInfo.EmissionDataFile)
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
		if err := jsWorker.LoadChatID(); err != nil {
			newUser := jsWorker.User{
				ID:      len(jsWorker.Users) + 1,
				Blocked: false,
				UserID:  update.Message.Chat.ID,
				Name:    update.Message.Chat.UserName,
			}
			jsWorker.Users = append(jsWorker.Users, newUser)
			jsWorker.SaveChatID()
		}
		if !jsWorker.SearchID(update.Message.Chat.ID) {
			newUser := jsWorker.User{
				ID:      len(jsWorker.Users) + 1,
				Blocked: false,
				UserID:  update.Message.Chat.ID,
				Name:    update.Message.Chat.UserName,
			}
			jsWorker.Users = append(jsWorker.Users, newUser)
			jsWorker.SaveChatID()
			logs.Logger.Info().Msg(fmt.Sprint("Find New ID: ", update.Message.Chat.ID, " ", update.Message.Chat.UserName))
		}
	}
}
