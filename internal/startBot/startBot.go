package startBot

import (
	"stalcraftBot/configs"
	"stalcraftBot/internal/emissionInfo"
	"stalcraftBot/internal/tgBot"
	"sync"
)

func StartBot() {
	configs.GetConfigsKeys()
	wg := new(sync.WaitGroup)
	wg.Add(2)

	go emissionInfo.GetEmissionData()
	go tgBot.BotChating()

	wg.Wait()
}
