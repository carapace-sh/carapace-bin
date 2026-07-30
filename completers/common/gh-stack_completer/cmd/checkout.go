package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gh"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var checkoutCmd = &cobra.Command{
	Use:     "checkout [<stack-number> | <pr-number> | <pr-url> | <branch>]",
	Short:   "Checkout a stack by stack number, PR number, PR URL, or branch name",
	GroupID: "stack",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(checkoutCmd).Standalone()

	rootCmd.AddCommand(checkoutCmd)

	carapace.Gen(checkoutCmd).PositionalCompletion(
		carapace.Batch(
			git.ActionLocalBranches(),
			gh.ActionPullRequests(gh.PullRequestOpts{}.Default()),
		).ToA(),
	)
}
