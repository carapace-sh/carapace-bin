package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search for apps in the App Store",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(searchCmd).Standalone()
	searchCmd.Flags().Bool("json", false, "Output JSON")
	searchCmd.Flags().Bool("price", false, "Output the price of each app")
	rootCmd.AddCommand(searchCmd)
}
