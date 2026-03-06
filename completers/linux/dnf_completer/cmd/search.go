package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
	Use:     "search [options] <pattern>...",
	Aliases: []string{"se"},
	Short:   "search for packages using keywords",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(searchCmd).Standalone()

	searchCmd.Flags().Bool("all", false, "search in description and URL fields")
	searchCmd.Flags().Bool("name", false, "limit search to name field")
	searchCmd.Flags().Bool("summary", false, "limit search to summary field")
	searchCmd.Flags().Bool("showduplicates", false, "show all versions of packages")

	rootCmd.AddCommand(searchCmd)

	carapace.Gen(searchCmd).PositionalCompletion(
		carapace.ActionValues(), // TODO: complete with keywords
	)
}
