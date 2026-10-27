package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ignore_helpCmd = &cobra.Command{
	Use:   "help",
	Short: "Print this message or the help of the given subcommand(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ignore_helpCmd).Standalone()

	ignoreCmd.AddCommand(ignore_helpCmd)

	carapace.Gen(ignore_helpCmd).PositionalCompletion(
		carapace.ActionCommands(ignoreCmd),
	)

	carapace.Gen(ignore_helpCmd).PositionalAnyCompletion(
		carapace.ActionCommands(ignoreCmd),
	)
}
