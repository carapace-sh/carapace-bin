package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net"
	"github.com/spf13/cobra"
)

var device_checkpointCmd = &cobra.Command{
	Use:   "checkpoint",
	Short: "Run a command with a configuration checkpoint",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(device_checkpointCmd).Standalone()

	device_checkpointCmd.Flags().String("timeout", "", "checkpoint timeout in seconds")
	deviceCmd.AddCommand(device_checkpointCmd)

	carapace.Gen(device_checkpointCmd).PositionalAnyCompletion(net.ActionDevices(net.AllDevices))
}
