package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var pr_revertCmd = &cobra.Command{
	Use:     "revert {<number> | <url> | <branch>}",
	Short:   "Revert a pull request",
	GroupID: "Targeted commands",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pr_revertCmd).Standalone()

	pr_revertCmd.Flags().StringP("body", "b", "", "Body for the revert pull request")
	pr_revertCmd.Flags().StringP("body-file", "F", "", "Read body text from `file` (use \"-\" to read from standard input)")
	pr_revertCmd.Flags().BoolP("draft", "d", false, "Mark revert pull request as a draft")
	pr_revertCmd.Flags().StringP("title", "t", "", "Title for the revert pull request")
	prCmd.AddCommand(pr_revertCmd)

	carapace.Gen(pr_revertCmd).FlagCompletion(carapace.ActionMap{
		"body-file": carapace.ActionFiles(),
	})

	carapace.Gen(pr_revertCmd).PositionalCompletion(
		action.ActionPullRequests(pr_revertCmd, action.PullRequestOpts{Merged: true}),
	)

}
