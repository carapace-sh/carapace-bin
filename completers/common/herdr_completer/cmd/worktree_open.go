package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var worktree_openCmd = &cobra.Command{
	Use:   "open",
	Short: "Open an existing Git worktree",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(worktree_openCmd).Standalone()

	worktree_openCmd.Flags().String("branch", "", "")
	worktree_openCmd.Flags().String("cwd", "", "")
	worktree_openCmd.Flags().Bool("focus", false, "")
	worktree_openCmd.Flags().Bool("json", false, "")
	worktree_openCmd.Flags().String("label", "", "")
	worktree_openCmd.Flags().Bool("no-focus", false, "")
	worktree_openCmd.Flags().String("path", "", "")
	worktree_openCmd.Flags().String("workspace", "", "")
	worktreeCmd.AddCommand(worktree_openCmd)

	carapace.Gen(worktree_openCmd).FlagCompletion(carapace.ActionMap{
		"cwd":  carapace.ActionFiles(),
		"path": carapace.ActionFiles(),
	})
}
