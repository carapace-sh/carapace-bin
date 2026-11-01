package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/lore_completer/cmd/action"
	"github.com/spf13/cobra"
)

var branch_mergeCmd = &cobra.Command{
	Use:   "merge",
	Short: "Merge two branches",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(branch_mergeCmd).Standalone()

	branch_mergeCmd.Flags().BoolP("help", "h", false, "Print help")
	branch_mergeCmd.Flags().String("id", "", "ID of the source branch to merge into the current branch")
	branch_mergeCmd.Flags().String("message", "", "Change the message for committing when no conflicts arise from the merge")
	branchCmd.AddCommand(branch_mergeCmd)

	carapace.Gen(branch_mergeCmd).PositionalCompletion(
		action.ActionBranches(branch_mergeCmd),
	)
}
