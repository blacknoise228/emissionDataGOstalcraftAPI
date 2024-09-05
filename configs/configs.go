package configs

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func GetConfigsKeys() {
	viper.SetConfigName("config_keys")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("ERROR Read config file: %v", err)
	}

	viper.BindEnv("stalcraft_token", "STALCRAFT_TOKEN")
	viper.BindEnv("stalcraft_id", "STALCRAFT_ID")
	viper.BindEnv("stalcraft_tg_token", "STALCRAFT_TG_TOKEN")

	viper.AutomaticEnv()

	fmt.Printf("\nStalcraft Token: %v\nStalcraft ID: %v\nTelegram Token: %v\n",
		viper.GetString("stalcraft_token"),
		viper.GetString("stalcraft_id"),
		viper.GetString("stalcraft_tg_token"))
}
func SetConfig() {

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("ERROR Read config file: %v", err)
	}

}
