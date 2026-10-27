package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var history_rebaseCmd = &cobra.Command{
	Use:   "rebase",
	Short: "Replay local commits onto a new base (same as top-level `grit rebase`)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(history_rebaseCmd).Standalone()

	history_rebaseCmd.Flags().StringS("X", "X", "", "Auto-resolve replay conflicts: `ours` keeps the onto side, `theirs` keeps replayed commits")
	history_rebaseCmd.Flags().Bool("abort", false, "Cancel an in-progress rebase and restore the branch to its pre-rebase tip")
	history_rebaseCmd.Flags().Bool("check", false, "Predict whether the rebase would succeed without changing refs, index, or worktree")
	history_rebaseCmd.Flags().Bool("continue", false, "Finish a rebase stopped on conflicts after resolving and staging paths")
	history_rebaseCmd.Flags().Bool("dry-run", false, "Print planned ref updates without moving refs (new commit objects are still written)")
	history_rebaseCmd.Flags().String("empty", "drop", "When a replayed commit would not change the tree: drop (default), keep an empty commit, or abort")
	history_rebaseCmd.Flags().StringSlice("exec", nil, "Run a shell command after each replayed commit is checked out (repeatable; incompatible with `--dry-run` and `--check`)")
	history_rebaseCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	history_rebaseCmd.Flags().Bool("keep-committer", false, "Keep each replayed commit's original committer line instead of the configured committer")
	history_rebaseCmd.Flags().String("onto", "", "Replay onto this commit instead of `UPSTREAM`")
	history_rebaseCmd.Flags().Bool("skip", false, "Skip the current commit and replay the rest")
	history_rebaseCmd.Flags().String("update-refs", "branches", "Which refs to rewrite: all descendant branches (default) or only `HEAD`")
	historyCmd.AddCommand(history_rebaseCmd)

	carapace.Gen(history_rebaseCmd).FlagCompletion(carapace.ActionMap{
		"X": carapace.ActionValuesDescribed(
			"ours", "Favor the upstream / onto side during rebase replay (matches `git rebase -X ours`)",
			"theirs", "Favor the commit being replayed (matches `git rebase -X theirs`)",
		),
		"empty":       carapace.ActionValues("abort", "drop", "keep"),
		"onto":        git.ActionRefs(git.RefOption{}.Default()),
		"update-refs": carapace.ActionValues("branches", "head"),
	})

	carapace.Gen(history_rebaseCmd).PositionalCompletion(
		git.ActionRefs(git.RefOption{}.Default()),
	)
}
