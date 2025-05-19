package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var gist_editCmd = &cobra.Command{
	Use:   "edit {<id> | <url>} [<filename>]",
	Short: "Edit one of your gists",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gist_editCmd).Standalone()

	gist_editCmd.Flags().StringP("add", "a", "", "Add a new file to the gist")
	gist_editCmd.Flags().StringP("desc", "d", "", "New description for the gist")
	gist_editCmd.Flags().StringP("filename", "f", "", "Select a file to edit")
	gist_editCmd.Flags().StringP("remove", "r", "", "Remove a file from the gist")
	gistCmd.AddCommand(gist_editCmd)

	carapace.Gen(gist_editCmd).FlagCompletion(carapace.ActionMap{
		"add": carapace.ActionFiles(),
		"filename": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if len(c.Args) > 0 {
				return action.ActionGistFiles(gist_editCmd, c.Args[0])
			} else {
				return carapace.ActionValues()
			}
		}),
	})

	carapace.Gen(gist_editCmd).PositionalCompletion(
		action.ActionGists(gist_editCmd),
	)
}
