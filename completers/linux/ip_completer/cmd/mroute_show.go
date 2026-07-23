package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net"
	"github.com/spf13/cobra"
)

var mroute_showCmd = &cobra.Command{
	Use:     "show",
	Aliases: []string{"list"},
	Short:   "list multicast routing cache entries",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mroute_showCmd).Standalone()

	mrouteCmd.AddCommand(mroute_showCmd)

	carapace.Gen(mroute_showCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if len(c.Args) > 0 {
				switch c.Args[len(c.Args)-1] {
				case "iif":
					return net.ActionDevices(net.IncludedDevices{Wifi: true, Ethernet: true})
				case "table":
					return carapace.ActionValues("local", "main", "default", "all")
				case "to":
					return net.ActionSubnets()
				case "from":
					return net.ActionSubnets()
				}
			}
			return carapace.ActionValues("to", "from", "iif", "table")
		}),
	)
}
