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
	replayCmd.Flags().String("onto", "", "starting point at which to create the new commits")
	common.AddCommitFormattingOptions(replayCmd)
	common.AddCommitLimitingOptions(replayCmd)
	common.AddCommitOrderingOptions(replayCmd)
	common.AddHistorySimplificationOptions(replayCmd)
	common.AddObjectTraversalOptions(replayCmd)
	rootCmd.AddCommand(replayCmd)

	carapace.Gen(replayCmd).FlagCompletion(carapace.ActionMap{
		"advance": git.ActionRefs(git.RefOption{}.Default()),
		"onto":    git.ActionLocalBranches(),
	})

	carapace.Gen(replayCmd).PositionalAnyCompletion(
		git.ActionRefRanges(git.RefOption{}.Default()),
	)
}
