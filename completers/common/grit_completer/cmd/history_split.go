package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var history_splitCmd = &cobra.Command{
	Use:   "split",
	Short: "Peel path-limited changes out of a commit into a new parent and replay descendants",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(history_splitCmd).Standalone()

	history_splitCmd.Flags().Bool("dry-run", false, "Print planned ref updates without moving refs (new commit objects are still written)")
	history_splitCmd.Flags().String("empty", "drop", "When a replayed commit would not change the tree: drop (default), keep an empty commit, or abort")
	history_splitCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	history_splitCmd.Flags().Bool("keep-committer", false, "Keep each replayed commit's original committer line instead of the configured committer")
	history_splitCmd.Flags().StringP("message", "m", "", "Message for the new parent commit (the peeled-off changes)")
	history_splitCmd.Flags().String("message-file", "", "Read the new parent message from a file (`-` for stdin)")
	history_splitCmd.Flags().Bool("pathspec-file-nul", false, "`--pathspec-from-file` entries are NUL-separated")
	history_splitCmd.Flags().String("pathspec-from-file", "", "Read pathspecs from a file (`-` for stdin) instead of or in addition to argv paths")
	history_splitCmd.Flags().String("rest-message", "", "Replace the remaining commit's message after the split")
	history_splitCmd.Flags().String("update-refs", "branches", "Which refs to rewrite: all descendant branches (default) or only `HEAD`")
	historyCmd.AddCommand(history_splitCmd)

	carapace.Gen(history_splitCmd).FlagCompletion(carapace.ActionMap{
		"empty":       carapace.ActionValues("abort", "drop", "keep"),
		"update-refs": carapace.ActionValues("branches", "head"),
	})

	carapace.Gen(history_splitCmd).PositionalCompletion(
		git.ActionRefs(git.RefOption{}.Default()),
	)
}
