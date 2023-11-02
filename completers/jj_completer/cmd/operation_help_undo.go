package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var operation_help_undoCmd = &cobra.Command{
	Use:   "undo",
	Short: "Create a new operation that undoes an earlier operation",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(operation_help_undoCmd).Standalone()

	operation_helpCmd.AddCommand(operation_help_undoCmd)
}
