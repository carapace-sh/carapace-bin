package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
	Use:     "search",
	Short:   "Search packages",
	Aliases: []string{"sr"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(searchCmd).Standalone()

	searchCmd.Flags().StringP("language", "l", "", "Summary and description language")
	searchCmd.Flags().StringP("repository", "r", "", "Name of the repository to search")
	searchCmd.Flags().Bool("installdb", false, "Search in installdb")
	searchCmd.Flags().Bool("sourcedb", false, "Search in sourcedb")
	searchCmd.Flags().Bool("name", false, "Search in package name")
	searchCmd.Flags().Bool("summary", false, "Search in package summary")
	searchCmd.Flags().Bool("description", false, "Search in package description")

	rootCmd.AddCommand(searchCmd)
}
