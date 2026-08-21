package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var security_group_rule_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List security group rules",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(security_group_rule_listCmd).Standalone()

	security_group_rule_listCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	security_group_rule_listCmd.Flags().Bool("egress", false, "List only rules applied to outgoing network traffic")
	security_group_rule_listCmd.Flags().String("ethertype", "", "List only rules with the specified Ethertype (IPv4 or IPv6)")
	security_group_rule_listCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	security_group_rule_listCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	security_group_rule_listCmd.Flags().Bool("ingress", false, "List only rules applied to incoming network traffic")
	security_group_rule_listCmd.Flags().String("limit", "", "The maximum number of entries to return per page.")
	security_group_rule_listCmd.Flags().Bool("long", false, "**Deprecated** This argument is no longer needed")
	security_group_rule_listCmd.Flags().String("marker", "", "The first position in the collection to return results from.")
	security_group_rule_listCmd.Flags().String("max-items", "", "The maximum number of entries to return in total, paging through multiple requests if needed.")
	security_group_rule_listCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	security_group_rule_listCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	security_group_rule_listCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	security_group_rule_listCmd.Flags().String("project", "", "List only rules with the specified project (name or ID)")
	security_group_rule_listCmd.Flags().String("project-domain", "", "Domain the project belongs to (name or ID).")
	security_group_rule_listCmd.Flags().String("protocol", "", "List only rules with the specified IP protocol (ah, dhcp, egp, esp, gre, icmp, igmp, ipv6-encap, ipv6-frag, ipv6-icmp, ipv6-nonxt, ipv6-opts, ipv6-route, ospf, pgm, rsvp, sctp, tcp, udp, udplite, vrrp and integer representations [0-255] or any; default: any (all protocols))")
	security_group_rule_listCmd.Flags().String("quote", "", "when to include quotes, defaults to nonnumeric")
	security_group_rule_listCmd.Flags().Bool("sort-ascending", false, "sort the column(s) in ascending order")
	security_group_rule_listCmd.Flags().String("sort-column", "", "specify the column(s) to sort the data (columns specified first have a priority, non-existing columns are ignored), can be repeated")
	security_group_rule_listCmd.Flags().Bool("sort-descending", false, "sort the column(s) in descending order")
	security_group_ruleCmd.AddCommand(security_group_rule_listCmd)
}
