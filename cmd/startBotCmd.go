package cmd

import (
	"fmt"
	"stalcraftBot/internal/jsWorker"
	"stalcraftBot/pkg/getData"

	"github.com/spf13/cobra"
)

var quantityUsers = &cobra.Command{
	Use:   "users",
	Short: "users quantity",
	Long:  "returned quantity of users",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Quantity users: ", jsWorker.QuantityUsers())
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

func init() {
	rootCmd.AddCommand(quantityUsers)
	rootCmd.AddCommand(promo)
}
