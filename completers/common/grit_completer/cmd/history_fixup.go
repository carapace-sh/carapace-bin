package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var history_fixupCmd = &cobra.Command{
	Use:   "fixup",
	Short: "Apply staged changes onto a commit and replay descendants",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(history_fixupCmd).Standalone()

	history_fixupCmd.Flags().Bool("dry-run", false, "Print planned ref updates without moving refs (new commit objects are still written)")
	history_fixupCmd.Flags().String("empty", "drop", "When a replayed commit would not change the tree: drop (default), keep an empty commit, or abort")
	history_fixupCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	history_fixupCmd.Flags().StringSlice("hunks", nil, "Unstaged hunk ids to fold (`grit diff --unstaged --json`, comma-separated or repeated)")
	history_fixupCmd.Flags().Bool("keep-committer", false, "Keep each replayed commit's original committer line instead of the configured committer")
	history_fixupCmd.Flags().StringSlice("lines", nil, "Unstaged new-side line ranges to fold (`path:start-end` or `path:line`, comma-separated or repeated)")
	history_fixupCmd.Flags().StringP("message", "m", "", "Replace the target commit message")
	history_fixupCmd.Flags().String("message-file", "", "Read the replacement message from a file (`-` for stdin)")
	history_fixupCmd.Flags().Bool("pathspec-file-nul", false, "`--pathspec-from-file` entries are NUL-separated")
	history_fixupCmd.Flags().String("pathspec-from-file", "", "Read pathspecs from a file (`-` for stdin) instead of or in addition to argv paths")
	history_fixupCmd.Flags().String("update-refs", "branches", "Which refs to rewrite: all descendant branches (default) or only `HEAD`")
	historyCmd.AddCommand(history_fixupCmd)

	carapace.Gen(history_fixupCmd).FlagCompletion(carapace.ActionMap{
		"empty":       carapace.ActionValues("abort", "drop", "keep"),
		"update-refs": carapace.ActionValues("branches", "head"),
	})

	carapace.Gen(history_fixupCmd).PositionalCompletion(
		git.ActionRefs(git.RefOption{}.Default()),
	)
}
