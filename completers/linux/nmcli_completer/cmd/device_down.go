package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net"
	"github.com/spf13/cobra"
)

var device_downCmd = &cobra.Command{
	Use:   "down",
	Short: "Disconnect device and prevent auto-activation",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(device_downCmd).Standalone()

	deviceCmd.AddCommand(device_downCmd)

	carapace.Gen(device_downCmd).PositionalAnyCompletion(net.ActionDevices(net.AllDevices))
}
