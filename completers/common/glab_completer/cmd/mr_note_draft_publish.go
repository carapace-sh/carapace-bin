package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var mr_note_draft_publishCmd = &cobra.Command{
	Use:   "publish [<id> | <branch>]",
	Short: "Publish all your pending review comments on a merge request. (EXPERIMENTAL)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mr_note_draft_publishCmd).Standalone()

	mr_note_draft_publishCmd.Flags().Bool("internal", false, "Mark the summary note as internal. Requires --message.")
	mr_note_draft_publishCmd.Flags().StringP("message", "m", "", "Summary note to add to the merge request when publishing.")
	mr_note_draft_publishCmd.Flags().String("reviewer-state", "", "Set the review state after publishing: requested_changes, reviewed. Does not record a formal approval.")
	mr_note_draft_publishCmd.Flags().BoolP("yes", "y", false, "Skip confirmation prompt.")
	mr_note_draftCmd.AddCommand(mr_note_draft_publishCmd)

	carapace.Gen(mr_note_draft_publishCmd).FlagCompletion(carapace.ActionMap{
		"reviewer-state": carapace.ActionValues("requested_changes", "reviewed"),
	})
	carapace.Gen(mr_note_draft_publishCmd).PositionalAnyCompletion(
		action.ActionMergeRequestsAndBranches(mr_note_draft_publishCmd, ""),
	)

}
