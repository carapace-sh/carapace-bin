package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_worktree_pruneCmd = &cobra.Command{
	Use:   "prune",
	Short: "Drop registrations for linked worktrees whose directories are gone",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_worktree_pruneCmd).Standalone()

	help_worktreeCmd.AddCommand(help_worktree_pruneCmd)
}
