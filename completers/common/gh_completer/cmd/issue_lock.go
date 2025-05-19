package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var issue_lockCmd = &cobra.Command{
	Use:     "lock {<number> | <url>}",
	Short:   "Lock issue conversation",
	GroupID: "Targeted commands",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(issue_lockCmd).Standalone()

	issue_lockCmd.Flags().StringP("reason", "r", "", "Optional reason for locking conversation (off_topic, resolved, spam, too_heated).")
	issueCmd.AddCommand(issue_lockCmd)

	carapace.Gen(issue_lockCmd).FlagCompletion(carapace.ActionMap{
		"reason": carapace.ActionValues("off_topic", "resolved", "spam", "too_heated"),
	})

	carapace.Gen(issue_lockCmd).PositionalCompletion(
		action.ActionIssues(issue_lockCmd, action.IssueOpts{Open: true, Closed: true}),
	)
}
