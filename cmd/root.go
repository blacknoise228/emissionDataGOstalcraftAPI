package cmd

import (
	"fmt"
	"os"
	"stalcraftbot/configs"
	"stalcraftbot/internal/logs"
	"stalcraftbot/internal/start"
	"stalcraftbot/pkg/api"

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
			Conf.API.BotAPI.PortTgBot = port
			fmt.Println("tgBot started")
			start.StartBot(Conf)
		}
		if startcrawler {
			Conf.API.BotAPI.PortTgBot = port
			fmt.Println("crawler started")
			start.StartCrawler(Conf)
		}
		if adminapi {
			Conf.API.AdminAPI.PortAdminAPI = port
			fmt.Println("adminAPI started")
			api.StartAdminAPI(Conf)
		}
	},
}
var loglvl = &cobra.Command{
	Use:   "loglvl",
	Short: "set your configurations",
	Long:  `Set up your configurations and change setup in config file`,
	Run: func(cmd *cobra.Command, args []string) {

		if debug {
			info, errors = false, false
			Conf.Logs.LogLvl = "debug"
			viper.Set("loglevel", Conf.Logs.LogLvl)
			viper.WriteConfig()
			fmt.Println("Log level is a DEBUG")
		}
		if info {
			debug, errors = false, false
			Conf.Logs.LogLvl = "info"
			viper.Set("loglevel", Conf.Logs.LogLvl)
			viper.WriteConfig()
			fmt.Println("Log level is a INFO")
		}
		if errors {
			debug, info = false, false
			Conf.Logs.LogLvl = "error"
			viper.Set("loglevel", Conf.Logs.LogLvl)
			viper.WriteConfig()
			fmt.Println("Log level is a ERRORS")
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

	rootCmd.Flags().BoolVarP(&startbot, "startbot", "b", false, "start telegram bot")
	rootCmd.Flags().IntVar(&port, "port", 8080, "set port for work api")
	rootCmd.Flags().BoolVarP(&startcrawler, "crawler", "c", false, "start stalcraft API info handler")
	rootCmd.Flags().BoolVarP(&adminapi, "adminapi", "a", false, "start adminAPI users control tool")
	rootCmd.AddCommand(loglvl)
	loglvl.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "log level = debug")
	loglvl.PersistentFlags().BoolVarP(&info, "info", "i", true, "log level = info")
	loglvl.PersistentFlags().BoolVarP(&errors, "errors", "e", false, "log level = errors")
}

// Banner and package info
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
