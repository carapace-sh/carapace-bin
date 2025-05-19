package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var issue_subscribeCmd = &cobra.Command{
	Use:     "subscribe <id>",
	Short:   "Subscribe to an issue.",
	Aliases: []string{"sub"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(issue_subscribeCmd).Standalone()

	issueCmd.AddCommand(issue_subscribeCmd)

	carapace.Gen(issue_subscribeCmd).PositionalCompletion(
		action.ActionIssues(issue_subscribeCmd, "opened"),
	)
}
