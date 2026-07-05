package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net"
	"github.com/spf13/cobra"
)

var device_upCmd = &cobra.Command{
	Use:   "up",
	Short: "Connect the device",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(device_upCmd).Standalone()

	deviceCmd.AddCommand(device_upCmd)

	carapace.Gen(device_upCmd).PositionalCompletion(
		net.ActionDevices(net.AllDevices),
	)
}
