package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_worktree_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Create a linked worktree at `path`",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_worktree_addCmd).Standalone()

	help_worktreeCmd.AddCommand(help_worktree_addCmd)
}
