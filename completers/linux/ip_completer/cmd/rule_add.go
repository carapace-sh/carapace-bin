package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/linux/ip_completer/cmd/action"
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
			if len(c.Args) > 0 {
				switch c.Args[len(c.Args)-1] {
				case "dport":
					return carapace.ActionValues()
				case "dsfield":
					return carapace.ActionValues()
				case "from":
					return net.ActionSubnets()
				case "fwmark":
					return carapace.ActionValues()
				case "goto":
					return carapace.ActionValues()
				case "iif":
					return net.ActionDevices(net.IncludedDevices{Wifi: true, Ethernet: true})
				case "ipproto":
					return action.ActionIpProtocols()
				case "lookup":
					return carapace.ActionValues("local", "main", "default", "all")
				case "map-to":
					return net.ActionIpv4Addresses()
				case "nat":
					return net.ActionIpv4Addresses()
				case "not":
					return carapace.ActionValues("from", "to", "tos", "dsfield", "fwmark", "iif", "oif", "priority", "preference", "order", "l3mdev", "uidrange", "ipproto", "sport", "dport", "tun_id")
				case "oif":
					return net.ActionDevices(net.IncludedDevices{Wifi: true, Ethernet: true})
				case "order":
					return carapace.ActionValues()
				case "preference":
					return carapace.ActionValues()
				case "priority":
					return carapace.ActionValues()
				case "proto":
					return carapace.ActionValues("kernel", "boot", "static")
				case "protocol":
					return carapace.ActionValues("kernel", "boot", "static")
				case "realms":
					return carapace.ActionValues()
				case "sport":
					return carapace.ActionValues()
				case "suppress_ifgroup":
					return carapace.ActionValues()
				case "suppress_prefixlength":
					return carapace.ActionValues()
				case "table":
					return carapace.ActionValues("local", "main", "default", "all")
				case "to":
					return net.ActionSubnets()
				case "tos":
					return carapace.ActionValues()
				case "tun_id":
					return carapace.ActionValues()
				case "type":
					return carapace.ActionValues("unicast", "blackhole", "unreachable", "prohibit", "nat")
				case "uidrange":
					return carapace.ActionValues()
				}
			}
			return carapace.ActionValues("not", "from", "to", "tos", "dsfield", "fwmark", "iif", "oif", "priority", "preference", "order", "l3mdev", "uidrange", "ipproto", "sport", "dport", "tun_id", "table", "lookup", "protocol", "proto", "nat", "map-to", "realms", "goto", "type", "suppress_prefixlength", "suppress_ifgroup")
		}),
	)
}
