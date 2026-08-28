package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/but"
	"github.com/spf13/cobra"
)

var moveCmd = &cobra.Command{
	Use:     "move",
	Short:   "Move commits and changes around",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: "editing commits",
}

func init() {
	carapace.Gen(moveCmd).Standalone()

	moveCmd.Flags().StringP("above", "A", "", "Place <SOURCES> above BRANCH_OR_COMMIT")
	moveCmd.Flags().Bool("allow-merged", false, "Allow targeting branches and commits that are already merged upstream")
	moveCmd.Flags().StringP("below", "B", "", "Place <SOURCES> below BRANCH_OR_COMMIT")
	moveCmd.Flags().StringP("branch", "b", "", "Place <SOURCES> on the branch BRANCH")
	moveCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	moveCmd.Flags().Bool("unstack", false, "Unstack <SOURCES> from their current stacks")
	rootCmd.AddCommand(moveCmd)

	carapace.Gen(moveCmd).FlagCompletion(carapace.ActionMap{
		"above":  but.ActionTargets(),
		"below":  but.ActionTargets(),
		"branch": but.ActionLocalBranches(),
	})

	carapace.Gen(moveCmd).PositionalCompletion(
		carapace.Batch(
			but.ActionCommits(),
			but.ActionCliIds(but.CliIdsOpts{Commits: true}),
		).ToA(),
	)
}
