package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/bluetoothctl"
	"github.com/spf13/cobra"
)

var connectCmd = &cobra.Command{
	Use:   "connect <device>",
	Short: "Connect device",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connectCmd).Standalone()
	rootCmd.AddCommand(connectCmd)
	carapace.Gen(connectCmd).PositionalCompletion(
		bluetoothctl.ActionDevices(bluetoothctl.DeviceOpts{}.Default()),
	)
}
