package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var history_undoCmd = &cobra.Command{
	Use:   "undo",
	Short: "Undo the last applied history rewrite using the recorded ref plan",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(history_undoCmd).Standalone()

	history_undoCmd.Flags().Bool("dry-run", false, "Print planned ref updates without moving refs (new commit objects are still written)")
	history_undoCmd.Flags().String("empty", "drop", "When a replayed commit would not change the tree: drop (default), keep an empty commit, or abort")
	history_undoCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	history_undoCmd.Flags().Bool("keep-committer", false, "Keep each replayed commit's original committer line instead of the configured committer")
	history_undoCmd.Flags().String("update-refs", "branches", "Which refs to rewrite: all descendant branches (default) or only `HEAD`")
	historyCmd.AddCommand(history_undoCmd)

	carapace.Gen(history_undoCmd).FlagCompletion(carapace.ActionMap{
		"empty":       carapace.ActionValues("abort", "drop", "keep"),
		"update-refs": carapace.ActionValues("branches", "head"),
	})
}
