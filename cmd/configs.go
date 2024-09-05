/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
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

		configs.SetConfig()
		if debug {
			info = false
			viper.Set("loglevel", "debug")
			viper.WriteConfig()
		}
		if info {
			debug = false
			viper.Set("loglevel", "info")
			viper.WriteConfig()
		}
	},
}

var (
	debug bool
	info  bool
)

func init() {
	rootCmd.AddCommand(configsCmd)
	configsCmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "log level = debug")
	configsCmd.PersistentFlags().BoolVarP(&info, "info", "i", true, "log level = info")
}
