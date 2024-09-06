package cmd

import (
	"fmt"
	"stalcraftBot/internal/jSon"
	"stalcraftBot/pkg/api"
	"stalcraftBot/pkg/getData"

	"github.com/spf13/cobra"
)

var quantityUsers = &cobra.Command{
	Use:   "users",
	Short: "users quantity",
	Long:  "returned quantity of users",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Quantity users: ", jSon.QuantityUsers())
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
	Short: "start adminAPI for administrative controls",
	Run: func(cmd *cobra.Command, args []string) {
		api.StartAdminAPI()
	},
}

func init() {
	rootCmd.AddCommand(quantityUsers)
	rootCmd.AddCommand(promo)
	rootCmd.AddCommand(adminapi)
}
