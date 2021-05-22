package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var issue_reopenCmd = &cobra.Command{
	Use:   "reopen {<number> | <url>}",
	Short: "Reopen issue",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	issueCmd.AddCommand(issue_reopenCmd)

	carapace.Gen(issue_reopenCmd).PositionalCompletion(
		action.ActionIssues(issue_reopenCmd, action.IssueOpts{Closed: true}),
	)
}
