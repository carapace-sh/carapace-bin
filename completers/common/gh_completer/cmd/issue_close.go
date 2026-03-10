package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var issue_closeCmd = &cobra.Command{
	Use:     "close {<number> | <url>}",
	Short:   "Close issue",
	GroupID: "Targeted commands",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(issue_closeCmd).Standalone()

	issue_closeCmd.Flags().StringP("comment", "c", "", "Leave a closing comment")
	issue_closeCmd.Flags().String("duplicate-of", "", "Mark as duplicate of another issue by number or URL")
	issue_closeCmd.Flags().StringP("reason", "r", "", "Reason for closing: {completed|not planned|duplicate}")
	issueCmd.AddCommand(issue_closeCmd)

	carapace.Gen(issue_closeCmd).FlagCompletion(carapace.ActionMap{
		"comment":      action.ActionBody(issue_closeCmd),
		"duplicate-of": action.ActionIssues(issue_closeCmd, action.IssueOpts{Closed: true, Open: true}),
		"reason":       carapace.ActionValues("completed", "not planned", "duplicate"),
	})

	carapace.Gen(issue_closeCmd).PositionalCompletion(
		action.ActionIssues(issue_closeCmd, action.IssueOpts{Open: true}),
	)
}
