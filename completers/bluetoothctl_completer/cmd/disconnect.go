package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/bluetoothctl"
	"github.com/spf13/cobra"
)

var disconnectCmd = &cobra.Command{
	Use:   "disconnect <device>",
	Short: "Disconnect device",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(disconnectCmd).Standalone()
	rootCmd.AddCommand(disconnectCmd)
	carapace.Gen(disconnectCmd).PositionalCompletion(
		bluetoothctl.ActionDevices(bluetoothctl.DeviceOpts{Connected: true}),
	)
}
