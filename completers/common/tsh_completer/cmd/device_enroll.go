package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var device_enrollCmd = &cobra.Command{
	Use:   "enroll",
	Short: "Enroll this device as a trusted device. Requires Teleport Enterprise.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(device_enrollCmd).Standalone()

	device_enrollCmd.Flags().String("token", "", "Device enrollment token")
	device_enrollCmd.MarkFlagRequired("token")
	deviceCmd.AddCommand(device_enrollCmd)
}
