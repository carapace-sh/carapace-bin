package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_notes_showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show the note for a revision",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_notes_showCmd).Standalone()

	help_notesCmd.AddCommand(help_notes_showCmd)
}
