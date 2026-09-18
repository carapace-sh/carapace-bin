package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var object_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set object properties",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(object_setCmd).Standalone()

	object_setCmd.Flags().String("property", "", "Set a property on this object (repeat option to set multiple properties)")
	object_setCmd.MarkFlagRequired("property")
	objectCmd.AddCommand(object_setCmd)
}
