package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/grit_completer/cmd/action"
	"github.com/spf13/cobra"
)

var stash_popCmd = &cobra.Command{
	Use:   "pop",
	Short: "Apply a stash entry and remove it when there are no conflicts",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stash_popCmd).Standalone()

	stash_popCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	stash_popCmd.Flags().Bool("index", false, "Restore the stashed index as well as the worktree")
	stashCmd.AddCommand(stash_popCmd)

	carapace.Gen(stash_popCmd).PositionalCompletion(
		action.ActionStashIndexes(),
	)
}
