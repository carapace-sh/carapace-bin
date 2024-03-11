package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cli_helpCmd = &cobra.Command{
	Use:   "help [COMMAND]...",
	Short: "Print this message or the help of the given subcommand(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cli_helpCmd).Standalone()

	cliCmd.AddCommand(cli_helpCmd)

	carapace.Gen(cli_helpCmd).PositionalAnyCompletion(
		carapace.ActionCommands(cli_helpCmd),
	)
}
