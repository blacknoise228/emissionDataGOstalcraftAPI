package main

import (
	"fmt"
	"sync"
	"time"

	"main.go/internal"
)

func main() {
	// this case show you work with demoAPI. you have to change to the actual token and url
	url := "https://eapi.stalcraft.net/ru/emission"
	token := "ZkoXovcbrWXeUKLyyjtBprhwIm0ECyiNnCDnCfQc"
	clientID := "627"
	respInfo := make(chan internal.EmissionInfo)
	wg := &sync.WaitGroup{}
	wg.Add(4)

	go func() {

		internal.BotReadSave()

	}()
	// send auth and receive info
	go func() {
		defer wg.Done()
		defer close(respInfo)
		for {
			resp, err := internal.RequestReceiveing(url, clientID, token)
			if err != nil {
				fmt.Println(err)
			}
			//json encode
			data := internal.EncodingJson(resp)
			respInfo <- data
			time.Sleep(30 * time.Second)
		}
	}()

	// emission start message
	go func() {
		defer wg.Done()
		dataRes := <-respInfo

		if dataRes.CurrentStart != "" {
			// print result for users
			currEm, err := internal.CurrentEmissionResult(dataRes)
			if err != nil {
				fmt.Println(err)
			}

			lastEm, err := internal.TimeResult(dataRes)
			if err != nil {
				fmt.Println(err)
			}
			textResult := fmt.Sprintf("\n%v\n%v", currEm, lastEm)
			//send telegram message
			internal.TelegramBot(textResult)
			dataRes.CurrentStart = ""
		}

	}()
	wg.Wait()
	fmt.Println("Work out")
}
