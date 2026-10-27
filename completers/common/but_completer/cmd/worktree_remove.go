package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/but"
	"github.com/spf13/cobra"
)

var worktree_removeCmd = &cobra.Command{
	Use:     "remove",
	Short:   "Remove a worktree from disk, like git worktree remove",
	Aliases: []string{"rm"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(worktree_removeCmd).Standalone()

	worktree_removeCmd.Flags().BoolP("force", "f", false, "Remove the worktree even if it has uncommitted changes")
	worktree_removeCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	worktreeCmd.AddCommand(worktree_removeCmd)

	carapace.Gen(worktree_removeCmd).PositionalCompletion(
		but.ActionWorktrees(but.WorktreesOpts{}.Default()),
	)
}
