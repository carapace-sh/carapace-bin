package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var worktree_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create and open a Git worktree",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(worktree_createCmd).Standalone()

	worktree_createCmd.Flags().String("base", "", "")
	worktree_createCmd.Flags().String("branch", "", "")
	worktree_createCmd.Flags().String("cwd", "", "")
	worktree_createCmd.Flags().Bool("focus", false, "")
	worktree_createCmd.Flags().Bool("json", false, "")
	worktree_createCmd.Flags().String("label", "", "")
	worktree_createCmd.Flags().Bool("no-focus", false, "")
	worktree_createCmd.Flags().String("path", "", "")
	worktree_createCmd.Flags().String("workspace", "", "")
	worktreeCmd.AddCommand(worktree_createCmd)

	carapace.Gen(worktree_createCmd).FlagCompletion(carapace.ActionMap{
		"cwd":  carapace.ActionFiles(),
		"path": carapace.ActionFiles(),
	})
}
