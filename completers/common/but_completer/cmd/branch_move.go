package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/but"
	"github.com/spf13/cobra"
)

var branch_moveCmd = &cobra.Command{
	Use:   "move",
	Short: "Move a branch on top of another branch, effectively stacking them",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(branch_moveCmd).Standalone()

	branch_moveCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	branchCmd.AddCommand(branch_moveCmd)

	carapace.Gen(branch_moveCmd).PositionalCompletion(
		carapace.Batch(
			but.ActionLocalBranches(),
			but.ActionCliIds(but.CliIdsOpts{Branches: true}),
		).ToA(),
		carapace.Batch(
			but.ActionLocalBranches(),
			but.ActionCliIds(but.CliIdsOpts{Branches: true}),
		).ToA().FilterArgs(),
	)
}
