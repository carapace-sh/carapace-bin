package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var branch_applyCmd = &cobra.Command{
	Use:   "apply",
	Short: "Apply a branch to the workspace",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(branch_applyCmd).Standalone()

	branch_applyCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	branchCmd.AddCommand(branch_applyCmd)

	carapace.Gen(branch_applyCmd).PositionalCompletion(
		carapace.Batch(
			git.ActionLocalBranches(),
			git.ActionRemoteBranches(""),
		).ToA(),
	)
}
