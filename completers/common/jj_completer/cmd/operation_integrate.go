package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/jj"
	"github.com/spf13/cobra"
)

var operation_integrateCmd = &cobra.Command{
	Use:   "integrate",
	Short: "Make an operation part of the operation log",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(operation_integrateCmd).Standalone()

	operation_integrateCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	operationCmd.AddCommand(operation_integrateCmd)

	carapace.Gen(operation_integrateCmd).PositionalCompletion(
		jj.ActionOperations(100),
	)
}
