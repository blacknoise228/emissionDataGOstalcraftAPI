package logs

import (
	"os"
	"stalcraftbot/configs"

	"github.com/rs/zerolog"
)

var Logger = zerolog.New(os.Stdout).With().Timestamp().Logger()

// Starting logger. Accepts the Config type and sets the loglevel based on Config.LogLvl
func StartLogger(conf *configs.Config) {

	switch conf.Logs.LogLvl {
	case "debug":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case "error":
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	default:
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}
	Logger.Debug().Msg("Logger is working")
}
