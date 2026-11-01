package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
	Use:   "search [options] <patterns>...",
	Short: "search for software matching all specified strings",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(searchCmd).Standalone()

	searchCmd.Flags().Bool("all", false, "Search also package description and URL")
	searchCmd.Flags().Bool("name", false, "Limit the search to the Name field")
	searchCmd.Flags().Bool("showduplicates", false, "Show all versions of the packages")
	searchCmd.Flags().Bool("summary", false, "Limit the search to the Summary field")

	rootCmd.AddCommand(searchCmd)
}
