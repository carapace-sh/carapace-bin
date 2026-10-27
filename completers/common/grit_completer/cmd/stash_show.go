package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/grit_completer/cmd/action"
	"github.com/spf13/cobra"
)

var stash_showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show the summary or patch for a stash entry",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stash_showCmd).Standalone()

	stash_showCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	stash_showCmd.Flags().BoolP("patch", "p", false, "Show the full patch instead of a diffstat")
	stashCmd.AddCommand(stash_showCmd)

	carapace.Gen(stash_showCmd).PositionalCompletion(
		action.ActionStashIndexes(),
	)
}
