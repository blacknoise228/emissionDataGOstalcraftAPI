package emissionInfo

import (
	"fmt"
	"os"
	"sync"
	"time"

	"stalcraftBot/internal/getData"
	"stalcraftBot/internal/jSon"
	logs "stalcraftBot/internal/logs"
	"stalcraftBot/internal/tgBot"
	"stalcraftBot/internal/timeRes"

	"github.com/spf13/viper"
)

func GetEmissionData() {
	var Data jSon.EmissionInfo
	// this case show you work with demoAPI. you have to change to the actual token and url
	url := "https://eapi.stalcraft.net/ru/emission"
	token := viper.GetString("stalcraft_token")
	clientID := viper.GetString("stalcraft_id")

	wg := &sync.WaitGroup{}
	wg.Add(1)

	// send auth and receive info
	go func() {

		for {
			resp, err := getData.RequestReceiveing(url, clientID, token)
			if err != nil {
				logs.Logger.Error().Err(err).Msg("Request receiveing error")
				time.Sleep(50 * time.Second)
				continue
			}
			//json encode
			Data, err = jSon.EncodingJson(resp)
			if err != nil {
				logs.Logger.Error().Err(err).Msg("EncodingJson error")
				time.Sleep(10 * time.Second)
			}
			lastEm, err := timeRes.TimeResult(Data)
			if err != nil {
				logs.Logger.Error().Err(err).Msg("TimeResult Data error")
			}
			SaveEmData(lastEm)

			if Data.CurrentStart != "" {
				// print result for users
				currEm, err := timeRes.CurrentEmissionResult(Data)
				if err != nil {
					logs.Logger.Error().Err(err).Msg("Current Emission Data error")
				}

				lastEm, err := timeRes.TimeResult(Data)
				if err != nil {
					logs.Logger.Error().Err(err).Msg("TimeResult Data error")
				}
				textResult := fmt.Sprintf("\n%v\n%v", currEm, lastEm)
				//send telegram message
				tgBot.SendMessageTG(textResult)
				Data.CurrentStart = ""
				time.Sleep(3 * time.Minute)
				tgBot.SendMessageTG("Еще немного и можно будет собирать артефакты!")
			}
			logs.Logger.Info().Msg(fmt.Sprint("Request done", Data))
			time.Sleep(60 * time.Second)

		}

	}()
	wg.Wait()
	fmt.Println("Work out")
}
func SaveEmData(data string) {
	file, err := os.Create("/var/tmp/emissionData.txt")
	if err != nil {
		logs.Logger.Error().Err(err).Msg("Create emData file error")
	}
	defer file.Close()
	file.WriteString(data)
	logs.Logger.Debug().Msg("Save emission data file done")
}
