package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gh"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var checkoutCmd = &cobra.Command{
	Use:   "checkout [<pr-number> | <branch>]",
	Short: "Checkout a stack from a PR number or branch name",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(checkoutCmd).Standalone()

	rootCmd.AddCommand(checkoutCmd)

	carapace.Gen(checkoutCmd).PositionalCompletion(
		carapace.Batch(
			git.ActionLocalBranches(),
			gh.ActionPullRequests(gh.PullRequestOpts{}), // TODO needs to support implicit repo mapping
		).ToA(),
	)
}
