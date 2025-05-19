package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var issue_reopenCmd = &cobra.Command{
	Use:     "reopen [<id> | <url>] [flags]",
	Short:   "Reopen a closed issue.",
	Aliases: []string{"open"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(issue_reopenCmd).Standalone()

	issueCmd.AddCommand(issue_reopenCmd)

	carapace.Gen(issue_reopenCmd).PositionalCompletion(
		action.ActionIssues(issue_reopenCmd, "closed"),
	)
}
