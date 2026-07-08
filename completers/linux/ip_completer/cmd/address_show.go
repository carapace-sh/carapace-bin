package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/linux/ip_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net"
	"github.com/spf13/cobra"
)

var address_showCmd = &cobra.Command{
	Use:     "show",
	Aliases: []string{"list"},
	Short:   "look at protocol addresses",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(address_showCmd).Standalone()

	addressCmd.AddCommand(address_showCmd)

	carapace.Gen(address_showCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if len(c.Args) > 0 {
				switch c.Args[len(c.Args)-1] {
				case "dev":
					return net.ActionDevices(net.IncludedDevices{Wifi: true, Ethernet: true})
				case "scope":
					return carapace.ActionValues("global", "site", "link", "host", "nowhere")
				case "to":
					return net.ActionSubnets()
				case "master":
					return net.ActionDevices(net.IncludedDevices{})
				case "vrf":
					return net.ActionDevices(net.IncludedDevices{})
				case "type":
					return action.ActionTypes()
				}
			}
			return carapace.ActionValues("dev", "scope", "to", "label", "up", "master", "nomaster", "type", "vrf")
		}),
	)
}
