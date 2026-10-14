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

	device_enrollCmd.Flags().Bool("current-device", false, "Attempts to register and enroll the current device. Requires device admin privileges.")
	device_enrollCmd.Flags().Bool("no-current-device", false, "Attempts to register and enroll the current device. Requires device admin privileges.")
	device_enrollCmd.Flags().String("token", "", "Device enrollment token.")
	device_enrollCmd.Flag("no-current-device").Hidden = true
	deviceCmd.AddCommand(device_enrollCmd)
}
