/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"stalcraftBot/internal/start"
	"stalcraftBot/pkg/api"

	"github.com/spf13/cobra"
)

var (
	startbot     bool
	startcrawler bool
	adminapi     bool
)
var rootCmd = &cobra.Command{
	Use:   "stalcraftbot",
	Short: "TelegramAPIbot for stalcraft:x game",
	Long: `

⡏⢉⣉⡉⠉⣿⣿⣉⡉⠉⣉⣹⣿⣿⠉⠉⠉⠉⢹⣿⣿⠉⢹⣿⣿⣿⣿⣿⠉⣉⣉⠉⣿⣿⡏⠉⠉⠉⠉⣿⣿⡏⠉⠉⠉⠉⣿⣿⡏⠉⠉⠉⠉⣿⣿⣍⡉⠉⣉⣹
⡇⠸⣿⡇⣠⣿⣿⣿⡇⠀⣿⣿⣿⣿⠀⢸⣿⠀⢸⣿⣿⠀⢸⣿⣿⣿⣿⣿⠀⣿⣿⣤⣿⣿⡇⠀⣿⡇⠀⣿⣿⡇⠀⣿⡇⠀⣿⣿⡇⠀⣿⣇⣤⣿⣿⣿⣿⠀⣿⣿
⣧⡀⠙⢿⣿⣿⣿⣿⡇⠀⣿⣿⣿⣿⠀⢸⣿⠀⢸⣿⣿⠀⢸⣿⣿⣿⣿⣿⠀⣿⣿⣿⣿⣿⡇⠀⠛⢀⣿⣿⣿⡇⠀⣿⡇⠀⣿⣿⡇⠀⠛⢻⣿⣿⣿⣿⣿⠀⣿⣿
⣿⣿⣆⠀⠹⣿⣿⣿⡇⠀⣿⣿⣿⣿⠀⢀⣀⠀⢸⣿⣿⠀⢸⣿⣿⣿⣿⣿⠀⣿⣿⣿⣿⣿⡇⠀⣴⡆⠀⣿⣿⡇⠀⣀⡀⠀⣿⣿⡇⠀⣶⣾⣿⣿⣿⣿⣿⠀⣿⣿
⡏⢸⣿⡇⠀⣿⣿⣿⡇⠀⣿⣿⣿⣿⠀⢸⣿⠀⢸⣿⣿⠀⢸⣿⠀⢻⣿⣿⠀⣿⣿⠀⣿⣿⡇⠀⣿⡇⠀⣿⣿⡇⠀⣿⡇⠀⣿⣿⡇⠀⣿⣿⣿⣿⣿⣿⣿⠀⣿⣿
⣇⣈⣉⣁⣀⣿⣿⣿⣇⣀⣿⣿⣿⣿⣀⣸⣿⣀⣸⣿⣿⣀⣈⣉⣀⣸⣿⣿⣀⣉⣉⣀⣿⣿⣇⣀⣿⣇⣀⣿⣿⣇⣀⣿⣇⣀⣿⣿⣇⣀⣿⣿⣿⣿⣿⣿⣿⣀⣿⣿


	TelegramAPIbot for stalcraft:x game
	This bot get emission info and send it for all users
	Bot show you quantity users
	Bot get promocodes from steam page`,
	Run: func(cmd *cobra.Command, args []string) {
		if startbot {
			fmt.Println("tgBot started")
			start.StartBot()
		}
		if startcrawler {
			fmt.Println("crawler started")
			start.StartCrawler()
		}
		if adminapi {
			fmt.Println("adminAPI started")
			api.StartAdminAPI()
		}
	},
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
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
	rootCmd.Flags().BoolVarP(&startcrawler, "crawler", "c", false, "start stalcraft API info handler")
	rootCmd.Flags().BoolVarP(&adminapi, "adminapi", "a", false, "start adminAPI users control tool")
}
