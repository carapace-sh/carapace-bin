package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_branch_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Deletes a branch from the workspace",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_branch_deleteCmd).Standalone()

	help_branchCmd.AddCommand(help_branch_deleteCmd)
}
