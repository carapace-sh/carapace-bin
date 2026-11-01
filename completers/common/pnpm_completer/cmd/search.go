package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
	Use:     "search",
	Short:   "Search for packages in the registry",
	Aliases: []string{"s", "se", "find"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(searchCmd).Standalone()

	searchCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	searchCmd.Flags().Bool("json", false, "Show search results in JSON format")
	searchCmd.Flags().String("registry", "", "Registry URL to search in")
	searchCmd.Flags().String("search-limit", "", "Maximum number of results to show (default: 20)")
	rootCmd.AddCommand(searchCmd)
}
