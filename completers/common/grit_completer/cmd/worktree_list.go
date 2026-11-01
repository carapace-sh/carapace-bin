package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var worktree_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List registered worktrees (main first, then linked)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(worktree_listCmd).Standalone()

	worktree_listCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	worktreeCmd.AddCommand(worktree_listCmd)
}
