package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/but"
	"github.com/spf13/cobra"
)

var landCmd = &cobra.Command{
	Use:     "land",
	Short:   "Land a branch directly onto the target branch",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: "server interactions",
}

func init() {
	carapace.Gen(landCmd).Standalone()

	landCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	landCmd.Flags().Bool("no-ff", false, "Always create a merge commit, even when the branch can be fast-forwarded")
	landCmd.Flags().Bool("whole-stack", false, "Land the entire stack: BRANCH must be the top segment, and the segments below it are published to the target along with it")
	landCmd.Flags().Bool("yes", false, "Skip the confirmation prompt")
	rootCmd.AddCommand(landCmd)

	carapace.Gen(landCmd).PositionalCompletion(
		carapace.Batch(
			but.ActionCliIds(but.CliIdsOpts{Branches: true, Stacks: true}),
			but.ActionLocalBranches(),
		).ToA(),
	)
}
