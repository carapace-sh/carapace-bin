package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var history_dropCmd = &cobra.Command{
	Use:   "drop",
	Short: "Remove a commit and replay descendants onto its parent",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(history_dropCmd).Standalone()

	history_dropCmd.Flags().Bool("dry-run", false, "Print planned ref updates without moving refs (new commit objects are still written)")
	history_dropCmd.Flags().String("empty", "drop", "When a replayed commit would not change the tree: drop (default), keep an empty commit, or abort")
	history_dropCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	history_dropCmd.Flags().Bool("keep-committer", false, "Keep each replayed commit's original committer line instead of the configured committer")
	history_dropCmd.Flags().String("update-refs", "branches", "Which refs to rewrite: all descendant branches (default) or only `HEAD`")
	historyCmd.AddCommand(history_dropCmd)

	carapace.Gen(history_dropCmd).FlagCompletion(carapace.ActionMap{
		"empty":       carapace.ActionValues("abort", "drop", "keep"),
		"update-refs": carapace.ActionValues("branches", "head"),
	})

	carapace.Gen(history_dropCmd).PositionalCompletion(
		git.ActionRefs(git.RefOption{}.Default()),
	)
}
