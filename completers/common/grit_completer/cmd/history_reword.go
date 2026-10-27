package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var history_rewordCmd = &cobra.Command{
	Use:   "reword",
	Short: "Replace a commit message and replay descendants onto the new history",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(history_rewordCmd).Standalone()

	history_rewordCmd.Flags().String("author", "", "Replacement author ident (`Name <email>` or bare email)")
	history_rewordCmd.Flags().Bool("dry-run", false, "Print planned ref updates without moving refs (new commit objects are still written)")
	history_rewordCmd.Flags().String("empty", "drop", "When a replayed commit would not change the tree: drop (default), keep an empty commit, or abort")
	history_rewordCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	history_rewordCmd.Flags().Bool("keep-committer", false, "Keep each replayed commit's original committer line instead of the configured committer")
	history_rewordCmd.Flags().StringP("message", "m", "", "New commit message")
	history_rewordCmd.Flags().String("message-file", "", "Read the new message from a file (`-` for stdin)")
	history_rewordCmd.Flags().String("update-refs", "branches", "Which refs to rewrite: all descendant branches (default) or only `HEAD`")
	historyCmd.AddCommand(history_rewordCmd)

	carapace.Gen(history_rewordCmd).FlagCompletion(carapace.ActionMap{
		"empty":       carapace.ActionValues("abort", "drop", "keep"),
		"update-refs": carapace.ActionValues("branches", "head"),
	})

	carapace.Gen(history_rewordCmd).FlagCompletion(carapace.ActionMap{
		"message-file": carapace.ActionFiles(),
	})

	carapace.Gen(history_rewordCmd).PositionalCompletion(
		git.ActionRefs(git.RefOption{}.Default()),
	)
}
