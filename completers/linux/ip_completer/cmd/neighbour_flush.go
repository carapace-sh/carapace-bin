package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net"
	"github.com/spf13/cobra"
)

var neighbour_flushCmd = &cobra.Command{
	Use:   "flush",
	Short: "flush neighbour entries",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neighbour_flushCmd).Standalone()

	neighbourCmd.AddCommand(neighbour_flushCmd)

	carapace.Gen(neighbour_flushCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if len(c.Args) > 0 {
				switch c.Args[len(c.Args)-1] {
				case "dev":
					return net.ActionDevices(net.IncludedDevices{Wifi: true, Ethernet: true})
				case "nud":
					return carapace.ActionValues("permanent", "noarp", "stale", "reachable", "none", "incomplete", "delay", "probe", "failed")
				case "vrf":
					return carapace.ActionValues()
				case "to":
					return net.ActionIpv4Addresses()
				}
			}
			return carapace.ActionValues("dev", "nud", "vrf", "to", "proxy", "unused", "nomaster")
		}),
	)
}
