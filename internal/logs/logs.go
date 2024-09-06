package logs

import (
	"os"
	"stalcraftBot/configs"

	"github.com/rs/zerolog"
	"github.com/spf13/viper"
)

var Logger = zerolog.New(os.Stdout).With().Timestamp().Logger()

func StartLogger() {

	configs.GetConfigs()
	switch viper.GetString("loglevel") {
	case "debug":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	default:
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}
	Logger.Debug().Msg("Logger is working")
}
