package start

import (
	"stalcraftBot/configs"
	"stalcraftBot/internal/emissionInfo"
	"stalcraftBot/internal/logs"
	"sync"
)

func StartCrawler() {

	logs.Logger.Debug().Msg("Func StartCrawler is Run")
	configs.GetConfigsKeys()
	wg := new(sync.WaitGroup)
	wg.Add(1)

	go emissionInfo.GetEmissionData()

	wg.Wait()

}
