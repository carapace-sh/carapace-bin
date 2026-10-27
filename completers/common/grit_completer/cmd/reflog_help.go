package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var reflog_helpCmd = &cobra.Command{
	Use:   "help",
	Short: "Print this message or the help of the given subcommand(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(reflog_helpCmd).Standalone()

	reflogCmd.AddCommand(reflog_helpCmd)

	carapace.Gen(reflog_helpCmd).PositionalCompletion(
		carapace.ActionCommands(reflogCmd),
	)

	carapace.Gen(reflog_helpCmd).PositionalAnyCompletion(
		carapace.ActionCommands(reflogCmd),
	)
}
