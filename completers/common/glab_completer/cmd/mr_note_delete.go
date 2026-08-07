package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mr_note_deleteCmd = &cobra.Command{
	Use:   "delete [<id> | <branch>] <note-id>",
	Short: "Delete a note from a merge request. (EXPERIMENTAL)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mr_note_deleteCmd).Standalone()

	mr_note_deleteCmd.Flags().BoolP("yes", "y", false, "Skip confirmation prompt.")
	mr_noteCmd.AddCommand(mr_note_deleteCmd)
}
