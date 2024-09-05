package startBot

import (
	"stalcraftBot/configs"
	"stalcraftBot/internal/emissionInfo"
	"stalcraftBot/internal/logs"
	"stalcraftBot/internal/tgBot"
	"sync"
)

func StartBot() {
	logs.Logger.Debug().Msg("Func StartBot is Run")
	configs.GetConfigsKeys()
	wg := new(sync.WaitGroup)
	wg.Add(2)

	go emissionInfo.GetEmissionData()
	go tgBot.BotChating()

	wg.Wait()
}
