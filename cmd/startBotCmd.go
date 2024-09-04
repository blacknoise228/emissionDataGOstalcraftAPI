package cmd

import (
	"stalcraftBot/internal/startBot"

	"github.com/spf13/cobra"
)

var startTgBot = &cobra.Command{
	Use:   "startBot",
	Short: "start telegram bot",
	Long:  "this command starting all func for bot",
	Run: func(cmd *cobra.Command, args []string) {
		startBot.StartBot()
	},
}

func init() {
	rootCmd.AddCommand(startTgBot)
}
