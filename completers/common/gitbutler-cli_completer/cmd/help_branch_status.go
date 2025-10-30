package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_branch_statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Provide the current state of all applied virtual branches",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_branch_statusCmd).Standalone()

	help_branchCmd.AddCommand(help_branch_statusCmd)
}
