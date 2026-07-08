package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net"
	"github.com/spf13/cobra"
)

var route_flushCmd = &cobra.Command{
	Use:   "flush",
	Short: "flush routes",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route_flushCmd).Standalone()

	routeCmd.AddCommand(route_flushCmd)

	carapace.Gen(route_flushCmd).PositionalAnyCompletion(
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
			case "to":
				return net.ActionSubnets()
			case "from":
				return net.ActionSubnets()
			default:
				return carapace.ActionValues("to", "from", "via", "dev", "src", "table", "scope", "proto", "metric")
			}
		}),
	)
}
