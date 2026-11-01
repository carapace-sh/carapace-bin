package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var worktree_help_removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a linked worktree",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(worktree_help_removeCmd).Standalone()

	worktree_helpCmd.AddCommand(worktree_help_removeCmd)
}
