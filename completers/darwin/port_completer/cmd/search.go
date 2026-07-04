package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search for ports by name or description",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(searchCmd).Standalone()
	searchCmd.Flags().Bool("case-sensitive", false, "Case-sensitive search")
	searchCmd.Flags().Bool("category", false, "Search in categories")
	searchCmd.Flags().Bool("description", false, "Search in descriptions")
	searchCmd.Flags().Bool("exact", false, "Match literal string exactly")
	searchCmd.Flags().Bool("glob", false, "Treat search string as glob pattern (default)")
	searchCmd.Flags().Bool("homepage", false, "Search in homepages")
	searchCmd.Flags().Bool("line", false, "Output in single-line format")
	searchCmd.Flags().Bool("maintainer", false, "Search by maintainer")
	searchCmd.Flags().Bool("name", false, "Search only port names")
	searchCmd.Flags().Bool("regex", false, "Treat search string as regular expression")
	rootCmd.AddCommand(searchCmd)
}
