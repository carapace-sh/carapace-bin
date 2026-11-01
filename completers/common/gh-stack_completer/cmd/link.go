package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gh"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var linkCmd = &cobra.Command{
	Use:     "link <stack-number | branch-or-pr> <branch-or-pr> [<branch-or-pr>...]",
	Short:   "Link PRs into a stack on GitHub without local tracking",
	GroupID: "remote",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(linkCmd).Standalone()

	linkCmd.Flags().String("base", "", "Base branch for the bottom of the stack (defaults to the repository default branch)")
	linkCmd.Flags().Bool("open", false, "Mark new and existing PRs as ready for review")
	linkCmd.Flags().String("remote", "", "Remote to push to (defaults to auto-detected remote)")
	rootCmd.AddCommand(linkCmd)

	carapace.Gen(linkCmd).FlagCompletion(carapace.ActionMap{
		"base":   git.ActionLocalBranches(),
		"remote": git.ActionRemotes(),
	})

	carapace.Gen(linkCmd).PositionalAnyCompletion(
		carapace.Batch(
			git.ActionLocalBranches(),
			gh.ActionPullRequests(gh.PullRequestOpts{}.Default()),
		).ToA().FilterArgs(),
	)
}
