package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var operation_undoCmd = &cobra.Command{
	Use:   "undo",
	Short: "Create a new operation that undoes an earlier operation",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(operation_undoCmd).Standalone()

	operation_undoCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	operation_undoCmd.Flags().StringSlice("what", []string{}, "What portions of the local state to restore (can be repeated)")
	operationCmd.AddCommand(operation_undoCmd)
}
