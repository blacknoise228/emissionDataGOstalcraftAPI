package internal

import (
	"encoding/json"
	"fmt"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const chatIDsFile = "/home/blacknoise/rbPi/stalcraftAPI/data/chat_ids.json"

// memory chats users
var ChatIDs []int64

// set your telegram bot token from @BotFather
var telegramToken string = "token"

// make bot
var bot, _ = tgbotapi.NewBotAPI(telegramToken)

func TelegramBot(s string) {

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

// save chatID to json file
func SaveChatID() {
	file, err := os.Create(chatIDsFile)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	err = json.NewEncoder(file).Encode(ChatIDs)
	if err != nil {
		fmt.Println(err)
	}
}

// load chatID from json file
func LoadChatID() {

	file, err := os.Open(chatIDsFile)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&ChatIDs)
	if err != nil {
		fmt.Println(err)
	}
}

func BotReadSave(message string) {
	// update chats and save id to json file
	go func() {
		LoadChatID()
		u := tgbotapi.NewUpdate(0)
		u.Timeout = 30
		updates := bot.GetUpdatesChan(u)
		go func() {
			for update := range updates {
				if update.Message != nil {
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
					bot.Send(msg)
				}
			}
		}()
		go func() {
			for update := range updates {
				for _, id := range ChatIDs {
					if update.Message.Chat.ID == id {
						return
					} else {
						ChatIDs = append(ChatIDs, update.Message.Chat.ID)
						SaveChatID()
						fmt.Println(ChatIDs)
					}
				}
			}
		}()

	}()
}
