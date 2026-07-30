package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gh"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var mergeCmd = &cobra.Command{
	Use:     "merge [<stack-number> | <pr-number>]",
	Short:   "Merge a stack of pull requests",
	GroupID: "remote",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mergeCmd).Standalone()

	mergeCmd.Flags().Bool("merge", false, "Merge with a merge commit")
	mergeCmd.Flags().String("merge-method", "", "Merge method to use: merge, squash, or rebase")
	mergeCmd.Flags().Bool("rebase", false, "Rebase and merge")
	mergeCmd.Flags().Bool("squash", false, "Squash and merge")
	mergeCmd.Flags().BoolP("yes", "y", false, "Merge without prompting for confirmation")
	rootCmd.AddCommand(mergeCmd)

	carapace.Gen(mergeCmd).FlagCompletion(carapace.ActionMap{
		"merge-method": carapace.ActionValues("merge", "squash", "rebase"),
	})

	carapace.Gen(mergeCmd).PositionalCompletion(
		carapace.Batch(
			git.ActionLocalBranches(),
			gh.ActionPullRequests(gh.PullRequestOpts{Open: true}),
		).ToA(),
	)
}
