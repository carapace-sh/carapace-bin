package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var history_squashCmd = &cobra.Command{
	Use:   "squash",
	Short: "Fold a revision range into the oldest commit and replay descendants",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(history_squashCmd).Standalone()

	history_squashCmd.Flags().Bool("combine", false, "Build a combined autosquash-style message from every commit in the range")
	history_squashCmd.Flags().Bool("dry-run", false, "Print planned ref updates without moving refs (new commit objects are still written)")
	history_squashCmd.Flags().String("empty", "drop", "When a replayed commit would not change the tree: drop (default), keep an empty commit, or abort")
	history_squashCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	history_squashCmd.Flags().Bool("keep-committer", false, "Keep each replayed commit's original committer line instead of the configured committer")
	history_squashCmd.Flags().StringP("message", "m", "", "Message for the squashed commit (default: keep the oldest commit's message)")
	history_squashCmd.Flags().String("message-file", "", "Read the squashed commit message from a file (`-` for stdin)")
	history_squashCmd.Flags().String("update-refs", "branches", "Which refs to rewrite: all descendant branches (default) or only `HEAD`")
	historyCmd.AddCommand(history_squashCmd)

	carapace.Gen(history_squashCmd).FlagCompletion(carapace.ActionMap{
		"empty":       carapace.ActionValues("abort", "drop", "keep"),
		"update-refs": carapace.ActionValues("branches", "head"),
	})

	carapace.Gen(history_squashCmd).PositionalAnyCompletion(
		git.ActionRefRanges(git.RefOption{}.Default()),
	)
}
