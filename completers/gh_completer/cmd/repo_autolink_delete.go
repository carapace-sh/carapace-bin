package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var repo_autolink_deleteCmd = &cobra.Command{
	Use:   "delete <id>",
	Short: "Delete an autolink reference",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repo_autolink_deleteCmd).Standalone()

	repo_autolink_deleteCmd.Flags().Bool("yes", false, "Confirm deletion without prompting")
	repo_autolinkCmd.AddCommand(repo_autolink_deleteCmd)

	carapace.Gen(repo_autolink_deleteCmd).PositionalCompletion(
		action.ActionAutolinks(repo_autolink_deleteCmd),
	)
}
