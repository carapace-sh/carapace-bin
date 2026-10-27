package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var revertCmd = &cobra.Command{
	Use:   "revert",
	Short: "Undo commits with new inverse commits",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(revertCmd).Standalone()

	revertCmd.Flags().StringS("X", "X", "", "Auto-resolve replay conflicts: `ours` keeps the current branch side, `theirs` keeps the replayed change")
	revertCmd.Flags().Bool("abort", false, "Cancel an in-progress revert and restore the pre-revert state")
	revertCmd.Flags().Bool("check", false, "Predict whether the revert would succeed without changing refs, the index, or the worktree")
	revertCmd.Flags().String("conflict", "", "Conflict marker style when a step would stop with conflicts")
	revertCmd.Flags().Bool("continue", false, "Finish a revert stopped on conflicts or an empty commit")
	revertCmd.Flags().Bool("dry-run", false, "List commits that would be reverted without changing refs or the worktree")
	revertCmd.Flags().String("empty", "stop", "When a revert step would not change the tree: stop (default), keep an empty commit, or drop the step")
	revertCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	revertCmd.Flags().String("mainline", "", "Parent number when reverting a merge commit (1-based, long only)")
	revertCmd.Flags().StringP("message", "m", "", "Replace the revert message (single commit only)")
	revertCmd.Flags().String("message-file", "", "Read the revert message from a file (single commit only)")
	revertCmd.Flags().Bool("no-commit", false, "Apply the revert to the index only; do not create a commit yet")
	revertCmd.Flags().Bool("skip", false, "Skip the current revert step and continue with the next commit")
	rootCmd.AddCommand(revertCmd)

	carapace.Gen(revertCmd).FlagCompletion(carapace.ActionMap{
		"X": carapace.ActionValuesDescribed(
			"ours", "Favor the upstream / onto side during rebase replay (matches `git rebase -X ours`)",
			"theirs", "Favor the commit being replayed (matches `git rebase -X theirs`)",
		),
		"conflict": carapace.ActionValuesDescribed(
			"diff3", "Three-section markers including the base section",
			"merge", "Standard two-section conflict markers",
			"zdiff3", "Zealous diff3 (`zdiff3`)",
		),
		"empty": carapace.ActionValuesDescribed(
			"drop", "Skip the commit and continue with the next one",
			"keep", "Record an empty revert commit for the step",
			"stop", "Stop and leave sequencer state for `--continue` / `--skip`",
		),
	})

	carapace.Gen(revertCmd).PositionalAnyCompletion(
		git.ActionRefRanges(git.RefOption{}.Default()),
	)
}
