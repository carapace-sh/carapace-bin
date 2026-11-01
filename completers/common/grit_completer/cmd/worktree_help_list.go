package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var worktree_help_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List registered worktrees (main first, then linked)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(worktree_help_listCmd).Standalone()

	worktree_helpCmd.AddCommand(worktree_help_listCmd)
}
