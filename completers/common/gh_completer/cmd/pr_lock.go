package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var pr_lockCmd = &cobra.Command{
	Use:     "lock {<number> | <url>}",
	Short:   "Lock pull request conversation",
	GroupID: "Targeted commands",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pr_lockCmd).Standalone()

	pr_lockCmd.Flags().StringP("reason", "r", "", "Optional reason for locking conversation (off_topic, resolved, spam, too_heated).")
	prCmd.AddCommand(pr_lockCmd)

	carapace.Gen(pr_lockCmd).FlagCompletion(carapace.ActionMap{
		"reason": carapace.ActionValues("off_topic", "resolved", "spam", "too_heated"),
	})

	carapace.Gen(pr_lockCmd).PositionalCompletion(
		action.ActionPullRequests(pr_lockCmd, action.PullRequestOpts{Open: true, Closed: true, Merged: true}),
	)
}
