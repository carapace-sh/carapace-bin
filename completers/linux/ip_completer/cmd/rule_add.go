package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net"
	"github.com/spf13/cobra"
)

var rule_addCmd = &cobra.Command{
	Use:   "add",
	Short: "insert a new rule",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rule_addCmd).Standalone()

	ruleCmd.AddCommand(rule_addCmd)

	carapace.Gen(rule_addCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			switch c.Args[len(c.Args)-1] {
			case "from":
				return net.ActionSubnets()
			case "to":
				return net.ActionSubnets()
			case "iif":
				return net.ActionDevices(net.IncludedDevices{Wifi: true, Ethernet: true})
			case "oif":
				return net.ActionDevices(net.IncludedDevices{Wifi: true, Ethernet: true})
			case "table":
				return carapace.ActionValues("local", "main", "default", "all")
			case "proto":
				return carapace.ActionValues("kernel", "boot", "static")
			case "not":
				return carapace.ActionValues()
			case "ipproto":
				return net.ActionProtocols()
			case "type":
				return carapace.ActionValues("unicast", "blackhole", "unreachable", "prohibit", "nat")
			case "nat":
				return net.ActionIpv4Addresses()
			default:
				return carapace.ActionValues("not", "from", "to", "tos", "fwmark", "iif", "oif", "priority", "l3mdev", "uidrange", "ipproto", "sport", "dport", "tun_id", "table", "protocol", "nat", "realms", "goto", "type", "suppress_prefixlength", "suppress_ifgroup")
			}
		}),
	)
}
