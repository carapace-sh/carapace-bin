package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net"
	"github.com/spf13/cobra"
)

var route_showCmd = &cobra.Command{
	Use:     "show",
	Aliases: []string{"list"},
	Short:   "list routes",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route_showCmd).Standalone()

	routeCmd.AddCommand(route_showCmd)

	carapace.Gen(route_showCmd).PositionalAnyCompletion(
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
				case "from":
					return net.ActionSubnets()
				case "metric":
					return carapace.ActionValues()
				}
			}
			return carapace.ActionValues("to", "from", "via", "dev", "src", "table", "scope", "proto", "metric", "cache", "cloned")
		}),
	)
}
