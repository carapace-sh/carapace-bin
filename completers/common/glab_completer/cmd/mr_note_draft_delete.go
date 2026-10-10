package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var mr_note_draft_deleteCmd = &cobra.Command{
	Use:   "delete [<id> | <branch>] <draft-note-id>",
	Short: "Delete one of your pending review comments from a merge request. (EXPERIMENTAL)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mr_note_draft_deleteCmd).Standalone()

	mr_note_draft_deleteCmd.Flags().BoolP("yes", "y", false, "Skip confirmation prompt.")
	mr_note_draftCmd.AddCommand(mr_note_draft_deleteCmd)

	carapace.Gen(mr_note_draft_deleteCmd).PositionalAnyCompletion(
		action.ActionMergeRequestsAndBranches(mr_note_draft_deleteCmd, ""),
	)

}
