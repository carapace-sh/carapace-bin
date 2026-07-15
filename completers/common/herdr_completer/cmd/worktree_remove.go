package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var worktree_removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a worktree checkout",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(worktree_removeCmd).Standalone()

	worktree_removeCmd.Flags().Bool("force", false, "")
	worktree_removeCmd.Flags().Bool("json", false, "")
	worktree_removeCmd.Flags().String("workspace", "", "")
	worktreeCmd.AddCommand(worktree_removeCmd)
}
