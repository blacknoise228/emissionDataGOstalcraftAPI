package start

import (
	"stalcraftbot/configs"
	"stalcraftbot/internal/emissionInfo"
	"stalcraftbot/internal/logs"
	"sync"
)

func StartCrawler(conf *configs.Config) {

	logs.Logger.Debug().Msg("Func StartCrawler is Run")
	wg := new(sync.WaitGroup)
	wg.Add(1)

	go emissionInfo.GetEmissionData(conf)

	wg.Wait()

}
