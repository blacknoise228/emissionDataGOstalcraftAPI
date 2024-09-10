package jsWorker

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"stalcraftBot/internal/logs"
)

const chatIDsFile = "/var/tmp/chat_ids.json"

// memory chats users
var BlackList []int64

// Save chatID to json file
func SaveChatID() {
	file, err := os.Create(chatIDsFile)
	if err != nil {
		logs.Logger.Error().Err(err).Msg("Create chat id file err")
	}
	defer file.Close()
	logs.Logger.Debug().Msg("Create chat id file done")
	for i := range Users {
		Users[i].ID = i + 1
	}
	err = json.NewEncoder(file).Encode(Users)
	if err != nil {
		logs.Logger.Error().Err(err).Msg("Save chat id to file error")
	}
	logs.Logger.Debug().Msg("Save chat id to file done")
}

// Load chatID from json file
func LoadChatID() error {

	file, err := os.Open(chatIDsFile)
	if err != nil {
		logs.Logger.Error().Err(err).Msg("Load chat id file error")
		return err
	}
	defer file.Close()
	logs.Logger.Debug().Msg("Open chat ids file done")

	err = json.NewDecoder(file).Decode(&Users)
	if err != nil {
		logs.Logger.Error().Err(err).Msg("Load Chat id error")
		return err
	}
	logs.Logger.Debug().Msg("Load chat ids from file done")
	return nil
}

// Finder in Users user id
func SearchID(num int64) bool {
	for _, v := range Users {
		if v.UserID == num {
			logs.Logger.Info().Msg(fmt.Sprint("Request ID: ", v.UserID))
			return true
		}
	}
	return false
}

// Returned a quantity of users
func QuantityUsers() int {
	LoadChatID()
	return len(Users)
}

// Loading last emission data from file
func LoadEmData(path string) string {
	file, err := os.Open(path)
	if err != nil {
		logs.Logger.Error().Err(err).Msg("Load emission data error")
	}
	defer file.Close()
	logs.Logger.Debug().Msg("Open emission data file done")
	reader, er := io.ReadAll(file)
	if er != nil {
		logs.Logger.Error().Err(er).Msg("load read emission data error")
	}
	logs.Logger.Debug().Msg("Load emission data from file done")
	return string(reader)
}

func SaveToBlackList(id int64) {
	LoadChatID()
	for i := range Users {
		if Users[i].UserID == id {
			Users[i].Blocked = true
			SaveChatID()
			return
		}
	}
}

// if the user is on the black list, then the function will return true
func SearchToBlackList(id int64) bool {
	LoadChatID()
	for _, i := range Users {
		if i.Blocked && i.UserID == id {

			return true
		}
	}
	return false
}
