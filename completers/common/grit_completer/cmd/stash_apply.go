package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/grit_completer/cmd/action"
	"github.com/spf13/cobra"
)

var stash_applyCmd = &cobra.Command{
	Use:   "apply",
	Short: "Apply a stash entry without removing it",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stash_applyCmd).Standalone()

	stash_applyCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	stash_applyCmd.Flags().Bool("index", false, "Restore the stashed index as well as the worktree")
	stashCmd.AddCommand(stash_applyCmd)

	carapace.Gen(stash_applyCmd).PositionalCompletion(
		action.ActionStashIndexes(),
	)
}
