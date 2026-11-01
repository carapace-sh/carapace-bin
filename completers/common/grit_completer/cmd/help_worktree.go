package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_worktreeCmd = &cobra.Command{
	Use:   "worktree",
	Short: "List, add, or remove linked worktrees",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_worktreeCmd).Standalone()

	helpCmd.AddCommand(help_worktreeCmd)
}
