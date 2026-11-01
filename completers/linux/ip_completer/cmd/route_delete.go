package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net"
	"github.com/spf13/cobra"
)

var route_deleteCmd = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"del"},
	Short:   "delete a route",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route_deleteCmd).Standalone()

	routeCmd.AddCommand(route_deleteCmd)

	carapace.Gen(route_deleteCmd).PositionalAnyCompletion(
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
				case "to":
					return net.ActionSubnets()
				case "encap":
					return carapace.ActionValues("bpf", "ioam6", "ip", "mpls", "seg6", "seg6local", "xfrm")
				case "pref":
					return carapace.ActionValues("high", "low", "medium")
				case "metric":
					return carapace.ActionValues()
				case "nexthop":
					return carapace.ActionValues()
				case "nhid":
					return carapace.ActionValues()
				}
			}
			return carapace.ActionValues("dev", "encap", "metric", "nexthop", "nhid", "pref", "proto", "scope", "src", "table", "to", "via")
		}),
	)
}
