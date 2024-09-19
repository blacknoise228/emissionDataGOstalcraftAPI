package tgBot

import (
	"database/sql"
	"fmt"
	"stalcraftbot/internal/logs"
	"stalcraftbot/internal/rediska"
	"stalcraftbot/pkg/getData"
	"stalcraftbot/pkg/postgres"

	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/spf13/viper"
)

// Make telegram bot and returned telegramBot
func MakeBot() (*tgbotapi.BotAPI, *sql.DB) {
	// set your telegram bot token from @BotFather
	var telegramToken string = viper.GetString("api.tgbot.token")
	var bot, err = tgbotapi.NewBotAPI(telegramToken)
	if err != nil {
		logs.Logger.Fatal().Msg(fmt.Sprintln("Make bot error", err))
	}
	logs.Logger.Debug().Msg(fmt.Sprintln("Make bot is done"))
	db := postgres.InitDB()
	return bot, db
}

// Sending Accepted message to all users from database
func SendMessageTG(s string) {
	bot, db := MakeBot()
	// print bot info and send message
	botUser, err := bot.GetMe()
	if err != nil {
		logs.Logger.Fatal().Err(err).Msg("Send message bot error ")
	}
	fmt.Printf("\nBot user: %v\n", botUser)
	// Load users from DB
	postgres.LoadChatID(db)
	// Send message
	for _, user := range postgres.Users {
		msg := tgbotapi.NewMessage(user.UserID, s)
		bot.Send(msg)
	}
}

// Function run cycle updating message from users and send response
func BotChating() {
	bot, db := MakeBot()
	time.Sleep(1 * time.Second)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 30
	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		// send message for user
		if update.Message == nil {
			logs.Logger.Debug().Msg("Messege is nil!!!!")
			continue
		} else if update.Message != nil {
			if update.Message.Text == "/last_emission" {
				lastEmm, err := rediska.LoadEmDataFromRedis()
				if err != nil {
					logs.Logger.Error().Msg(fmt.Sprintf("Reading from REDIS ERRROR: %v", err))
				}
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, lastEmm)
				bot.Send(msg)
				logs.Logger.Info().Msg(fmt.Sprint("User: ", update.Message.Chat.UserName))
				logs.Logger.Debug().Msg("/last_emission msg send done")
			}
			if update.Message.Text == "/start" {
				lastEmm, err := rediska.LoadEmDataFromRedis()
				if err != nil {
					logs.Logger.Error().Msg(fmt.Sprintf("Reading from REDIS ERRROR: %v", err))
				}
				msg := tgbotapi.NewMessage(update.Message.Chat.ID,
					"Здорово, мужик! Ты подписался на оповещение о выбросах!\n"+lastEmm)
				bot.Send(msg)
				logs.Logger.Info().Msg(fmt.Sprint("User: ", update.Message.Chat.UserName))
				logs.Logger.Debug().Msg("/start msg send done")
			}
			if update.Message.Text == "/promocodes" {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, getData.ParseFunc())
				bot.Send(msg)
				logs.Logger.Info().Msg(fmt.Sprint("User: ", update.Message.Chat.UserName))
				logs.Logger.Debug().Msg("/last_emission msg send done")
			}
			if !postgres.SearchID(db, update.Message.Chat.ID) {
				logs.Logger.Info().Msg("Find new user")
				user := postgres.User{
					UserID: update.Message.Chat.ID,
					Name:   update.Message.Chat.UserName,
				}
				postgres.SaveChatID(db, user)
			}
		}
	}
}
