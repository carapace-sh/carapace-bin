package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/but"
	"github.com/spf13/cobra"
)

var worktree_unarchiveCmd = &cobra.Command{
	Use:   "unarchive",
	Short: "Show an archived worktree in the workspace again",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(worktree_unarchiveCmd).Standalone()

	worktree_unarchiveCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	worktreeCmd.AddCommand(worktree_unarchiveCmd)

	carapace.Gen(worktree_unarchiveCmd).PositionalCompletion(
		but.ActionWorktrees(but.WorktreesOpts{}.Default()),
	)
}
