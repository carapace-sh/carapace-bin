package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var hooks_helpCmd = &cobra.Command{
	Use:   "help",
	Short: "Print this message or the help of the given subcommand(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(hooks_helpCmd).Standalone()

	hooksCmd.AddCommand(hooks_helpCmd)

	carapace.Gen(hooks_helpCmd).PositionalCompletion(
		carapace.ActionCommands(hooksCmd),
	)

	carapace.Gen(hooks_helpCmd).PositionalAnyCompletion(
		carapace.ActionCommands(hooksCmd),
	)
}
