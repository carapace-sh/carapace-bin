package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net"
	"github.com/spf13/cobra"
)

var device_disconnectCmd = &cobra.Command{
	Use:   "disconnect",
	Short: "Disconnect devices",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(device_disconnectCmd).Standalone()

	deviceCmd.AddCommand(device_disconnectCmd)

	carapace.Gen(device_disconnectCmd).PositionalAnyCompletion(net.ActionDevices(net.AllDevices))
}
