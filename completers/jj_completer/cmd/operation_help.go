package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var operation_helpCmd = &cobra.Command{
	Use:   "help",
	Short: "Print this message or the help of the given subcommand(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(operation_helpCmd).Standalone()

	operationCmd.AddCommand(operation_helpCmd)

	carapace.Gen(operation_helpCmd).PositionalAnyCompletion(
		carapace.ActionCommands(operationCmd),
	)
}
