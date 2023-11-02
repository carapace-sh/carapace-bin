package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var help_branch_forgetCmd = &cobra.Command{
	Use:   "forget",
	Short: "Forget everything about a branch, including its local and remote targets",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_branch_forgetCmd).Standalone()

	help_branchCmd.AddCommand(help_branch_forgetCmd)
}
