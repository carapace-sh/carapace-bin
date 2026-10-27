package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var pickCmd = &cobra.Command{
	Use:   "pick",
	Short: "Cherry-pick commits onto the current branch",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pickCmd).Standalone()

	pickCmd.Flags().StringS("X", "X", "", "Auto-resolve replay conflicts: `ours` keeps the current branch side, `theirs` keeps the replayed change")
	pickCmd.Flags().Bool("abort", false, "Cancel an in-progress pick and restore the pre-pick state")
	pickCmd.Flags().Bool("check", false, "Predict whether the pick would succeed without changing refs, the index, or the worktree")
	pickCmd.Flags().String("conflict", "", "Conflict marker style when a step would stop with conflicts")
	pickCmd.Flags().Bool("continue", false, "Finish a pick stopped on conflicts or an empty commit")
	pickCmd.Flags().Bool("dry-run", false, "List commits that would be picked without changing refs or the worktree")
	pickCmd.Flags().String("empty", "stop", "When a pick step would not change the tree: stop (default), keep an empty commit, or drop the step")
	pickCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	pickCmd.Flags().String("mainline", "", "Parent number when picking a merge commit (1-based, long only)")
	pickCmd.Flags().StringP("message", "m", "", "Replace the pick message (single commit only)")
	pickCmd.Flags().String("message-file", "", "Read the pick message from a file (single commit only)")
	pickCmd.Flags().Bool("no-commit", false, "Apply the pick to the index only; do not create a commit yet")
	pickCmd.Flags().BoolP("record-origin", "x", false, "Append `(cherry picked from commit …)` to each new commit message")
	pickCmd.Flags().Bool("skip", false, "Skip the current pick step and continue with the next commit")
	rootCmd.AddCommand(pickCmd)

	carapace.Gen(pickCmd).FlagCompletion(carapace.ActionMap{
		"X": carapace.ActionValuesDescribed(
			"ours", "Favor the upstream / onto side during rebase replay (matches `git rebase -X ours`)",
			"theirs", "Favor the commit being replayed (matches `git rebase -X theirs`)",
		),
		"conflict": carapace.ActionValuesDescribed(
			"diff3", "Three-section markers including the base section",
			"merge", "Standard two-section conflict markers",
			"zdiff3", "Zealous diff3 (`zdiff3`)",
		),
		"empty": carapace.ActionValues("drop", "keep", "stop"),
	})

	carapace.Gen(pickCmd).PositionalAnyCompletion(
		git.ActionRefRanges(git.RefOption{}.Default()),
	)
}
