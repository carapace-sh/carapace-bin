package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/bluetoothctl"
	"github.com/spf13/cobra"
)

var unblockCmd = &cobra.Command{
	Use:   "unblock <device>",
	Short: "Unblock device",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(unblockCmd).Standalone()
	rootCmd.AddCommand(unblockCmd)
	carapace.Gen(unblockCmd).PositionalCompletion(
		bluetoothctl.ActionDevices(bluetoothctl.DeviceOpts{}.Default()),
	)
}
