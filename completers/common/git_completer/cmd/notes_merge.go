package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var notes_mergeCmd = &cobra.Command{
	Use:   "merge",
	Short: "Merge the given notes ref into the current notes ref",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(notes_mergeCmd).Standalone()

	notes_mergeCmd.Flags().Bool("abort", false, "abort/reset an in-progress git notes merge")
	notes_mergeCmd.Flags().Bool("commit", false, "finalize an in-progress git notes merge")
	notes_mergeCmd.Flags().BoolP("quiet", "q", false, "when merging notes, operate quietly")
	notes_mergeCmd.Flags().StringP("strategy", "s", "", "resolve notes conflicts using the given strategy")
	notes_mergeCmd.Flags().BoolP("verbose", "v", false, "when merging notes, be more verbose")
	notesCmd.AddCommand(notes_mergeCmd)

	notes_mergeCmd.MarkFlagsMutuallyExclusive("abort", "commit", "strategy")

	carapace.Gen(notes_mergeCmd).FlagCompletion(carapace.ActionMap{
		"strategy": git.ActionNotesMergeStrategies(),
	})

	carapace.Gen(notes_mergeCmd).PositionalCompletion(
		git.ActionNotes(),
	)
}
