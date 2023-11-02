package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var help_operation_restoreCmd = &cobra.Command{
	Use:   "restore",
	Short: "Create a new operation that restores the repo to an earlier state",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_operation_restoreCmd).Standalone()

	help_operationCmd.AddCommand(help_operation_restoreCmd)
}
