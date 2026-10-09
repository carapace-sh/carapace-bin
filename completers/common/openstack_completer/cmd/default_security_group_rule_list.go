package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var default_security_group_rule_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List security group rules used for new default security groups.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(default_security_group_rule_listCmd).Standalone()

	default_security_group_rule_listCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	default_security_group_rule_listCmd.Flags().Bool("egress", false, "List only default rules which will be applied to outgoing network traffic")
	default_security_group_rule_listCmd.Flags().String("ethertype", "", "List default rules by the Ethertype (IPv4 or IPv6)")
	default_security_group_rule_listCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	default_security_group_rule_listCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	default_security_group_rule_listCmd.Flags().Bool("ingress", false, "List only default rules which will be applied to incoming network traffic")
	default_security_group_rule_listCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	default_security_group_rule_listCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	default_security_group_rule_listCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	default_security_group_rule_listCmd.Flags().String("protocol", "", "List only default rules with the specified IP protocol (ah, dhcp, egp, esp, gre, icmp, igmp, ipv6-encap, ipv6-frag, ipv6-icmp, ipv6-nonxt, ipv6-opts, ipv6-route, ospf, pgm, rsvp, sctp, tcp, udp, udplite, vrrp and integer representations [0-255] or any; default: any (all protocols))")
	default_security_group_rule_listCmd.Flags().String("quote", "", "when to include quotes, defaults to nonnumeric")
	default_security_group_rule_listCmd.Flags().Bool("sort-ascending", false, "sort the column(s) in ascending order")
	default_security_group_rule_listCmd.Flags().String("sort-column", "", "specify the column(s) to sort the data (columns specified first have a priority, non-existing columns are ignored), can be repeated")
	default_security_group_rule_listCmd.Flags().Bool("sort-descending", false, "sort the column(s) in descending order")
	default_security_group_ruleCmd.AddCommand(default_security_group_rule_listCmd)
}
