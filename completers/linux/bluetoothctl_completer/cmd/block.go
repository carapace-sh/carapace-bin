package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/bluetoothctl"
	"github.com/spf13/cobra"
)

var blockCmd = &cobra.Command{
	Use:   "block <device>",
	Short: "Block device",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(blockCmd).Standalone()
	rootCmd.AddCommand(blockCmd)
	carapace.Gen(blockCmd).PositionalCompletion(
		bluetoothctl.ActionDevices(bluetoothctl.DeviceOpts{}.Default()),
	)
}
