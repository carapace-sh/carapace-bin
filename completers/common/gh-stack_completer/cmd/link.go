package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gh"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var linkCmd = &cobra.Command{
	Use:   "link <branch-or-pr> <branch-or-pr> [<branch-or-pr>...]",
	Short: "Link PRs into a stack on GitHub without local tracking",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(linkCmd).Standalone()

	linkCmd.Flags().String("base", "", "Base branch for the bottom of the stack")
	linkCmd.Flags().Bool("draft", false, "Create new PRs as drafts")
	linkCmd.Flags().String("remote", "", "Remote to push to (defaults to auto-detected remote)")
	rootCmd.AddCommand(linkCmd)

	carapace.Gen(linkCmd).FlagCompletion(carapace.ActionMap{
		"base":   git.ActionLocalBranches(),
		"remote": git.ActionRemotes(),
	})

	carapace.Gen(linkCmd).PositionalAnyCompletion(
		carapace.Batch(
			git.ActionLocalBranches(),
			gh.ActionPullRequests(gh.PullRequestOpts{}), // TODO needs to support implicit repo mapping
		).ToA().FilterArgs(),
	)
}
