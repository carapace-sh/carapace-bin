package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net"
	"github.com/spf13/cobra"
)

var device_wifi_showPasswordCmd = &cobra.Command{
	Use:   "show-password",
	Short: "Show details of active Wi-Fi networks including secrets",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(device_wifi_showPasswordCmd).Standalone()

	device_wifiCmd.AddCommand(device_wifi_showPasswordCmd)

	carapace.Gen(device_wifi_showPasswordCmd).PositionalCompletion(
		carapace.ActionValues("ifname"),
		net.ActionDevices(net.IncludedDevices{Wifi: true}),
	)
}
