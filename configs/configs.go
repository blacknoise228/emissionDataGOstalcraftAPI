package configs

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	LogLvl           string
	PortAdminAPI     int
	PortTgBot        int
	StalcraftID      string
	StalcraftTgToken string
	StalcraftToken   string
}

var (
	configName = "config"
	configType = "yaml"
	configPath = "./configs"
)

// Init config file and return structure
func InitConfig() *Config {

	SetConfig(configName, configType, configPath)

	viper.BindEnv("stalcraft_token", "STALCRAFT_TOKEN")
	viper.BindEnv("stalcraft_id", "STALCRAFT_ID")
	viper.BindEnv("stalcraft_tg_token", "STALCRAFT_TG_TOKEN")

	viper.AutomaticEnv()

	conf := Config{
		LogLvl:           viper.GetString("loglevel"),
		PortAdminAPI:     viper.GetInt("port_adminapi"),
		PortTgBot:        viper.GetInt("port_tgbot"),
		StalcraftID:      viper.GetString("stalcraft_id"),
		StalcraftTgToken: viper.GetString("stalcraft_tg_token"),
		StalcraftToken:   viper.GetString("stalcraft_token"),
	}
	return &conf
}

// Setup configuration file
func SetConfig(name string, configType string, path string) {

	viper.SetConfigName(name)
	viper.SetConfigType(configType)
	viper.AddConfigPath(path)

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("ERROR Read config file: %v", err)
	}

}
