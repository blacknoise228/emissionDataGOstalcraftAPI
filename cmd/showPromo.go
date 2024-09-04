package cmd

import (
	"stalcraftBot/internal/getData"

	"github.com/spf13/cobra"
)

var promo = &cobra.Command{
	Use:   "promo",
	Short: "printing promocodes",
	Long:  "parse websyte and print actual promocodes",
	Run: func(cmd *cobra.Command, args []string) {
		getData.ParseFunc()
	},
}

func init() {
	rootCmd.AddCommand(promo)
}
