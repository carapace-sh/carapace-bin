package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var history_moveCmd = &cobra.Command{
	Use:   "move",
	Short: "Reorder a commit on the first-parent chain of `HEAD`",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(history_moveCmd).Standalone()

	history_moveCmd.Flags().String("after", "", "Place the commit immediately after this revision")
	history_moveCmd.Flags().String("before", "", "Place the commit immediately before this revision")
	history_moveCmd.Flags().Bool("dry-run", false, "Print planned ref updates without moving refs (new commit objects are still written)")
	history_moveCmd.Flags().String("empty", "drop", "When a replayed commit would not change the tree: drop (default), keep an empty commit, or abort")
	history_moveCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	history_moveCmd.Flags().Bool("keep-committer", false, "Keep each replayed commit's original committer line instead of the configured committer")
	history_moveCmd.Flags().String("update-refs", "branches", "Which refs to rewrite: all descendant branches (default) or only `HEAD`")
	historyCmd.AddCommand(history_moveCmd)

	carapace.Gen(history_moveCmd).FlagCompletion(carapace.ActionMap{
		"after":       git.ActionRefs(git.RefOption{}.Default()),
		"before":      git.ActionRefs(git.RefOption{}.Default()),
		"empty":       carapace.ActionValues("abort", "drop", "keep"),
		"update-refs": carapace.ActionValues("branches", "head"),
	})

	carapace.Gen(history_moveCmd).PositionalCompletion(
		git.ActionRefs(git.RefOption{}.Default()),
	)
}
