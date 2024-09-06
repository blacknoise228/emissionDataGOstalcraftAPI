/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"stalcraftBot/configs"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// configsCmd represents the configs command
var configsCmd = &cobra.Command{
	Use:   "loglvl",
	Short: "set your configurations",
	Long:  `Set up your configurations and change setup in config file`,
	Run: func(cmd *cobra.Command, args []string) {

		configs.GetConfigs()
		if debug {
			info, errors = false, false
			viper.Set("loglevel", "debug")
			viper.WriteConfig()
			fmt.Println("Log level is a DEBUG")
		}
		if info {
			debug, errors = false, false
			viper.Set("loglevel", "info")
			viper.WriteConfig()
			fmt.Println("Log level is a INFO")
		}
		if errors {
			debug, info = false, false
			viper.Set("loglevel", "info")
			viper.WriteConfig()
			fmt.Println("Log level is a ERRORS")
		}
	},
}

var (
	debug  bool
	info   bool
	errors bool
)

func init() {
	rootCmd.AddCommand(configsCmd)
	configsCmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "log level = debug")
	configsCmd.PersistentFlags().BoolVarP(&info, "info", "i", true, "log level = info")
	configsCmd.PersistentFlags().BoolVarP(&errors, "errors", "e", false, "log level = errors")
}
