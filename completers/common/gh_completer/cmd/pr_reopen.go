package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var pr_reopenCmd = &cobra.Command{
	Use:     "reopen {<number> | <url> | <branch>}",
	Short:   "Reopen a pull request",
	GroupID: "Targeted commands",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pr_reopenCmd).Standalone()

	pr_reopenCmd.Flags().StringP("comment", "c", "", "Add a reopening comment")
	prCmd.AddCommand(pr_reopenCmd)

	carapace.Gen(pr_reopenCmd).FlagCompletion(carapace.ActionMap{
		"comment": action.ActionBody(pr_reopenCmd),
	})

	carapace.Gen(pr_reopenCmd).PositionalCompletion(
		action.ActionPullRequests(pr_reopenCmd, action.PullRequestOpts{Closed: true}),
	)
}
