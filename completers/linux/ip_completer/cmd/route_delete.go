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
					return carapace.ActionValues("local", "main", "default", "all")
				case "scope":
					return carapace.ActionValues("global", "site", "link", "host", "nowhere")
				case "proto":
					return carapace.ActionValues("kernel", "boot", "static", "redirect")
				case "to":
					return net.ActionSubnets()
				case "encap":
					return carapace.ActionValues("mpls", "ip", "bpf", "seg6", "seg6local", "ioam6", "xfrm")
				case "pref":
					return carapace.ActionValues("low", "medium", "high")
				}
			}
			return carapace.ActionValues("to", "via", "dev", "src", "table", "scope", "proto", "metric", "encap", "pref", "nexthop", "nhid")
		}),
	)
}
