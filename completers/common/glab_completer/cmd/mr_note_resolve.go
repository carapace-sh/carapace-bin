package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var mr_note_resolveCmd = &cobra.Command{
	Use:   "resolve [<id> | <branch>] <discussion-id>",
	Short: "Resolve a discussion on a merge request. (EXPERIMENTAL)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mr_note_resolveCmd).Standalone()

	mr_noteCmd.AddCommand(mr_note_resolveCmd)

	carapace.Gen(mr_note_resolveCmd).PositionalAnyCompletion(
		action.ActionMergeRequestsAndBranches(mr_note_resolveCmd, ""),
	)
}
