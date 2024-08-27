package main

import (
	"fmt"
	"sync"
	"time"

	"main.go/internal"
)

var Data internal.EmissionInfo

func main() {
	// this case show you work with demoAPI. you have to change to the actual token and url
	url := "https://eapi.stalcraft.net/ru/emission"
	token := "ZkoXovcbrWXeUKLyyjtBprhwIm0ECyiNnCDnCfQc"
	clientID := "627"
	wg := &sync.WaitGroup{}
	wg.Add(4)

	go func() {

		internal.BotReadSave()

	}()
	// send auth and receive info
	go func() {

		for {
			resp, err := internal.RequestReceiveing(url, clientID, token)
			if err != nil {
				fmt.Println(err)
				time.Sleep(60 * time.Second)
				continue
			}
			//json encode
			Data, err = internal.EncodingJson(resp)
			if err != nil {
				fmt.Println(err)
				time.Sleep(10 * time.Second)
			}

			if Data.CurrentStart != "" {
				// print result for users
				currEm, err := internal.CurrentEmissionResult(Data)
				if err != nil {
					fmt.Println(err)
				}

				lastEm, err := internal.TimeResult(Data)
				if err != nil {
					fmt.Println(err)
				}
				textResult := fmt.Sprintf("\n%v\n%v", currEm, lastEm)
				//send telegram message
				internal.TelegramBot(textResult)
				Data.CurrentStart = ""
				time.Sleep(3 * time.Minute)
				internal.TelegramBot("Еще немного и можно будет собирать артефакты!")
			}
			fmt.Println("Request done", time.Now().Format(time.TimeOnly), Data)
			time.Sleep(60 * time.Second)

		}

	}()

	wg.Wait()
	fmt.Println("Work out")
}
