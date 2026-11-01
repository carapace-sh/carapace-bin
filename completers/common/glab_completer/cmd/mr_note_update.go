package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var mr_note_updateCmd = &cobra.Command{
	Use:   "update [<id> | <branch>] <note-id>",
	Short: "Update the body of a note on a merge request. (EXPERIMENTAL)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mr_note_updateCmd).Standalone()

	mr_note_updateCmd.Flags().StringSlice("attach", nil, "(EXPERIMENTAL) Upload a file and reference it at the end of the note. Use \"-\" to read the file from standard input. Repeat the flag to attach multiple files.")
	mr_note_updateCmd.Flags().StringP("message", "m", "", "New note body. If omitted, opens an editor or reads from stdin.")
	mr_noteCmd.AddCommand(mr_note_updateCmd)

	carapace.Gen(mr_note_updateCmd).FlagCompletion(carapace.ActionMap{
		"attach": carapace.ActionFiles(),
	})

	carapace.Gen(mr_note_updateCmd).PositionalAnyCompletion(
		action.ActionMergeRequestsAndBranches(mr_note_updateCmd, ""),
	)
}
