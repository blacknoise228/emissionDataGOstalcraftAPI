package configs

import (
	"bytes"
	"fmt"
	"log"

	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

// Init config file and return structure
func InitConfig() *Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.SetConfigFile("./configs/config.yaml")

	viper.BindEnv("api.tgbot.token", "STALCRAFT_TG_TOKEN")
	viper.BindEnv("stalcraft.id", "STALCRAFT_ID")
	viper.BindEnv("stalcraft.token", "STALCRAFT_TOKEN")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("ERROR Read config file: %v", err)
	}
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("ERROR Unmarshall config: %v", err)
	}
	buf := bytes.NewBuffer(nil)
	_ = yaml.NewEncoder(buf).Encode(config)
	fmt.Println("Effective configuration:")
	fmt.Println(buf.String())
	return &config
}
