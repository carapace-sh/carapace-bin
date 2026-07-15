package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var worktree_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List worktree workspaces",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(worktree_listCmd).Standalone()

	worktree_listCmd.Flags().String("cwd", "", "")
	worktree_listCmd.Flags().Bool("json", false, "")
	worktree_listCmd.Flags().String("workspace", "", "")
	worktreeCmd.AddCommand(worktree_listCmd)

	carapace.Gen(worktree_listCmd).FlagCompletion(carapace.ActionMap{
		"cwd": carapace.ActionFiles(),
	})
}
