package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var issue_deleteCmd = &cobra.Command{
	Use:     "delete <id>",
	Short:   "Delete an issue.",
	Aliases: []string{"del"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(issue_deleteCmd).Standalone()

	issueCmd.AddCommand(issue_deleteCmd)

	carapace.Gen(issue_deleteCmd).PositionalCompletion(
		action.ActionIssues(issue_deleteCmd, "all"),
	)
}
