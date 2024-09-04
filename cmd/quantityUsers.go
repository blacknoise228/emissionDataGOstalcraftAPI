package cmd

import (
	"stalcraftBot/internal/tgBot"

	"github.com/spf13/cobra"
)

var quantityUsers = &cobra.Command{
	Use:   "users",
	Short: "users quantity",
	Long:  "returned quantity of users",
	Run: func(cmd *cobra.Command, args []string) {
		tgBot.QuantityUsers()
	},
}

func init() {
	rootCmd.AddCommand(quantityUsers)
}
