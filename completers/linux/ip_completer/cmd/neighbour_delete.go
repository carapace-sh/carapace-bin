package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net"
	"github.com/spf13/cobra"
)

var neighbour_deleteCmd = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"del"},
	Short:   "delete a neighbour entry",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neighbour_deleteCmd).Standalone()

	neighbourCmd.AddCommand(neighbour_deleteCmd)

	carapace.Gen(neighbour_deleteCmd).PositionalCompletion(
		net.ActionIpv4Addresses(),
	)

	carapace.Gen(neighbour_deleteCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if len(c.Args) > 0 {
				switch c.Args[len(c.Args)-1] {
				case "dev":
					return net.ActionDevices(net.IncludedDevices{Wifi: true, Ethernet: true})
				case "lladdr":
					return carapace.ActionValues().NoSpace()
				case "nud":
					return carapace.ActionValues("permanent", "noarp", "stale", "reachable", "none", "incomplete", "delay", "probe", "failed")
				case "vrf":
					return carapace.ActionValues()
				}
			}
			return carapace.ActionValues("dev", "lladdr", "nud", "proxy", "vrf", "nomaster")
		}),
	)
}
