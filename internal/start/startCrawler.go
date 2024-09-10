package start

import (
	"stalcraftBot/configs"
	"stalcraftBot/internal/emissionInfo"
	"stalcraftBot/internal/logs"

	"sync"
)

func StartCrawler(conf *configs.Config) {

	logs.Logger.Debug().Msg("Func StartCrawler is Run")
	wg := new(sync.WaitGroup)
	wg.Add(1)

	go emissionInfo.GetEmissionData(conf)

	wg.Wait()

}
