package start

import (
	"stalcraftBot/configs"
	"stalcraftBot/internal/logs"
	"stalcraftBot/internal/tgBot"
	"stalcraftBot/pkg/api"

	"sync"
)

func StartBot(conf *configs.Config) {
	logs.Logger.Debug().Msg("Func StartBot is Run")
	wg := new(sync.WaitGroup)
	wg.Add(2)

	go api.DataMessageAPI(conf)
	go tgBot.BotChating()

	wg.Wait()
}
