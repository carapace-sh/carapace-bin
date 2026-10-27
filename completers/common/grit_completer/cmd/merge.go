package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var mergeCmd = &cobra.Command{
	Use:     "merge",
	Short:   "Merge another branch into the current one",
	GroupID: "branches",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mergeCmd).Standalone()

	mergeCmd.Flags().StringS("X", "X", "", "Prefer ours or theirs on overlapping edits (`-X ours` / `-X theirs`)")
	mergeCmd.Flags().Bool("abort", false, "Cancel an in-progress merge and restore the pre-merge state")
	mergeCmd.Flags().Bool("allow-unrelated-histories", false, "Merge histories that share no commits (uses the empty tree as merge base)")
	mergeCmd.Flags().Bool("check", false, "Predict whether the merge would succeed without changing refs, index, or worktree")
	mergeCmd.Flags().String("conflict", "", "Conflict marker style (`merge`, `diff3`, or `zdiff3`)")
	mergeCmd.Flags().Bool("continue", false, "Finish an in-progress merge after resolving conflicts (or after `--no-commit`)")
	mergeCmd.Flags().Bool("dry-run", false, "Report the planned merge outcome without changing refs, index, or worktree")
	mergeCmd.Flags().Bool("ff-only", false, "Refuse when a fast-forward is not possible (exit code 128)")
	mergeCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	mergeCmd.Flags().StringP("message", "m", "", "Commit message text")
	mergeCmd.Flags().String("message-file", "", "Read the message from a file (`-` for stdin)")
	mergeCmd.Flags().Bool("no-commit", false, "Apply the merge to the index and worktree but do not commit yet")
	mergeCmd.Flags().Bool("no-ff", false, "Record a merge commit even when a fast-forward is possible")
	mergeCmd.Flags().Bool("squash", false, "Squash merged commits into staged changes and write `SQUASH_MSG`")
	rootCmd.AddCommand(mergeCmd)

	carapace.Gen(mergeCmd).FlagCompletion(carapace.ActionMap{
		"X": carapace.ActionValuesDescribed(
			"ours", "Prefer the current branch when both sides changed the same region",
			"theirs", "Prefer the branch being merged when both sides changed the same region",
		),
		"conflict": carapace.ActionValuesDescribed(
			"diff3", "Three-section markers including the base section",
			"merge", "Standard two-section conflict markers",
			"zdiff3", "Zealous diff3 (`zdiff3`)",
		),
	})

	carapace.Gen(mergeCmd).PositionalCompletion(
		git.ActionRefs(git.RefOption{
			LocalBranches:  true,
			RemoteBranches: true,
			Tags:           true,
		}).FilterArgs(),
	)
}
