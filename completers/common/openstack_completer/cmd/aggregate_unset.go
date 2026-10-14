package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var aggregate_unsetCmd = &cobra.Command{
	Use:   "unset",
	Short: "Unset aggregate properties",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(aggregate_unsetCmd).Standalone()

	aggregate_unsetCmd.Flags().String("property", "", "Property to remove from aggregate (repeat option to remove multiple properties)")
	aggregateCmd.AddCommand(aggregate_unsetCmd)
}
