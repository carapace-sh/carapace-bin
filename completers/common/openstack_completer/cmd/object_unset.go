package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var object_unsetCmd = &cobra.Command{
	Use:   "unset",
	Short: "Unset object properties",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(object_unsetCmd).Standalone()

	object_unsetCmd.Flags().String("property", "", "Property to remove from object (repeat option to remove multiple properties)")
	object_unsetCmd.MarkFlagRequired("property")
	objectCmd.AddCommand(object_unsetCmd)
}
