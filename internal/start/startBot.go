package start

import (
	"stalcraftBot/configs"
	"stalcraftBot/internal/logs"
	"stalcraftBot/internal/tgBot"
	"stalcraftBot/pkg/api"

	"sync"
)

func StartBot() {
	logs.Logger.Debug().Msg("Func StartBot is Run")
	configs.GetConfigs()
	wg := new(sync.WaitGroup)
	wg.Add(2)

	go api.DataMessageAPI()
	go tgBot.BotChating()

	wg.Wait()
}
