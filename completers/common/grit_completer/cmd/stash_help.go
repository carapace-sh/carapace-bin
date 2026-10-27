package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var stash_helpCmd = &cobra.Command{
	Use:   "help",
	Short: "Print this message or the help of the given subcommand(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stash_helpCmd).Standalone()

	stashCmd.AddCommand(stash_helpCmd)

	carapace.Gen(stash_helpCmd).PositionalCompletion(
		carapace.ActionCommands(stashCmd),
	)

	carapace.Gen(stash_helpCmd).PositionalAnyCompletion(
		carapace.ActionCommands(stashCmd),
	)
}
