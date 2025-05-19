package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var pr_updateBranchCmd = &cobra.Command{
	Use:     "update-branch [<number> | <url> | <branch>]",
	Short:   "Update a pull request branch",
	GroupID: "Targeted commands",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pr_updateBranchCmd).Standalone()

	pr_updateBranchCmd.Flags().Bool("rebase", false, "Update PR branch by rebasing on top of latest base branch")
	prCmd.AddCommand(pr_updateBranchCmd)

	carapace.Gen(pr_updateBranchCmd).PositionalCompletion(
		action.ActionPullRequests(pr_updateBranchCmd, action.PullRequestOpts{Open: true}),
	)
}
