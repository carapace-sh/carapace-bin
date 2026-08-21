package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var aggregate_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set aggregate properties",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(aggregate_setCmd).Standalone()

	aggregate_setCmd.Flags().String("name", "", "Set aggregate name")
	aggregate_setCmd.Flags().Bool("no-property", false, "Remove all properties from <aggregate> (specify both --property and --no-property to overwrite the current properties)")
	aggregate_setCmd.Flags().String("property", "", "Property to set on <aggregate> (repeat option to set multiple properties)")
	aggregate_setCmd.Flags().String("zone", "", "Set availability zone name")
	aggregateCmd.AddCommand(aggregate_setCmd)
}
