package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/jj"
	"github.com/spf13/cobra"
)

var operation_abandonCmd = &cobra.Command{
	Use:   "abandon",
	Short: "Abandon operation history",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(operation_abandonCmd).Standalone()

	operation_abandonCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	operationCmd.AddCommand(operation_abandonCmd)

	carapace.Gen(operation_abandonCmd).PositionalCompletion(
		jj.ActionOperations(100),
	)
}
