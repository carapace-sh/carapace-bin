package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/bluetoothctl"
	"github.com/spf13/cobra"
)

var cancelPairingCmd = &cobra.Command{
	Use:   "cancel-pairing <device>",
	Short: "Cancel pairing with device",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cancelPairingCmd).Standalone()
	rootCmd.AddCommand(cancelPairingCmd)
	carapace.Gen(cancelPairingCmd).PositionalCompletion(
		bluetoothctl.ActionDevices(bluetoothctl.DeviceOpts{Paired: true}),
	)
}
