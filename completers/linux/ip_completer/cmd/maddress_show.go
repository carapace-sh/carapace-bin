package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net"
	"github.com/spf13/cobra"
)

var maddress_showCmd = &cobra.Command{
	Use:     "show",
	Aliases: []string{"list"},
	Short:   "list multicast addresses",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(maddress_showCmd).Standalone()

	maddressCmd.AddCommand(maddress_showCmd)

	carapace.Gen(maddress_showCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			switch c.Args[len(c.Args)-1] {
			case "dev":
				return net.ActionDevices(net.IncludedDevices{Wifi: true, Ethernet: true})
			default:
				return carapace.ActionValues("dev")
			}
		}),
	)
}
