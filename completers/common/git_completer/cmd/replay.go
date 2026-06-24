package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/git_completer/cmd/common"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var replayCmd = &cobra.Command{
	Use:     "replay",
	Short:   "EXPERIMENTAL: Replay commits on a new base, works with bare repos too",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_low_level_manipulator].ID,
}

func init() {
	carapace.Gen(replayCmd).Standalone()

	replayCmd.Flags().String("advance", "", "starting point at which to create the new commits")
	replayCmd.Flags().Bool("contained", false, "update all branches that point at commits in the revision range")
	replayCmd.Flags().String("onto", "", "starting point at which to create the new commits")
	replayCmd.Flags().String("ref", "", "override which reference is updated with the result of the replay")
	replayCmd.Flags().String("ref-action", "", "control how references are updated")
	replayCmd.Flags().String("revert", "", "starting point at which to create the reverted commits")
	common.AddCommitFormattingOptions(replayCmd)
	common.AddCommitLimitingOptions(replayCmd)
	common.AddCommitOrderingOptions(replayCmd)
	common.AddHistorySimplificationOptions(replayCmd)
	common.AddObjectTraversalOptions(replayCmd)
	rootCmd.AddCommand(replayCmd)

	replayCmd.MarkFlagsMutuallyExclusive("onto", "advance", "revert")

	replayCmd.Flag("ref-action").NoOptDefVal = "update"

	carapace.Gen(replayCmd).FlagCompletion(carapace.ActionMap{
		"advance":    git.ActionRefs(git.RefOption{}.Default()),
		"onto":       git.ActionLocalBranches(),
		"ref":        git.ActionRefs(git.RefOption{}.Default()),
		"ref-action": carapace.ActionValuesDescribed("update", "update refs directly using an atomic transaction", "print", "output update-ref commands for pipeline use"),
		"revert":     git.ActionLocalBranches(),
	})

	carapace.Gen(replayCmd).PositionalAnyCompletion(
		git.ActionRefRanges(git.RefOption{}.Default()),
	)
}
