package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var default_security_group_rule_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Add a new security group rule to the default security group template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(default_security_group_rule_createCmd).Standalone()

	default_security_group_rule_createCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	default_security_group_rule_createCmd.Flags().String("description", "", "Set default security group rule description")
	default_security_group_rule_createCmd.Flags().String("dst-port", "", "Destination port, may be a single port or a starting and ending port range: 137:139.")
	default_security_group_rule_createCmd.Flags().Bool("egress", false, "Rule will apply to outgoing network traffic")
	default_security_group_rule_createCmd.Flags().String("ethertype", "", "Ethertype of network traffic (IPv4, IPv6; default: based on IP protocol)")
	default_security_group_rule_createCmd.Flags().String("extra-property", "", "Additional parameters can be passed using this property.")
	default_security_group_rule_createCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	default_security_group_rule_createCmd.Flags().Bool("for-custom-sg", false, "Set this default security group rule to be used in all custom security groups created manually by users")
	default_security_group_rule_createCmd.Flags().Bool("for-default-sg", false, "Set this default security group rule to be used in all default security groups created automatically for each project")
	default_security_group_rule_createCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	default_security_group_rule_createCmd.Flags().String("icmp-code", "", "ICMP code for ICMP IP protocols")
	default_security_group_rule_createCmd.Flags().String("icmp-type", "", "ICMP type for ICMP IP protocols")
	default_security_group_rule_createCmd.Flags().Bool("ingress", false, "Rule will apply to incoming network traffic (default)")
	default_security_group_rule_createCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	default_security_group_rule_createCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	default_security_group_rule_createCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	default_security_group_rule_createCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	default_security_group_rule_createCmd.Flags().String("protocol", "", "IP protocol (ah, dccp, egp, esp, gre, icmp, igmp, ipv66-encap, ipv6-frag, ipv6-icmp, ipv6-nonxt, ipv6-opts, ipv6-route, ospf, pgm, rsvp, sctp, tcp, udp, udplite, vrrp and integer representations [0-255] or any; default: any (all protocols))")
	default_security_group_rule_createCmd.Flags().String("remote-address-group", "", "Remote address group (ID)")
	default_security_group_rule_createCmd.Flags().String("remote-group", "", "Remote security group (ID)")
	default_security_group_rule_createCmd.Flags().String("remote-ip", "", "Remote IP address block (may use CIDR notation; default for IPv4 rule: 0.0.0.0/0, default for IPv6 rule: ::/0)")
	default_security_group_rule_createCmd.Flags().String("variable", "", "==SUPPRESS==")
	default_security_group_ruleCmd.AddCommand(default_security_group_rule_createCmd)
}
