/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"stalcraftBot/configs"
	"stalcraftBot/internal/logs"
	"stalcraftBot/internal/start"
	"stalcraftBot/pkg/api"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	Conf         = configs.InitConfig()
	startbot     bool
	startcrawler bool
	adminapi     bool
	port         int
	debug        bool
	info         bool
	errors       bool
)
var rootCmd = &cobra.Command{
	Use:   "stalcraftbot",
	Short: "TelegramAPIbot for stalcraft:x game",
	Long:  stringInfo,
	Run: func(cmd *cobra.Command, args []string) {
		logs.StartLogger(Conf)
		if startbot {
			Conf.PortTgBot = port
			fmt.Println("tgBot started")
			start.StartBot(Conf)
		}
		if startcrawler {
			fmt.Println("crawler started")
			start.StartCrawler(Conf)
		}
		if adminapi {
			Conf.PortAdminAPI = port
			fmt.Println("adminAPI started")
			api.StartAdminAPI(Conf)
		}
	},
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}
var configsCmd = &cobra.Command{
	Use:   "loglvl",
	Short: "set your configurations",
	Long:  `Set up your configurations and change setup in config file`,
	Run: func(cmd *cobra.Command, args []string) {

		if debug {
			info, errors = false, false
			Conf.LogLvl = "debug"
			viper.Set("loglevel", Conf.LogLvl)
			viper.WriteConfig()
			fmt.Println("Log level is a DEBUG")
		}
		if info {
			debug, errors = false, false
			Conf.LogLvl = "info"
			viper.Set("loglevel", Conf.LogLvl)
			viper.WriteConfig()
			fmt.Println("Log level is a INFO")
		}
		if errors {
			debug, info = false, false
			Conf.LogLvl = "error"
			viper.Set("loglevel", Conf.LogLvl)
			viper.WriteConfig()
			fmt.Println("Log level is a ERRORS")
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.stalcraftBot.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolVarP(&startbot, "startbot", "b", false, "start telegram bot")
	rootCmd.Flags().IntVar(&port, "port", 8080, "set port for work api")
	rootCmd.Flags().BoolVarP(&startcrawler, "crawler", "c", false, "start stalcraft API info handler")
	rootCmd.Flags().BoolVarP(&adminapi, "adminapi", "a", false, "start adminAPI users control tool")
	rootCmd.AddCommand(configsCmd)
	configsCmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "log level = debug")
	configsCmd.PersistentFlags().BoolVarP(&info, "info", "i", true, "log level = info")
	configsCmd.PersistentFlags().BoolVarP(&errors, "errors", "e", false, "log level = errors")
}

var stringInfo = `

⡏⢉⣉⡉⠉⣿⣿⣉⡉⠉⣉⣹⣿⣿⠉⠉⠉⠉⢹⣿⣿⠉⢹⣿⣿⣿⣿⣿⠉⣉⣉⠉⣿⣿⡏⠉⠉⠉⠉⣿⣿⡏⠉⠉⠉⠉⣿⣿⡏⠉⠉⠉⠉⣿⣿⣍⡉⠉⣉⣹
⡇⠸⣿⡇⣠⣿⣿⣿⡇⠀⣿⣿⣿⣿⠀⢸⣿⠀⢸⣿⣿⠀⢸⣿⣿⣿⣿⣿⠀⣿⣿⣤⣿⣿⡇⠀⣿⡇⠀⣿⣿⡇⠀⣿⡇⠀⣿⣿⡇⠀⣿⣇⣤⣿⣿⣿⣿⠀⣿⣿
⣧⡀⠙⢿⣿⣿⣿⣿⡇⠀⣿⣿⣿⣿⠀⢸⣿⠀⢸⣿⣿⠀⢸⣿⣿⣿⣿⣿⠀⣿⣿⣿⣿⣿⡇⠀⠛⢀⣿⣿⣿⡇⠀⣿⡇⠀⣿⣿⡇⠀⠛⢻⣿⣿⣿⣿⣿⠀⣿⣿
⣿⣿⣆⠀⠹⣿⣿⣿⡇⠀⣿⣿⣿⣿⠀⢀⣀⠀⢸⣿⣿⠀⢸⣿⣿⣿⣿⣿⠀⣿⣿⣿⣿⣿⡇⠀⣴⡆⠀⣿⣿⡇⠀⣀⡀⠀⣿⣿⡇⠀⣶⣾⣿⣿⣿⣿⣿⠀⣿⣿
⡏⢸⣿⡇⠀⣿⣿⣿⡇⠀⣿⣿⣿⣿⠀⢸⣿⠀⢸⣿⣿⠀⢸⣿⠀⢻⣿⣿⠀⣿⣿⠀⣿⣿⡇⠀⣿⡇⠀⣿⣿⡇⠀⣿⡇⠀⣿⣿⡇⠀⣿⣿⣿⣿⣿⣿⣿⠀⣿⣿
⣇⣈⣉⣁⣀⣿⣿⣿⣇⣀⣿⣿⣿⣿⣀⣸⣿⣀⣸⣿⣿⣀⣈⣉⣀⣸⣿⣿⣀⣉⣉⣀⣿⣿⣇⣀⣿⣇⣀⣿⣿⣇⣀⣿⣇⣀⣿⣿⣇⣀⣿⣿⣿⣿⣿⣿⣿⣀⣿⣿


	TelegramAPIbot for stalcraft:x game
	This bot get emission info and send it for all users
	Bot show you quantity users
	Bot get promocodes from steam page`
