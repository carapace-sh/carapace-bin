package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/eopkg"
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
	searchCmd.Flags().StringP("repository", "r", "", "name of the source or package repository")
	searchCmd.Flags().Bool("summary", false, "only search in the summary field of packages")

	rootCmd.AddCommand(searchCmd)

	carapace.Gen(searchCmd).FlagCompletion(carapace.ActionMap{
		"repository": eopkg.ActionRepositories(),
	})
}
