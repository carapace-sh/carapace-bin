package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/bluetoothctl"
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove <device>",
	Short: "Remove device",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(removeCmd).Standalone()
	rootCmd.AddCommand(removeCmd)
	carapace.Gen(removeCmd).PositionalCompletion(
		bluetoothctl.ActionDevices(bluetoothctl.DeviceOpts{}.Default()),
	)
}
