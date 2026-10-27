package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/grit_completer/cmd/action"
	"github.com/spf13/cobra"
)

var worktree_unlockCmd = &cobra.Command{
	Use:   "unlock",
	Short: "Clear the lock on a linked worktree",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(worktree_unlockCmd).Standalone()

	worktree_unlockCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	worktreeCmd.AddCommand(worktree_unlockCmd)

	carapace.Gen(worktree_unlockCmd).PositionalCompletion(
		action.ActionWorktreePaths(),
	)
}
