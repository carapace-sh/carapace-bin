package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/bluetoothctl"
	"github.com/spf13/cobra"
)

var untrustCmd = &cobra.Command{
	Use:   "untrust <device>",
	Short: "Untrust device",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(untrustCmd).Standalone()
	rootCmd.AddCommand(untrustCmd)
	carapace.Gen(untrustCmd).PositionalCompletion(
		bluetoothctl.ActionDevices(bluetoothctl.DeviceOpts{Trusted: true}),
	)
}
