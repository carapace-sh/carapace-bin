package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var mr_note_draft_createCmd = &cobra.Command{
	Use:   "create [<id> | <branch>]",
	Short: "Add a pending review comment to a merge request. (EXPERIMENTAL)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mr_note_draft_createCmd).Standalone()

	mr_note_draft_createCmd.Flags().StringSlice("attach", nil, "(EXPERIMENTAL) Upload a file and reference it at the end of the comment. Use \"-\" to read the file from standard input. Repeat the flag to attach multiple files.")
	mr_note_draft_createCmd.Flags().String("file", "", "File path for a diff comment, like <path/to/file>. Targets the latest merge request diff version.")
	mr_note_draft_createCmd.Flags().String("line", "", "Line in the new version. A single line number, like 42, or a range, like 10:15.")
	mr_note_draft_createCmd.Flags().StringP("message", "m", "", "Comment message. If omitted, opens an editor or reads from stdin.")
	mr_note_draft_createCmd.Flags().String("old-line", "", "Line in the old version, for commenting on a removed line.")
	mr_note_draft_createCmd.Flags().String("reply", "", "Reply to an existing discussion. Accepts a full discussion ID or a unique prefix of at least 8 characters.")
	mr_note_draftCmd.AddCommand(mr_note_draft_createCmd)

	carapace.Gen(mr_note_draft_createCmd).FlagCompletion(carapace.ActionMap{
		"attach": carapace.ActionFiles(),
		"file":   carapace.ActionFiles(),
	})

	carapace.Gen(mr_note_draft_createCmd).PositionalAnyCompletion(
		action.ActionMergeRequestsAndBranches(mr_note_draft_createCmd, ""),
	)
}
