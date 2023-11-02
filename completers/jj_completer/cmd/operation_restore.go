package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var operation_restoreCmd = &cobra.Command{
	Use:   "restore",
	Short: "Create a new operation that restores the repo to an earlier state",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(operation_restoreCmd).Standalone()

	operation_restoreCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	operation_restoreCmd.Flags().StringSlice("what", []string{}, "What portions of the local state to restore (can be repeated)")
	operationCmd.AddCommand(operation_restoreCmd)
}
