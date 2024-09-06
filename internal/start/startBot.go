package start

import (
	"stalcraftBot/configs"
	"stalcraftBot/internal/logs"
	"stalcraftBot/internal/tgBot"
	"sync"
)

func StartBot() {
	logs.Logger.Debug().Msg("Func StartBot is Run")
	configs.GetConfigsKeys()
	wg := new(sync.WaitGroup)
	wg.Add(1)

	go tgBot.BotChating()

	wg.Wait()
}
