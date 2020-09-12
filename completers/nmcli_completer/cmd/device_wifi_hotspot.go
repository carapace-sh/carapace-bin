package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/nmcli_completer/cmd/action"
	"github.com/spf13/cobra"
)

var device_wifi_hotspotCmd = &cobra.Command{
	Use:   "hotspot",
	Short: "Create a Wi-Fi hotspot",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(device_wifi_hotspotCmd).Standalone()

	device_wifiCmd.AddCommand(device_wifi_hotspotCmd)

	carapace.Gen(device_wifi_hotspotCmd).PositionalCompletion(
		carapace.ActionValues("ifname"),
		action.ActionDevices("wifi"),
		carapace.ActionValues("con-name"),
		carapace.ActionValues(""),
		carapace.ActionValues("ssid"),
		carapace.ActionValues(""),
		carapace.ActionValues("band"),
		carapace.ActionValues("a", "bg"),
		carapace.ActionValues("channel"),
		carapace.ActionValues(""),
		carapace.ActionValues("password"),
		carapace.ActionValues(""),
	)
}
