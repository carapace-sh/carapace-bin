package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inputs_helpCmd = &cobra.Command{
	Use:   "help",
	Short: "Print this message or the help of the given subcommand(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inputs_helpCmd).Standalone()

	inputsCmd.AddCommand(inputs_helpCmd)

	carapace.Gen(inputs_helpCmd).PositionalAnyCompletion(
		carapace.ActionCommands(inputsCmd),
	)
}
