package jSon

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"stalcraftBot/internal/logs"
)

const chatIDsFile = "/var/tmp/chat_ids.json"

// save chatID to json file
func SaveChatID() {
	file, err := os.Create(chatIDsFile)
	if err != nil {
		logs.Logger.Err(err).Msg("Create chat id file err")
	}
	defer file.Close()
	logs.Logger.Debug().Msg("Create chat id file done")
	for i := range Users {
		Users[i].ID = i + 1
	}
	err = json.NewEncoder(file).Encode(Users)
	if err != nil {
		logs.Logger.Err(err).Msg("Save chat id to file error")
	}
	logs.Logger.Debug().Msg("Save chat id to file done")
}

// load chatID from json file
func LoadChatID() {

	file, err := os.Open(chatIDsFile)
	if err != nil {
		logs.Logger.Err(err).Msg("Load chat id file error")
	}
	defer file.Close()
	logs.Logger.Debug().Msg("Open chat ids file done")

	err = json.NewDecoder(file).Decode(&Users)
	if err != nil {
		logs.Logger.Err(err).Msg("Load Chat id error")
	}
	logs.Logger.Debug().Msg("Load chat ids from file done")
}

// finder in ChatIDs user id
func SearchID(num int64) bool {
	i := 0
	b := false
	for _, v := range Users {

		if v.UserID == num {
			i++
			logs.Logger.Info().Msg(fmt.Sprint("Request ID: ", v.UserID))
		}
	}
	if i > 0 {

		b = true
	}

	return b
}

func QuantityUsers() int {
	LoadChatID()
	return len(Users)
}

func LoadEmData() string {
	file, err := os.Open("/var/tmp/emissionData.txt")
	if err != nil {
		logs.Logger.Err(err).Msg("Load emission data error")
	}
	defer file.Close()
	logs.Logger.Debug().Msg("Open emission data file done")
	reader, er := io.ReadAll(file)
	if er != nil {
		logs.Logger.Err(er).Msg("load read emission data error")
	}
	logs.Logger.Debug().Msg("Load emission data from file done")
	return string(reader)

}
