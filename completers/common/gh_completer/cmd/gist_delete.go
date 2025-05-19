package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var gist_deleteCmd = &cobra.Command{
	Use:   "delete {<id> | <url>}",
	Short: "Delete a gist",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gist_deleteCmd).Standalone()

	gist_deleteCmd.Flags().Bool("yes", false, "Confirm deletion without prompting")
	gistCmd.AddCommand(gist_deleteCmd)

	carapace.Gen(gist_deleteCmd).PositionalCompletion(
		action.ActionGists(gist_deleteCmd),
	)
}
