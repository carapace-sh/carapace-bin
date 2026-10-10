package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/carapace-sh/carapace-jq/pkg/actions/tools/jq"
	"github.com/spf13/cobra"
)

var mr_note_draft_listCmd = &cobra.Command{
	Use:   "list [<id> | <branch>]",
	Short: "List your pending review comments on a merge request. (EXPERIMENTAL)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mr_note_draft_listCmd).Standalone()

	mr_note_draft_listCmd.Flags().String("file", "", "Show only pending diff comments on this file path.")
	mr_note_draft_listCmd.Flags().String("jq", "", "Filter JSON output with a jq expression.")
	mr_note_draft_listCmd.Flags().StringP("output", "F", "text", "Format output as: text, json.")
	mr_note_draftCmd.AddCommand(mr_note_draft_listCmd)

	carapace.Gen(mr_note_draft_listCmd).FlagCompletion(carapace.ActionMap{
		"file":   carapace.ActionFiles(),
		"jq":     jq.ActionFilters(),
		"output": carapace.ActionValues("text", "json"),
	})
	carapace.Gen(mr_note_draft_listCmd).PositionalAnyCompletion(
		action.ActionMergeRequestsAndBranches(mr_note_draft_listCmd, ""),
	)

}
