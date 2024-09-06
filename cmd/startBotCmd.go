package cmd

import (
	"fmt"
	"stalcraftBot/internal/getData"
	"stalcraftBot/internal/startBot"
	"stalcraftBot/internal/tgBot"
	"stalcraftBot/pkg/api"

	"github.com/spf13/cobra"
)

var startTgBot = &cobra.Command{
	Use:   "startbot",
	Short: "start telegram bot",
	Long:  "this command starting all func for bot",
	Run: func(cmd *cobra.Command, args []string) {
		startBot.StartBot()
	},
}
var quantityUsers = &cobra.Command{
	Use:   "users",
	Short: "users quantity",
	Long:  "returned quantity of users",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(tgBot.QuantityUsers())
	},
}
var promo = &cobra.Command{
	Use:   "promo",
	Short: "printing promocodes",
	Long:  "parse websyte and print actual promocodes",
	Run: func(cmd *cobra.Command, args []string) {
		getData.ParseFunc()
	},
}
var adminapi = &cobra.Command{
	Use:   "adminapi",
	Short: "start admin API",
	Long:  "starting administrative API service",
	Run: func(cmd *cobra.Command, args []string) {
		api.StartAdminAPI()
	},
}

func init() {
	rootCmd.AddCommand(startTgBot)
	rootCmd.AddCommand(quantityUsers)
	rootCmd.AddCommand(promo)
	rootCmd.AddCommand(adminapi)
}
