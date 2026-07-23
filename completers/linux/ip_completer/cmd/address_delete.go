package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/linux/ip_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net"
	"github.com/spf13/cobra"
)

var address_deleteCmd = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"del"},
	Short:   "delete protocol address",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(address_deleteCmd).Standalone()

	addressCmd.AddCommand(address_deleteCmd)

	carapace.Gen(address_deleteCmd).PositionalCompletion(
		net.ActionIpv4Addresses(),
		carapace.ActionValues("dev"),
		net.ActionDevices(net.IncludedDevices{Wifi: true, Ethernet: true}),
	)

	carapace.Gen(address_deleteCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionConfFlags().Invoke(c).Filter(c.Args[2:]...).ToA()
		}),
	)
}
