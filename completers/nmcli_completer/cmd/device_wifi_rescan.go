package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/nmcli_completer/cmd/action"
	"github.com/spf13/cobra"
)

var device_wifi_rescanCmd = &cobra.Command{
	Use:   "rescan",
	Short: "Request that NetworkManager immediately re-scan for available access points",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(device_wifi_rescanCmd).Standalone()

	device_wifiCmd.AddCommand(device_wifi_rescanCmd)

	carapace.Gen(device_wifi_rescanCmd).PositionalCompletion(
		carapace.ActionValues("ifname"),
		action.ActionDevices("wifi"),
	)

	carapace.Gen(device_wifi_rescanCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(args []string) carapace.Action {
			if len(args)%2 == 0 {
				return carapace.ActionValues("ssid")
			} else {
				return action.ActionSsids()
			}
		}),
	)
}
