package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var history_absorbCmd = &cobra.Command{
	Use:   "absorb",
	Short: "Assign staged hunks into recent commits and replay descendants once",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(history_absorbCmd).Standalone()

	history_absorbCmd.Flags().String("base", "", "Oldest commit boundary (exclusive) for candidate commits")
	history_absorbCmd.Flags().Bool("dry-run", false, "Print planned ref updates without moving refs (new commit objects are still written)")
	history_absorbCmd.Flags().String("empty", "drop", "When a replayed commit would not change the tree: drop (default), keep an empty commit, or abort")
	history_absorbCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	history_absorbCmd.Flags().Bool("keep-committer", false, "Keep each replayed commit's original committer line instead of the configured committer")
	history_absorbCmd.Flags().String("update-refs", "branches", "Which refs to rewrite: all descendant branches (default) or only `HEAD`")
	historyCmd.AddCommand(history_absorbCmd)

	carapace.Gen(history_absorbCmd).FlagCompletion(carapace.ActionMap{
		"base":        git.ActionRefs(git.RefOption{}.Default()),
		"empty":       carapace.ActionValues("abort", "drop", "keep"),
		"update-refs": carapace.ActionValues("branches", "head"),
	})
}
