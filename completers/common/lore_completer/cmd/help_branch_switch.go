package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_branch_switchCmd = &cobra.Command{
	Use:   "switch",
	Short: "Switch to a different branch",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_branch_switchCmd).Standalone()

	help_branchCmd.AddCommand(help_branch_switchCmd)
}
