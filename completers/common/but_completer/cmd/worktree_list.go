package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var worktree_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List worktrees, most recently updated first",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(worktree_listCmd).Standalone()

	worktree_listCmd.Flags().Bool("active", false, "List all active worktrees")
	worktree_listCmd.Flags().Bool("archived", false, "List all archived worktrees")
	worktree_listCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	worktreeCmd.AddCommand(worktree_listCmd)
}
