package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/jj"
	"github.com/spf13/cobra"
)

var operation_undoCmd = &cobra.Command{
	Use:   "undo [OPTIONS] [OPERATION]",
	Short: "Create a new operation that undoes an earlier operation",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(operation_undoCmd).Standalone()

	operation_undoCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	operation_undoCmd.Flags().StringSlice("what", nil, "What portions of the local state to restore (can be repeated)")
	operationCmd.AddCommand(operation_undoCmd)

	carapace.Gen(operation_undoCmd).FlagCompletion(carapace.ActionMap{
		"what": carapace.ActionValues("repo", "remote-tracking"),
	})

	carapace.Gen(operation_undoCmd).PositionalCompletion(
		jj.ActionOperations(100),
	)
}
