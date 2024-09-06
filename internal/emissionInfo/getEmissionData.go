package emissionInfo

import (
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"

	"stalcraftBot/internal/jsWorker"
	logs "stalcraftBot/internal/logs"
	"stalcraftBot/internal/timeRes"
	"stalcraftBot/pkg/getData"

	"github.com/spf13/viper"
)

const EmissionDataFile string = "/var/tmp/emissionData.txt"
const CurrentEmissionDataFile string = "/var/tmp/currentEmissionData.txt"

func GetEmissionData() {
	var Data jsWorker.EmissionInfo
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
				logs.Logger.Err(err).Msg("Request receiveing error")
				time.Sleep(50 * time.Second)
				continue
			}
			Data, err = jsWorker.EncodingJson(resp)
			if err != nil {
				logs.Logger.Err(err).Msg("EncodingJson error")
				time.Sleep(10 * time.Second)
				continue
			}
			lastEm, err := timeRes.TimeResult(Data)
			if err != nil {
				logs.Logger.Err(err).Msg("TimeResult Data error")
				continue
			}
			SaveLastEmissionDataToFile(lastEm)

			//Data.CurrentStart = "2019-08-24T14:15:22Z"
			if Data.CurrentStart != "" {
				CurrentEmissionDataSendToBotAPI(Data)
			}
			logs.Logger.Info().Msg(fmt.Sprint("Request done", Data))
			time.Sleep(60 * time.Second)

		}

	}()
	wg.Wait()
	fmt.Println("Work out")
}
func SaveLastEmissionDataToFile(data string) {
	file, err := os.Create(EmissionDataFile)
	if err != nil {
		logs.Logger.Error().Err(err).Msg("Create emData file error")
	}
	defer file.Close()
	file.WriteString(data)
	logs.Logger.Debug().Msg("Save emission data file done")
}
func SaveCurrentEmissionDataToFile(data string) {
	file, err := os.Create(CurrentEmissionDataFile)
	if err != nil {
		logs.Logger.Error().Err(err).Msg("Create currentEmData file error")
	}
	defer file.Close()
	file.WriteString(data)
	logs.Logger.Debug().Msg("Save current emission data file done")
}
func CurrentEmissionDataSendToBotAPI(data jsWorker.EmissionInfo) {
	// print result for users
	for {
		currEm, err := timeRes.CurrentEmissionResult(data)
		if err != nil {
			logs.Logger.Err(err).Msg("Current Emission Data error")
		}

		lastEm, err := timeRes.TimeResult(data)
		if err != nil {
			logs.Logger.Err(err).Msg("TimeResult Data error")
		}
		textResult := fmt.Sprintf("\n%v\n%v", currEm, lastEm)
		SaveCurrentEmissionDataToFile(textResult)
		//send to telegram botAPI message
		resp, err := http.Get("http://localhost:1234/emdata")
		if err != nil {
			logs.Logger.Err(err).Msg("Error send signal to botAPI")
			time.Sleep(5 * time.Second)
			continue
		}
		logs.Logger.Info().Msg("Send current emission data to botAPI done")
		resp.Body.Close()
		data.CurrentStart = ""
		time.Sleep(4 * time.Minute)
		return
	}
}
