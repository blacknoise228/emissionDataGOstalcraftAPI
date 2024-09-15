package configs

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

// Init config file and return structure
func InitConfig() *Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.SetConfigFile("./configs/config.yaml")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("ERROR Read config file: %v", err)
	}
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("ERROR Unmarshall config: %v", err)
	}
	fmt.Println(config, viper.AllSettings())
	return &config
}
