package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gc_helpCmd = &cobra.Command{
	Use:   "help",
	Short: "Print this message or the help of the given subcommand(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gc_helpCmd).Standalone()

	gcCmd.AddCommand(gc_helpCmd)

	carapace.Gen(gc_helpCmd).PositionalCompletion(
		carapace.ActionCommands(gcCmd),
	)

	carapace.Gen(gc_helpCmd).PositionalAnyCompletion(
		carapace.ActionCommands(gcCmd),
	)
}
