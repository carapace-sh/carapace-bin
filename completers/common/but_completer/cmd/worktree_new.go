package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/but"
	"github.com/spf13/cobra"
)

var worktree_newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create a worktree on a new branch at the workspace base or above a commit, or on an existing branch that isn't in the workspace",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(worktree_newCmd).Standalone()

	worktree_newCmd.Flags().StringP("above", "A", "", "Start the new branch at COMMIT instead of the workspace base")
	worktree_newCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	worktreeCmd.AddCommand(worktree_newCmd)

	carapace.Gen(worktree_newCmd).FlagCompletion(carapace.ActionMap{
		"above": but.ActionCliIds(but.CliIdsOpts{Commits: true}),
	})

	carapace.Gen(worktree_newCmd).PositionalCompletion(
		but.ActionLocalBranches(),
	)
}
