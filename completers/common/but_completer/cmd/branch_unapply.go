package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/but"
	"github.com/spf13/cobra"
)

var branch_unapplyCmd = &cobra.Command{
	Use:   "unapply BRANCH_NAME",
	Short: "Unapply a branch from the workspace",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(branch_unapplyCmd).Standalone()

	branch_unapplyCmd.Flags().BoolP("force", "f", false, "Force unapply without confirmation")
	branch_unapplyCmd.Flags().BoolP("help", "h", false, "Print help")
	branchCmd.AddCommand(branch_unapplyCmd)

	carapace.Gen(branch_unapplyCmd).PositionalCompletion(
		but.ActionLocalBranches(),
	)
}
