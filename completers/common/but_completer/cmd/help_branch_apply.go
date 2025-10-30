package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_branch_applyCmd = &cobra.Command{
	Use:   "apply",
	Short: "Add a branch to the workspace",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_branch_applyCmd).Standalone()

	help_branchCmd.AddCommand(help_branch_applyCmd)
}
