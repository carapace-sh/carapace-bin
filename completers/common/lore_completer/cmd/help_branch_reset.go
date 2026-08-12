package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_branch_resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Reset local latest pointer for a branch",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_branch_resetCmd).Standalone()

	help_branchCmd.AddCommand(help_branch_resetCmd)
}
