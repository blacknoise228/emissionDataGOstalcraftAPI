package emissionInfo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"stalcraftbot/configs"
	"stalcraftbot/internal/jsWorker"
	"stalcraftbot/internal/logs"
	"stalcraftbot/internal/rediska"
	"stalcraftbot/internal/timeRes"
	"stalcraftbot/pkg/getData"
	"strconv"
	"strings"

	"sync"
	"time"
)

const EmissionDataFile string = "/var/tmp/emissionData.txt"
const CurrentEmissionDataFile string = "/var/tmp/currentEmissionData.txt"

// Send request to stalcraftAPI server, save last emission data to file.
// If current emission data not "", send current emission data to tgBotAPI
func GetEmissionData(conf *configs.Config) {
	var Data jsWorker.EmissionInfo
	// this case show you work with demoAPI. you have to change to the actual token and url
	url := "https://eapi.stalcraft.net/ru/emission"
	token := conf.Stalcraft.StalcraftToken
	clientID := conf.Stalcraft.StalcraftID
	port := strconv.Itoa(conf.API.BotAPI.PortTgBot)
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
			//Data.CurrentStart = "2019-08-24T14:15:22Z"
			if Data.CurrentStart != "" {
				CurrentEmissionDataSendToBotAPI(Data, port)
				Data.CurrentStart = ""
			}

			lastEm, err := timeRes.TimeResult(Data)
			if err != nil {
				logs.Logger.Err(err).Msg("TimeResult Data error")
				continue
			}
			if err = rediska.SaveLastEmissionDataToRedis(lastEm); err != nil {
				logs.Logger.Error().Msg(fmt.Sprintf("Saving to REDIS ERRROR: %v", err))
			}

			logs.Logger.Info().Msg(fmt.Sprint("Request done", Data))
			time.Sleep(60 * time.Second)
		}

	}()
	wg.Wait()
	fmt.Println("Work out")
}

func CurrentEmissionDataSendToBotAPI(data jsWorker.EmissionInfo, port string) {
	// print result for users
	fmt.Println(data)
	for {
		//send to telegram botAPI message
		jData, err := json.Marshal(data)
		if err != nil {
			logs.Logger.Err(err).Msg("Marshal api message error")
			continue
		}
		reader := strings.NewReader(string(jData))
		resp, err := http.Post("http://bot:"+port+"/emdata", "json", reader)
		if err != nil {
			logs.Logger.Err(err).Msg("Error send signal to botAPI")
			time.Sleep(5 * time.Second)
			continue
		}
		defer resp.Body.Close()
		logs.Logger.Info().Msg("Send current emission data to botAPI done")
		time.Sleep(4 * time.Minute)
		return
	}
}
