package internal

import (
	"encoding/json"
	"fmt"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const chatIDsFile = "/home/chat_ids.json"

// memory chats users
var ChatIDs []int64

// set your telegram bot token from @BotFather
var telegramToken string = "7544255529:AAGxUryzd9Io2k4pcLzXwrwcdjk8HEvB134"

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

func BotReadSave() {
	// update chats and save id to json file
	url := "https://eapi.stalcraft.net/ru/emission"
	token := "ZkoXovcbrWXeUKLyyjtBprhwIm0ECyiNnCDnCfQc"
	clientID := "627"

	go func() {
		LoadChatID()
		u := tgbotapi.NewUpdate(0)
		u.Timeout = 30
		updates := bot.GetUpdatesChan(u)

		for update := range updates {
			// receive emm info and send message for user
			resp, err := RequestReceiveing(url, clientID, token)
			if err != nil {
				fmt.Println(err)
			}
			data := EncodingJson(resp)

			if update.Message != nil {
				lastEmm, err := TimeResult(data)
				if err != nil {
					fmt.Println(err)
				}
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, lastEmm)
				bot.Send(msg)
			}
			if !find(update.Message.Chat.ID) {
				ChatIDs = append(ChatIDs, update.Message.Chat.ID)
				SaveChatID()
				fmt.Println("Find New ID: ", update.Message.Chat.ID)
			}
		}

	}()

}

// finder in ChatIDs user id
func find(num int64) bool {
	i := 0
	b := false
	for _, id := range ChatIDs {

		if id == num {
			i++
			fmt.Println("Request ID: ", id)
		}
	}
	if i > 0 {

		b = true
	}

	return b
}
