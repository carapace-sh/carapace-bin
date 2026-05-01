package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search packages by querying search.nixos.org",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(searchCmd).Standalone()

	searchCmd.Flags().StringP("channel", "c", "nixos-unstable", "Name of the channel to query")
	searchCmd.Flags().BoolP("json", "j", false, "Output results as JSON")
	searchCmd.Flags().StringP("limit", "l", "30", "Number of search results to display")
	searchCmd.Flags().BoolP("platforms", "P", false, "Show supported platforms for each package")

	rootCmd.AddCommand(searchCmd)
}
