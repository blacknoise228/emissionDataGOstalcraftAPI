package emissionInfo

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"stalcraftBot/internal/jsWorker"
	"stalcraftBot/internal/logs"
	"stalcraftBot/internal/timeRes"
	"stalcraftBot/pkg/getData"

	"sync"
	"time"

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
	port := viper.GetString("port_tgbot")
	wg := &sync.WaitGroup{}
	wg.Add(1)

	// send auth and receive info
	go func() {

		for {
			Resp, err := getData.RequestReceiveing(url, clientID, token)
			if err != nil {
				logs.Logger.Err(err).Msg("Request receiveing error")
				time.Sleep(50 * time.Second)
				continue
			}
			Data, err = jsWorker.EncodingJson(Resp)
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
				CurrentEmissionDataSendToBotAPI(Resp.Body, port)
				Data.CurrentStart = ""
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
func CurrentEmissionDataSendToBotAPI(data io.Reader, port string) {
	// print result for users
	fmt.Println(data)
	for {
		//send to telegram botAPI message
		resp, err := http.Post("http://localhost:"+port+"/emdata", "json", data)
		if err != nil {
			logs.Logger.Err(err).Msg("Error send signal to botAPI")
			time.Sleep(5 * time.Second)
			continue
		}
		logs.Logger.Info().Msg("Send current emission data to botAPI done")
		resp.Body.Close()
		time.Sleep(4 * time.Minute)
		return
	}
}
