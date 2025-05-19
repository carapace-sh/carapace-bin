package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var gist_renameCmd = &cobra.Command{
	Use:   "rename {<id> | <url>} <old-filename> <new-filename>",
	Short: "Rename a file in a gist",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gist_renameCmd).Standalone()

	gistCmd.AddCommand(gist_renameCmd)

	carapace.Gen(gist_renameCmd).PositionalCompletion(
		action.ActionGists(gist_renameCmd),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionGistFiles(gist_renameCmd, c.Args[0])
		}),
	)
}
