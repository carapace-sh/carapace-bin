package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var notes_getRefCmd = &cobra.Command{
	Use:   "get-ref",
	Short: "Print the current notes ref",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(notes_getRefCmd).Standalone()

	notesCmd.AddCommand(notes_getRefCmd)
}
