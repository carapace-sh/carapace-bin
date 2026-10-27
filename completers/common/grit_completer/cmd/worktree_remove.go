package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var worktree_removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a linked worktree",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(worktree_removeCmd).Standalone()

	worktree_removeCmd.Flags().BoolP("force", "f", false, "Remove even when the worktree has local changes or is locked")
	worktree_removeCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	worktreeCmd.AddCommand(worktree_removeCmd)

	carapace.Gen(worktree_removeCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
