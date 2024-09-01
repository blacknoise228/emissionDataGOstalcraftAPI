package emissionInfo

import (
	"fmt"
	"os"
	"sync"
	"time"

	"stalcraftBot/internal/getData"
	"stalcraftBot/internal/jSon"
	"stalcraftBot/internal/tgBot"
	"stalcraftBot/internal/timeRes"
)

func GetEmissionData() {
	var Data jSon.EmissionInfo
	// this case show you work with demoAPI. you have to change to the actual token and url
	url := "https://eapi.stalcraft.net/ru/emission"
	token := "ZkoXovcbrWXeUKLyyjtBprhwIm0ECyiNnCDnCfQc"
	clientID := "627"

	wg := &sync.WaitGroup{}
	wg.Add(1)

	// send auth and receive info
	go func() {

		for {
			resp, err := getData.RequestReceiveing(url, clientID, token)
			if err != nil {
				fmt.Println(err)
				time.Sleep(60 * time.Second)
				continue
			}
			//json encode
			Data, err = jSon.EncodingJson(resp)
			if err != nil {
				fmt.Println(err)
				time.Sleep(10 * time.Second)
			}
			lastEm, err := timeRes.TimeResult(Data)
			if err != nil {
				fmt.Println(err)
			}
			SaveEmData(lastEm)

			if Data.CurrentStart != "" {
				// print result for users
				currEm, err := timeRes.CurrentEmissionResult(Data)
				if err != nil {
					fmt.Println(err)
				}

				lastEm, err := timeRes.TimeResult(Data)
				if err != nil {
					fmt.Println(err)
				}
				textResult := fmt.Sprintf("\n%v\n%v", currEm, lastEm)
				//send telegram message
				tgBot.SendMessageTG(textResult)
				Data.CurrentStart = ""
				time.Sleep(3 * time.Minute)
				tgBot.SendMessageTG("Еще немного и можно будет собирать артефакты!")
			}
			fmt.Println("Request done", time.Now().Format(time.TimeOnly), Data)
			time.Sleep(60 * time.Second)

		}

	}()
	wg.Wait()
	fmt.Println("Work out")
}
func SaveEmData(data string) {
	file, err := os.Create("/var/tmp/emissionData.txt")
	if err != nil {
		fmt.Println("Create emData:", err)
	}
	defer file.Close()
	file.WriteString(data)
}
