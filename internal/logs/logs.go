package logs

import (
	"os"
	"stalcraftBot/configs"

	"github.com/rs/zerolog"
)

var Logger = zerolog.New(os.Stdout).With().Timestamp().Logger()

func StartLogger(conf *configs.Config) {

	switch conf.LogLvl {
	case "debug":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case "error":
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	default:
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}
	Logger.Debug().Msg("Logger is working")
}
