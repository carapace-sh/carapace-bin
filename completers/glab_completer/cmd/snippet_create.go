package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var snippet_createCmd = &cobra.Command{
	Use:     "create [flags] -t <title> <file1> [<file2>...]",
	Short:   "Create a new snippet.",
	Aliases: []string{"new"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(snippet_createCmd).Standalone()

	snippet_createCmd.Flags().StringP("description", "d", "", "Description of the snippet.")
	snippet_createCmd.Flags().StringP("filename", "f", "", "Filename of the snippet in GitLab.")
	snippet_createCmd.Flags().BoolP("personal", "p", false, "Create a personal snippet.")
	snippet_createCmd.Flags().StringP("title", "t", "", "(required) Title of the snippet.")
	snippet_createCmd.Flags().StringP("visibility", "v", "", "Limit by visibility: 'public', 'internal', or 'private'")
	snippetCmd.AddCommand(snippet_createCmd)

	carapace.Gen(snippet_createCmd).FlagCompletion(carapace.ActionMap{
		"visibility": carapace.ActionValues("public", "internal", "private").StyleF(style.ForKeyword),
	})

	carapace.Gen(snippet_createCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
