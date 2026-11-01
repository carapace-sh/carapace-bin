package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net"
	"github.com/spf13/cobra"
)

var maddress_deleteCmd = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"del"},
	Short:   "delete a static multicast address",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(maddress_deleteCmd).Standalone()

	maddressCmd.AddCommand(maddress_deleteCmd)

	carapace.Gen(maddress_deleteCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if len(c.Args) > 0 {
				switch c.Args[len(c.Args)-1] {
				case "dev":
					return net.ActionDevices(net.IncludedDevices{Wifi: true, Ethernet: true})
				}
			}
			return carapace.ActionValues("dev")
		}),
	)
}
