package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var accelerator_device_profileCmd = &cobra.Command{
	Use:   "profile",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(accelerator_device_profileCmd).Standalone()

	accelerator_deviceCmd.AddCommand(accelerator_device_profileCmd)
}
