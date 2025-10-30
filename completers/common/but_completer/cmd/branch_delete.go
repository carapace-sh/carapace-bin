package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/but"
	"github.com/spf13/cobra"
)

var branch_deleteCmd = &cobra.Command{
	Use:   "delete BRANCH_NAME",
	Short: "Deletes a branch from the workspace",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(branch_deleteCmd).Standalone()

	branch_deleteCmd.Flags().BoolP("force", "f", false, "Force deletion without confirmation")
	branch_deleteCmd.Flags().BoolP("help", "h", false, "Print help")
	branchCmd.AddCommand(branch_deleteCmd)

	carapace.Gen(branch_deleteCmd).PositionalCompletion(
		but.ActionLocalBranches(),
	)
}
