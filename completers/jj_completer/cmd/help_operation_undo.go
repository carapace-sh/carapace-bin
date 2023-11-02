package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var help_operation_undoCmd = &cobra.Command{
	Use:   "undo",
	Short: "Create a new operation that undoes an earlier operation",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_operation_undoCmd).Standalone()

	help_operationCmd.AddCommand(help_operation_undoCmd)
}
