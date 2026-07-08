package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net"
	"github.com/spf13/cobra"
)

var route_addCmd = &cobra.Command{
	Use:   "add",
	Short: "add a new route",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route_addCmd).Standalone()

	routeCmd.AddCommand(route_addCmd)

	carapace.Gen(route_addCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			switch c.Args[len(c.Args)-1] {
			case "dev":
				return net.ActionDevices(net.IncludedDevices{Wifi: true, Ethernet: true})
			case "via":
				return net.ActionIpv4Addresses()
			case "src":
				return net.ActionIpv4Addresses()
			case "table":
				return carapace.ActionValues("local", "main", "default", "all")
			case "scope":
				return carapace.ActionValues("global", "site", "link", "host", "nowhere")
			case "proto":
				return carapace.ActionValues("kernel", "boot", "static", "redirect")
			case "encap":
				return carapace.ActionValues("mpls", "ip", "bpf", "seg6", "seg6local", "ioam6", "xfrm")
			case "pref":
				return carapace.ActionValues("low", "medium", "high")
			case "to":
				return net.ActionSubnets()
			default:
				return carapace.ActionValues("to", "via", "dev", "src", "table", "scope", "proto", "metric", "mtu", "encap", "pref", "onlink", "nexthop", "nhid", "window", "rtt", "initcwnd", "congctl", "expires", "ttl-propagate")
			}
		}),
	)
}
