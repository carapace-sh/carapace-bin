package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/but"
	"github.com/spf13/cobra"
)

var worktree_archiveCmd = &cobra.Command{
	Use:   "archive",
	Short: "Hide a worktree from the workspace",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(worktree_archiveCmd).Standalone()

	worktree_archiveCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	worktreeCmd.AddCommand(worktree_archiveCmd)

	carapace.Gen(worktree_archiveCmd).PositionalCompletion(
		but.ActionWorktrees(but.WorktreesOpts{}.Default()),
	)
}
