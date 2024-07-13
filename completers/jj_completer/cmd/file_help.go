package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var file_helpCmd = &cobra.Command{
	Use:   "help",
	Short: "Print this message or the help of the given subcommand(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(file_helpCmd).Standalone()

	fileCmd.AddCommand(file_helpCmd)

	carapace.Gen(file_helpCmd).PositionalAnyCompletion(
		carapace.ActionCommands(utilCmd),
	)
}
