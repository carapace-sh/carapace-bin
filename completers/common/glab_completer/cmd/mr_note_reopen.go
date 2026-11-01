package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var mr_note_reopenCmd = &cobra.Command{
	Use:   "reopen [<id> | <branch>] <discussion-id>",
	Short: "Reopen a discussion on a merge request. (EXPERIMENTAL)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mr_note_reopenCmd).Standalone()

	mr_noteCmd.AddCommand(mr_note_reopenCmd)

	carapace.Gen(mr_note_reopenCmd).PositionalAnyCompletion(
		action.ActionMergeRequestsAndBranches(mr_note_reopenCmd, ""),
	)
}
