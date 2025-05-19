package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/git_completer/cmd/common"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var revListCmd = &cobra.Command{
	Use:     "rev-list",
	Short:   "Lists commit objects in reverse chronological order",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_low_level_interrogator].ID,
}

func init() {
	carapace.Gen(revListCmd).Standalone()

	common.AddBisectionHelperOptions(revListCmd)
	common.AddCommitFormattingOptions(revListCmd)
	common.AddCommitLimitingOptions(revListCmd)
	common.AddCommitOrderingOptions(revListCmd)
	common.AddHistorySimplificationOptions(revListCmd)
	common.AddObjectTraversalOptions(revListCmd)
	rootCmd.AddCommand(revListCmd)

	carapace.Gen(revListCmd).PositionalAnyCompletion(
		git.ActionRefs(git.RefOption{}.Default()),
	)

	carapace.Gen(revListCmd).DashAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			batch := carapace.Batch()
			for _, ref := range revListCmd.Flags().Args()[:revListCmd.Flags().ArgsLenAtDash()] {
				batch = append(batch, git.ActionRefFiles(ref))
			}
			return batch.ToA()
		}),
	)
}
