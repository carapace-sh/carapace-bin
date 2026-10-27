package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_notes_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List notes attached to commits",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_notes_listCmd).Standalone()

	help_notesCmd.AddCommand(help_notes_listCmd)
}
