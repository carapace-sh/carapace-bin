package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/but"
	"github.com/spf13/cobra"
)

var pickCmd = &cobra.Command{
	Use:     "pick",
	Short:   "Cherry-pick commits into an applied branch",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: "branching and committing",
}

func init() {
	carapace.Gen(pickCmd).Standalone()

	pickCmd.Flags().StringP("above", "A", "", "Place the picked commits above BRANCH_OR_COMMIT")
	pickCmd.Flags().Bool("allow-merged", false, "Allow targeting branches and commits that are already merged upstream")
	pickCmd.Flags().StringP("below", "B", "", "Place the picked commits below BRANCH_OR_COMMIT")
	pickCmd.Flags().StringP("branch", "b", "", "Place the picked commits on the branch BRANCH")
	pickCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	rootCmd.AddCommand(pickCmd)

	carapace.Gen(pickCmd).PositionalCompletion(
		carapace.Batch(
			but.ActionCliIds(but.CliIdsOpts{}.Default()),
			but.ActionTargets(), // TODO targets correct?
		).ToA(),
		but.ActionLocalBranches(),
	)
}
