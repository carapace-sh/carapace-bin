package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_notes_removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove notes from one or more revisions",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_notes_removeCmd).Standalone()

	help_notesCmd.AddCommand(help_notes_removeCmd)
}
