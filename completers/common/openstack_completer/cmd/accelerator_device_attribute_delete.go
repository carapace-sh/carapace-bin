package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var accelerator_device_attribute_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete attribute(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(accelerator_device_attribute_deleteCmd).Standalone()

	accelerator_device_attributeCmd.AddCommand(accelerator_device_attribute_deleteCmd)
}
