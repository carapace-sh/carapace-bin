package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var worktree_help_lockCmd = &cobra.Command{
	Use:   "lock",
	Short: "Mark a linked worktree as locked (prevents accidental removal)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(worktree_help_lockCmd).Standalone()

	worktree_helpCmd.AddCommand(worktree_help_lockCmd)
}
