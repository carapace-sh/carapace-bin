package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var snippet_createCmd = &cobra.Command{
	Use:     "create [path]",
	Short:   "Create new snippet",
	Aliases: []string{"new"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(snippet_createCmd).Standalone()
	snippet_createCmd.Flags().StringP("description", "d", "", "Description of the snippet")
	snippet_createCmd.Flags().StringP("filename", "f", "", "Filename of the snippet in GitLab")
	snippet_createCmd.Flags().StringP("title", "t", "", "Title of the snippet")
	snippet_createCmd.Flags().StringP("visibility", "v", "private", "Limit by visibility {public, internal, or private}")
	snippetCmd.AddCommand(snippet_createCmd)

	carapace.Gen(snippet_createCmd).FlagCompletion(carapace.ActionMap{
		"visibility": carapace.ActionValues("public", "internal", "private"),
	})

	carapace.Gen(snippet_createCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
