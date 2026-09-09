package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var security_group_rule_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new security group rule",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(security_group_rule_createCmd).Standalone()

	security_group_rule_createCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	security_group_rule_createCmd.Flags().String("description", "", "Set security group rule description")
	security_group_rule_createCmd.Flags().String("dst-port", "", "Destination port, may be a single port or a starting and ending port range: 137:139.")
	security_group_rule_createCmd.Flags().Bool("egress", false, "Rule applies to outgoing network traffic")
	security_group_rule_createCmd.Flags().String("ethertype", "", "Ethertype of network traffic (IPv4, IPv6; default: based on IP protocol)")
	security_group_rule_createCmd.Flags().String("extra-property", "", "Additional parameters can be passed using this property.")
	security_group_rule_createCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	security_group_rule_createCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	security_group_rule_createCmd.Flags().String("icmp-code", "", "ICMP code for ICMP IP protocols")
	security_group_rule_createCmd.Flags().String("icmp-type", "", "ICMP type for ICMP IP protocols")
	security_group_rule_createCmd.Flags().Bool("ingress", false, "Rule applies to incoming network traffic (default)")
	security_group_rule_createCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	security_group_rule_createCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	security_group_rule_createCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	security_group_rule_createCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	security_group_rule_createCmd.Flags().String("project", "", "Owner's project (name or ID)")
	security_group_rule_createCmd.Flags().String("project-domain", "", "Domain the project belongs to (name or ID).")
	security_group_rule_createCmd.Flags().String("protocol", "", "IP protocol (ah, dccp, egp, esp, gre, icmp, igmp, ipv6-encap, ipv6-frag, ipv6-icmp, ipv6-nonxt, ipv6-opts, ipv6-route, ospf, pgm, rsvp, sctp, tcp, udp, udplite, vrrp and integer representations [0-255] or any; default: any (all protocols))")
	security_group_rule_createCmd.Flags().String("remote-address-group", "", "Remote address group (name or ID)")
	security_group_rule_createCmd.Flags().String("remote-group", "", "Remote security group (name or ID)")
	security_group_rule_createCmd.Flags().String("remote-ip", "", "Remote IP address block (may use CIDR notation; default for IPv4 rule: 0.0.0.0/0, default for IPv6 rule: ::/0)")
	security_group_rule_createCmd.Flags().String("variable", "", "==SUPPRESS==")
	security_group_ruleCmd.AddCommand(security_group_rule_createCmd)
}
