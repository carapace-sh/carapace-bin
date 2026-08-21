package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var aggregate_removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(aggregate_removeCmd).Standalone()

	aggregateCmd.AddCommand(aggregate_removeCmd)
}
