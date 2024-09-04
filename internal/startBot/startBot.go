package startBot

import (
	"stalcraftBot/internal/emissionInfo"
	"stalcraftBot/internal/tgBot"
	"sync"
)

func StartBot() {
	wg := new(sync.WaitGroup)
	wg.Add(2)

	go emissionInfo.GetEmissionData()
	go tgBot.BotChating()

	wg.Wait()
}
