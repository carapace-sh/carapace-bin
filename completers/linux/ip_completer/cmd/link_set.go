package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/linux/ip_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net"
	"github.com/spf13/cobra"
)

var link_setCmd = &cobra.Command{
	Use:   "set",
	Short: "change device attributes",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(link_setCmd).Standalone()

	linkCmd.AddCommand(link_setCmd)

	carapace.Gen(link_setCmd).PositionalCompletion(
		net.ActionDevices(net.IncludedDevices{Wifi: true, Ethernet: true}),
	)

	carapace.Gen(link_setCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if len(c.Args) > 0 {
				switch c.Args[len(c.Args)-1] {
				case "dev":
					return net.ActionDevices(net.IncludedDevices{Wifi: true, Ethernet: true})
				case "master":
					return net.ActionDevices(net.IncludedDevices{})
				case "vrf":
					return carapace.ActionValues()
				case "netns":
					return carapace.ActionValues()
				case "address":
					return carapace.ActionValues().NoSpace()
				case "arp":
					return carapace.ActionValues("on", "off")
				case "multicast":
					return carapace.ActionValues("on", "off")
				case "dynamic":
					return carapace.ActionValues("on", "off")
				case "group":
					return carapace.ActionValues()
				case "type":
					return action.ActionTypes()
				case "mtu":
					return carapace.ActionValues()
				case "name":
					return carapace.ActionValues().NoSpace()
				case "txqueuelen":
					return carapace.ActionValues()
				}
			}
			return carapace.ActionValues("up", "down", "dev", "master", "nomaster", "vrf", "mtu", "address", "arp", "multicast", "dynamic", "group", "name", "netns", "txqueuelen", "type")
		}),
	)
}
