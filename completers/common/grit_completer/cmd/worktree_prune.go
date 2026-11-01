package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var worktree_pruneCmd = &cobra.Command{
	Use:   "prune",
	Short: "Drop registrations for linked worktrees whose directories are gone",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(worktree_pruneCmd).Standalone()

	worktree_pruneCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	worktreeCmd.AddCommand(worktree_pruneCmd)
}
