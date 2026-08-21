package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
	Use:     "search <term>",
	Aliases: []string{"sr"},
	Short:   "finds packages using the specified search term, which can be a regular expression when quoted",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(searchCmd).Standalone()

	searchCmd.Flags().Bool("description", false, "only search in the description field of packages")
	searchCmd.Flags().BoolP("installdb", "i", false, "only search installed packages, ignoring repository candidates")
	searchCmd.Flags().StringP("language", "l", "", "only search for summaries/descriptions with the matching language code")
	searchCmd.Flags().Bool("name", false, "only search in the name field of packages")
	searchCmd.Flags().StringP("repository", "r", "", "only search within the specified repository")
	searchCmd.Flags().BoolP("sourcedb", "s", false, "only search source repositories")
	searchCmd.Flags().Bool("summary", false, "only search in the summary field of packages")

	rootCmd.AddCommand(searchCmd)
}
