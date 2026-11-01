package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_notes_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add or replace a note on a revision",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_notes_addCmd).Standalone()

	help_notesCmd.AddCommand(help_notes_addCmd)
}
