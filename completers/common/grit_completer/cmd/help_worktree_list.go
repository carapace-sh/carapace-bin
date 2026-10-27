package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_worktree_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List registered worktrees (main first, then linked)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_worktree_listCmd).Standalone()

	help_worktreeCmd.AddCommand(help_worktree_listCmd)
}
