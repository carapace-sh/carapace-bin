package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_worktree_removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a linked worktree",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_worktree_removeCmd).Standalone()

	help_worktreeCmd.AddCommand(help_worktree_removeCmd)
}
