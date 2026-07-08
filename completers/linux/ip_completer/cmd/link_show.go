package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/linux/ip_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net"
	"github.com/spf13/cobra"
)

var link_showCmd = &cobra.Command{
	Use:     "show",
	Aliases: []string{"list"},
	Short:   "display device attributes",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(link_showCmd).Standalone()

	linkCmd.AddCommand(link_showCmd)

	carapace.Gen(link_showCmd).PositionalCompletion(
		net.ActionDevices(net.IncludedDevices{Wifi: true, Ethernet: true}),
	)

	carapace.Gen(link_showCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			switch c.Args[len(c.Args)-1] {
			case "dev":
				return net.ActionDevices(net.IncludedDevices{Wifi: true, Ethernet: true})
			case "master":
				return net.ActionDevices(net.IncludedDevices{})
			case "vrf":
				return carapace.ActionValues()
			case "type":
				return action.ActionTypes()
			default:
				return carapace.ActionValues("dev", "up", "master", "nomaster", "vrf", "type", "group")
			}
		}),
	)
}
