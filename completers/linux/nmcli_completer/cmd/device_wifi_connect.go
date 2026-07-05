package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/linux/nmcli_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net"
	"github.com/spf13/cobra"
)

var device_wifi_wifi_connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "Connect to a Wi-Fi network",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(device_wifi_wifi_connectCmd).Standalone()

	device_wifiCmd.AddCommand(device_wifi_wifi_connectCmd)

	carapace.Gen(device_wifi_wifi_connectCmd).PositionalCompletion(
		net.ActionSsids(),
	)

	carapace.Gen(device_wifi_wifi_connectCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if len(c.Args) > 0 {
				switch c.Args[len(c.Args)-1] {
				case "ifname":
					return net.ActionDevices(net.IncludedDevices{Wifi: true})
				case "bssid":
					return net.ActionBssids()
				case "wep-key-type":
					return carapace.ActionValues("key", "phrase")
				case "private", "hidden":
					return action.ActionYesNo()
				case "password", "name":
					return carapace.ActionValues()
				}
			}
			return carapace.ActionValues("password", "wep-key-type", "ifname", "bssid", "name", "private", "hidden")
		}),
	)
}
