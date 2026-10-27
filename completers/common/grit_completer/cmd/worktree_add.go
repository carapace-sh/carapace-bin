package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var worktree_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Create a linked worktree at `path`",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(worktree_addCmd).Standalone()

	worktree_addCmd.Flags().StringP("branch", "b", "", "Create and check out a new branch")
	worktree_addCmd.Flags().Bool("detach", false, "Check out in detached HEAD mode")
	worktree_addCmd.Flags().BoolP("force", "f", false, "Reuse a stale registration when the working tree directory is missing")
	worktree_addCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	worktreeCmd.AddCommand(worktree_addCmd)

	carapace.Gen(worktree_addCmd).FlagCompletion(carapace.ActionMap{
		"branch": git.ActionRefs(git.RefOption{LocalBranches: true}),
	})

	carapace.Gen(worktree_addCmd).PositionalCompletion(
		carapace.ActionDirectories(),
		git.ActionRefs(git.RefOption{}.Default()),
	)
}
