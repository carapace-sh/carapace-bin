package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var notes_helpCmd = &cobra.Command{
	Use:   "help",
	Short: "Print this message or the help of the given subcommand(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(notes_helpCmd).Standalone()

	notesCmd.AddCommand(notes_helpCmd)

	carapace.Gen(notes_helpCmd).PositionalCompletion(
		carapace.ActionCommands(notesCmd),
	)

	carapace.Gen(notes_helpCmd).PositionalAnyCompletion(
		carapace.ActionCommands(notesCmd),
	)
}
