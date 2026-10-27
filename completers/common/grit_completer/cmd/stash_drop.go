package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/grit_completer/cmd/action"
	"github.com/spf13/cobra"
)

var stash_dropCmd = &cobra.Command{
	Use:   "drop",
	Short: "Remove a stash entry without applying it",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stash_dropCmd).Standalone()

	stash_dropCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	stashCmd.AddCommand(stash_dropCmd)

	carapace.Gen(stash_dropCmd).PositionalCompletion(
		action.ActionStashIndexes(),
	)
}
