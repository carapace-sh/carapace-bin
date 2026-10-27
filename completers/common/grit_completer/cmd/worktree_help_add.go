package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var worktree_help_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Create a linked worktree at `path`",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(worktree_help_addCmd).Standalone()

	worktree_helpCmd.AddCommand(worktree_help_addCmd)
}
