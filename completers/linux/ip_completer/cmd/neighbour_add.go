package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net"
	"github.com/spf13/cobra"
)

var neighbour_addCmd = &cobra.Command{
	Use:   "add",
	Short: "add a new neighbour entry",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neighbour_addCmd).Standalone()

	neighbourCmd.AddCommand(neighbour_addCmd)

	carapace.Gen(neighbour_addCmd).PositionalCompletion(
		net.ActionIpv4Addresses(),
	)

	carapace.Gen(neighbour_addCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if len(c.Args) > 0 {
				switch c.Args[len(c.Args)-1] {
				case "dev":
					return net.ActionDevices(net.IncludedDevices{Wifi: true, Ethernet: true})
				case "lladdr":
					return carapace.ActionValues().NoSpace()
				case "nud":
					return carapace.ActionValues("delay", "failed", "incomplete", "noarp", "none", "permanent", "probe", "reachable", "stale")
				case "vrf":
					return carapace.ActionValues()
				}
			}
			return carapace.ActionValues("dev", "extern_learn", "extern_valid", "lladdr", "managed", "nomaster", "nud", "proxy", "router", "use", "vrf")
		}),
	)
}
