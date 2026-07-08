package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net"
	"github.com/spf13/cobra"
)

var rule_showCmd = &cobra.Command{
	Use:     "show",
	Aliases: []string{"list", "lst"},
	Short:   "list rules",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rule_showCmd).Standalone()

	ruleCmd.AddCommand(rule_showCmd)

	carapace.Gen(rule_showCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if len(c.Args) > 0 {
				switch c.Args[len(c.Args)-1] {
				case "dport":
					return carapace.ActionValues()
				case "from":
					return net.ActionSubnets()
				case "fwmark":
					return carapace.ActionValues()
				case "iif":
					return net.ActionDevices(net.IncludedDevices{Wifi: true, Ethernet: true})
				case "ipproto":
					return net.ActionProtocols()
				case "not":
					return carapace.ActionValues()
				case "oif":
					return net.ActionDevices(net.IncludedDevices{Wifi: true, Ethernet: true})
				case "priority":
					return carapace.ActionValues()
				case "sport":
					return carapace.ActionValues()
				case "suppress_ifgroup":
					return carapace.ActionValues()
				case "suppress_prefixlength":
					return carapace.ActionValues()
				case "to":
					return net.ActionSubnets()
				case "tos":
					return carapace.ActionValues()
				case "tun_id":
					return carapace.ActionValues()
				case "uidrange":
					return carapace.ActionValues()
				}
			}
			return carapace.ActionValues("not", "from", "to", "tos", "fwmark", "iif", "oif", "priority", "l3mdev", "uidrange", "ipproto", "sport", "dport", "tun_id", "suppress_prefixlength", "suppress_ifgroup")
		}),
	)
}
