package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var issue_unsubscribeCmd = &cobra.Command{
	Use:     "unsubscribe <id>",
	Short:   "Unsubscribe from an issue.",
	Aliases: []string{"unsub"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(issue_unsubscribeCmd).Standalone()

	issueCmd.AddCommand(issue_unsubscribeCmd)

	carapace.Gen(issue_unsubscribeCmd).PositionalCompletion(
		action.ActionIssues(issue_unsubscribeCmd, "opened"),
	)
}
