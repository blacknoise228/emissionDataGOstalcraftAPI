package tgBot

import (
	"encoding/json"
	"fmt"
	"os"
)

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

// finder in ChatIDs user id
func searchID(num int64) bool {
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

func QuantityUsers() {
	LoadChatID()
	fmt.Println(len(ChatIDs))
}
