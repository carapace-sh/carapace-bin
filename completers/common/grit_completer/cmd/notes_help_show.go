package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var notes_help_showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show the note for a revision",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(notes_help_showCmd).Standalone()

	notes_helpCmd.AddCommand(notes_help_showCmd)
}
