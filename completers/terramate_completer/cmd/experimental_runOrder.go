package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var experimental_runOrderCmd = &cobra.Command{
	Use:   "run-order",
	Short: "Show the topological ordering of the stacks",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(experimental_runOrderCmd).Standalone()

	experimentalCmd.AddCommand(experimental_runOrderCmd)

	carapace.Gen(experimental_runOrderCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)
}
