package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gh"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var mergeCmd = &cobra.Command{
	Use:   "merge [<pr-or-branch>]",
	Short: "Merge a stack of PRs",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mergeCmd).Standalone()

	rootCmd.AddCommand(mergeCmd)

	carapace.Gen(mergeCmd).PositionalCompletion(
		carapace.Batch(
			git.ActionLocalBranches(),
			gh.ActionPullRequests(gh.PullRequestOpts{}), // TODO needs to support implicit repo mapping
		).ToA(),
	)
}
