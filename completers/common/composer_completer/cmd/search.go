package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Searches for packages",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(searchCmd).Standalone()

	searchCmd.Flags().StringP("format", "f", "", "Format of the output: text or json")
	searchCmd.Flags().BoolP("only-name", "N", false, "Search only in package names")
	searchCmd.Flags().BoolP("only-vendor", "O", false, "Search only for vendor / organization names, returns only \"vendor\" as result")
	searchCmd.Flags().StringP("type", "t", "", "Search for a specific package type")

	rootCmd.AddCommand(searchCmd)

	carapace.Gen(searchCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "text"),
	})
}
