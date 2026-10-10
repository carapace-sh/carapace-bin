package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var mr_note_draft_updateCmd = &cobra.Command{
	Use:   "update [<id> | <branch>] <draft-note-id>",
	Short: "Update the body of one of your pending review comments. (EXPERIMENTAL)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mr_note_draft_updateCmd).Standalone()

	mr_note_draft_updateCmd.Flags().StringSlice("attach", nil, "(EXPERIMENTAL) Upload a file and reference it at the end of the comment. Use \"-\" to read the file from standard input. Repeat the flag to attach multiple files.")
	mr_note_draft_updateCmd.Flags().StringP("message", "m", "", "New comment body. If omitted, opens an editor or reads from stdin.")
	mr_note_draftCmd.AddCommand(mr_note_draft_updateCmd)

	carapace.Gen(mr_note_draft_updateCmd).FlagCompletion(carapace.ActionMap{
		"attach": carapace.ActionFiles(),
	})

	carapace.Gen(mr_note_draft_updateCmd).PositionalAnyCompletion(
		action.ActionMergeRequestsAndBranches(mr_note_draft_updateCmd, ""),
	)
}
