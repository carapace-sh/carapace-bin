package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var worktree_help_pruneCmd = &cobra.Command{
	Use:   "prune",
	Short: "Drop registrations for linked worktrees whose directories are gone",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(worktree_help_pruneCmd).Standalone()

	worktree_helpCmd.AddCommand(worktree_help_pruneCmd)
}
