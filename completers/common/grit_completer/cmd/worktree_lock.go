package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/grit_completer/cmd/action"
	"github.com/spf13/cobra"
)

var worktree_lockCmd = &cobra.Command{
	Use:   "lock",
	Short: "Mark a linked worktree as locked (prevents accidental removal)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(worktree_lockCmd).Standalone()

	worktree_lockCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	worktree_lockCmd.Flags().String("reason", "", "Optional reason stored in the lock file")
	worktreeCmd.AddCommand(worktree_lockCmd)

	carapace.Gen(worktree_lockCmd).PositionalCompletion(
		action.ActionWorktreePaths(),
	)
}
