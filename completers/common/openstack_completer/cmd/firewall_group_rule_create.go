package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var firewall_group_rule_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new firewall rule",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(firewall_group_rule_createCmd).Standalone()

	firewall_group_rule_createCmd.Flags().String("action", "", "Action for the firewall rule")
	firewall_group_rule_createCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	firewall_group_rule_createCmd.Flags().String("description", "", "Description of the firewall rule")
	firewall_group_rule_createCmd.Flags().String("destination-firewall-group", "", "Destination firewall group (name or ID)")
	firewall_group_rule_createCmd.Flags().String("destination-ip-address", "", "Destination IP address or subnet")
	firewall_group_rule_createCmd.Flags().String("destination-port", "", "Destination port number or range(integer in [1, 65535] or range like 123:456)")
	firewall_group_rule_createCmd.Flags().Bool("disable-rule", false, "Disable this rule")
	firewall_group_rule_createCmd.Flags().Bool("enable-rule", false, "Enable this rule (default is enabled)")
	firewall_group_rule_createCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	firewall_group_rule_createCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	firewall_group_rule_createCmd.Flags().String("ip-version", "", "Set IP version 4 or 6 (default is 4)")
	firewall_group_rule_createCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	firewall_group_rule_createCmd.Flags().String("name", "", "(Deprecated, please pass name as a positional argument) Name of the firewall rule")
	firewall_group_rule_createCmd.Flags().Bool("no-destination-firewall-group", false, "No associated destination firewall group")
	firewall_group_rule_createCmd.Flags().Bool("no-destination-ip-address", false, "Detach destination IP address")
	firewall_group_rule_createCmd.Flags().Bool("no-destination-port", false, "Detach destination port number or range")
	firewall_group_rule_createCmd.Flags().Bool("no-share", false, "Restrict use of the firewall rule to the current project")
	firewall_group_rule_createCmd.Flags().Bool("no-source-firewall-group", false, "No associated source firewall group")
	firewall_group_rule_createCmd.Flags().Bool("no-source-ip-address", false, "Detach source IP address")
	firewall_group_rule_createCmd.Flags().Bool("no-source-port", false, "Detach source port number or range")
	firewall_group_rule_createCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	firewall_group_rule_createCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	firewall_group_rule_createCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	firewall_group_rule_createCmd.Flags().String("project", "", "Owner's project (name or ID)")
	firewall_group_rule_createCmd.Flags().String("project-domain", "", "Domain the project belongs to (name or ID).")
	firewall_group_rule_createCmd.Flags().String("protocol", "", "IP protocol (ah, dccp, egp, esp, gre, icmp, igmp, ipv6-encap, ipv6-frag, ipv6-icmp, ipv6-nonxt, ipv6-opts, ipv6-route, ospf, pgm, rsvp, sctp, tcp, udp, udplite, vrrp and integer representations [0-255] or any; default: any (all protocols))")
	firewall_group_rule_createCmd.Flags().Bool("share", false, "Share the firewall rule to be used in all projects (by default, it is restricted to be used by the current project).")
	firewall_group_rule_createCmd.Flags().String("source-firewall-group", "", "Source firewall group (name or ID)")
	firewall_group_rule_createCmd.Flags().String("source-ip-address", "", "Source IP address or subnet")
	firewall_group_rule_createCmd.Flags().String("source-port", "", "Source port number or range (integer in [1, 65535] or range like 123:456)")
	firewall_group_rule_createCmd.Flags().String("variable", "", "==SUPPRESS==")
	firewall_group_ruleCmd.AddCommand(firewall_group_rule_createCmd)
}
