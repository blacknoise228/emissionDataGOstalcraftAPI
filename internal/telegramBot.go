package internal

import (
	"fmt"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

// memory chats users
var ChatIDs = []int64{611685177}

// set your telegram bot token from @BotFather
var telegramToken string = "7544255529:AAGxUryzd9Io2k4pcLzXwrwcdjk8HEvB134"

// make bot
var bot, _ = telego.NewBot(telegramToken)

func TelegramBot(s string) {

	// print bot info
	botUser, err := bot.GetMe()
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Printf("\nBot user: %v\n", botUser)

	// Send message
	for _, id := range ChatIDs {
		msg := tu.Message(tu.ID(id), s)
		_, _ = bot.SendMessage(msg)
	}
}
