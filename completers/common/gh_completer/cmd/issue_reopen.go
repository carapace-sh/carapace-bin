package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var issue_reopenCmd = &cobra.Command{
	Use:     "reopen {<number> | <url>}",
	Short:   "Reopen issue",
	GroupID: "Targeted commands",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(issue_reopenCmd).Standalone()

	issue_reopenCmd.Flags().StringP("comment", "c", "", "Add a reopening comment")
	issueCmd.AddCommand(issue_reopenCmd)

	carapace.Gen(issue_reopenCmd).FlagCompletion(carapace.ActionMap{
		"comment": action.ActionBody(issue_reopenCmd),
	})

	carapace.Gen(issue_reopenCmd).PositionalCompletion(
		action.ActionIssues(issue_reopenCmd, action.IssueOpts{Closed: true}),
	)
}
