package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/linux/ip_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net"
	"github.com/spf13/cobra"
)

var link_addCmd = &cobra.Command{
	Use:   "add",
	Short: "add a virtual link",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(link_addCmd).Standalone()

	linkCmd.AddCommand(link_addCmd)

	carapace.Gen(link_addCmd).PositionalCompletion(
		carapace.ActionValues("name"),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return carapace.ActionValues().NoSpace()
		}),
		carapace.ActionValues("type"),
		action.ActionTypes(),
	)

	carapace.Gen(link_addCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if len(c.Args) > 0 {
				switch c.Args[len(c.Args)-1] {
				case "dev":
					return net.ActionDevices(net.IncludedDevices{Wifi: true, Ethernet: true})
				case "type":
					return action.ActionTypes()
				case "master":
					return net.ActionDevices(net.IncludedDevices{})
				case "vrf":
					return carapace.ActionValues()
				case "netns":
					return carapace.ActionValues()
				case "address":
					return carapace.ActionValues().NoSpace()
				}
			}
			return carapace.ActionValues("dev", "type", "master", "vrf", "netns", "address", "mtu", "up", "down", "group", "name")
		}),
	)
}
