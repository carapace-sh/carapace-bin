package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var mr_note_createCmd = &cobra.Command{
	Use:   "create [<id> | <branch>]",
	Short: "Create a comment or discussion on a merge request. (EXPERIMENTAL)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mr_note_createCmd).Standalone()

	mr_note_createCmd.Flags().String("file", "", "File path for a diff comment, like <path/to/file>. Targets the latest merge request diff version.")
	mr_note_createCmd.Flags().String("line", "", "Line in the new version. A single line number, like 42, or a range, like 10:15.")
	mr_note_createCmd.Flags().StringP("message", "m", "", "Comment or note message.")
	mr_note_createCmd.Flags().Int("old-line", 0, "Line in the old version, for commenting on a removed line.")
	mr_note_createCmd.Flags().String("reply", "", "Reply to an existing discussion. Accepts a full discussion ID or a unique prefix of at least 8 characters.")
	mr_note_createCmd.Flags().Bool("resolvable", true, "Create the note as a resolvable discussion thread. Set to false to create a non-resolvable note.")
	mr_note_createCmd.Flags().Bool("unique", false, "Don't create a note if a note with the same body already exists. Reads all merge request comments first.")
	mr_noteCmd.AddCommand(mr_note_createCmd)

	carapace.Gen(mr_note_createCmd).FlagCompletion(carapace.ActionMap{
		"file": carapace.ActionFiles(),
	})

	carapace.Gen(mr_note_createCmd).PositionalAnyCompletion(
		action.ActionMergeRequestsAndBranches(mr_note_createCmd, ""),
	)
}
