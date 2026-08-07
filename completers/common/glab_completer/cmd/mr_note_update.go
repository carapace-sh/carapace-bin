package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mr_note_updateCmd = &cobra.Command{
	Use:   "update [<id> | <branch>] <note-id>",
	Short: "Update the body of a note on a merge request. (EXPERIMENTAL)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mr_note_updateCmd).Standalone()

	mr_note_updateCmd.Flags().StringP("message", "m", "", "New note body. If omitted, opens an editor or reads from stdin.")
	mr_noteCmd.AddCommand(mr_note_updateCmd)
}
