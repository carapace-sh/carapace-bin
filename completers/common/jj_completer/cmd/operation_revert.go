package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var operation_revertCmd = &cobra.Command{
	Use:   "revert",
	Short: "Create a new operation that reverts an earlier operation",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(operation_revertCmd).Standalone()

	operation_revertCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	operation_revertCmd.Flags().StringSlice("what", nil, "What portions of the local state to restore (can be repeated)")
	operationCmd.AddCommand(operation_revertCmd)

	carapace.Gen(operation_revertCmd).FlagCompletion(carapace.ActionMap{
		"what": carapace.ActionValues("repo", "remote-tracking"),
	})

	// TODO operation completion
}
