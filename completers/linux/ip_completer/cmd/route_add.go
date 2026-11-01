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
			if len(c.Args) > 0 {
				switch c.Args[len(c.Args)-1] {
				case "dev":
					return net.ActionDevices(net.IncludedDevices{Wifi: true, Ethernet: true})
				case "via":
					return net.ActionIpv4Addresses()
				case "src":
					return net.ActionIpv4Addresses()
				case "table":
					return carapace.ActionValues("all", "default", "local", "main")
				case "scope":
					return carapace.ActionValues("global", "host", "link", "nowhere", "site")
				case "proto":
					return carapace.ActionValues("boot", "kernel", "redirect", "static")
				case "encap":
					return carapace.ActionValues("bpf", "ioam6", "ip", "mpls", "seg6", "seg6local", "xfrm")
				case "pref":
					return carapace.ActionValues("high", "low", "medium")
				case "to":
					return net.ActionSubnets()
				case "congctl":
					return carapace.ActionValues()
				case "expires":
					return carapace.ActionValues()
				case "initcwnd":
					return carapace.ActionValues()
				case "metric":
					return carapace.ActionValues()
				case "mtu":
					return carapace.ActionValues()
				case "nexthop":
					return carapace.ActionValues()
				case "nhid":
					return carapace.ActionValues()
				case "rtt":
					return carapace.ActionValues()
				case "ttl-propagate":
					return carapace.ActionValues()
				case "window":
					return carapace.ActionValues()
				}
			}
			return carapace.ActionValues("congctl", "dev", "encap", "expires", "initcwnd", "metric", "mtu", "nexthop", "nhid", "onlink", "pref", "proto", "rtt", "scope", "src", "table", "to", "ttl-propagate", "via", "window")
		}),
	)
}
