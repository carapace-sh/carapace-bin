package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_branch_listAllCmd = &cobra.Command{
	Use:   "list-all",
	Short: "List all branches that can be relevant",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_branch_listAllCmd).Standalone()

	help_branchCmd.AddCommand(help_branch_listAllCmd)
}
