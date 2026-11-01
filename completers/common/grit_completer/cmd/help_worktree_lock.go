package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_worktree_lockCmd = &cobra.Command{
	Use:   "lock",
	Short: "Mark a linked worktree as locked (prevents accidental removal)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_worktree_lockCmd).Standalone()

	help_worktreeCmd.AddCommand(help_worktree_lockCmd)
}
