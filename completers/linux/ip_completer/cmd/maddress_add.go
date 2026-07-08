package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net"
	"github.com/spf13/cobra"
)

var maddress_addCmd = &cobra.Command{
	Use:   "add",
	Short: "add a static multicast address",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(maddress_addCmd).Standalone()

	maddressCmd.AddCommand(maddress_addCmd)

	carapace.Gen(maddress_addCmd).PositionalAnyCompletion(
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
