package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var firewall_group_rule_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set firewall rule properties",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(firewall_group_rule_setCmd).Standalone()

	firewall_group_rule_setCmd.Flags().String("action", "", "Action for the firewall rule")
	firewall_group_rule_setCmd.Flags().String("description", "", "Description of the firewall rule")
	firewall_group_rule_setCmd.Flags().String("destination-firewall-group", "", "Destination firewall group (name or ID)")
	firewall_group_rule_setCmd.Flags().String("destination-ip-address", "", "Destination IP address or subnet")
	firewall_group_rule_setCmd.Flags().String("destination-port", "", "Destination port number or range(integer in [1, 65535] or range like 123:456)")
	firewall_group_rule_setCmd.Flags().Bool("disable-rule", false, "Disable this rule")
	firewall_group_rule_setCmd.Flags().Bool("enable-rule", false, "Enable this rule (default is enabled)")
	firewall_group_rule_setCmd.Flags().String("ip-version", "", "Set IP version 4 or 6 (default is 4)")
	firewall_group_rule_setCmd.Flags().String("name", "", "Name of the firewall rule")
	firewall_group_rule_setCmd.Flags().Bool("no-destination-firewall-group", false, "No associated destination firewall group")
	firewall_group_rule_setCmd.Flags().Bool("no-destination-ip-address", false, "Detach destination IP address")
	firewall_group_rule_setCmd.Flags().Bool("no-destination-port", false, "Detach destination port number or range")
	firewall_group_rule_setCmd.Flags().Bool("no-share", false, "Restrict use of the firewall rule to the current project")
	firewall_group_rule_setCmd.Flags().Bool("no-source-firewall-group", false, "No associated source firewall group")
	firewall_group_rule_setCmd.Flags().Bool("no-source-ip-address", false, "Detach source IP address")
	firewall_group_rule_setCmd.Flags().Bool("no-source-port", false, "Detach source port number or range")
	firewall_group_rule_setCmd.Flags().String("protocol", "", "IP protocol (ah, dccp, egp, esp, gre, icmp, igmp, ipv6-encap, ipv6-frag, ipv6-icmp, ipv6-nonxt, ipv6-opts, ipv6-route, ospf, pgm, rsvp, sctp, tcp, udp, udplite, vrrp and integer representations [0-255] or any; default: any (all protocols))")
	firewall_group_rule_setCmd.Flags().Bool("share", false, "Share the firewall rule to be used in all projects (by default, it is restricted to be used by the current project).")
	firewall_group_rule_setCmd.Flags().String("source-firewall-group", "", "Source firewall group (name or ID)")
	firewall_group_rule_setCmd.Flags().String("source-ip-address", "", "Source IP address or subnet")
	firewall_group_rule_setCmd.Flags().String("source-port", "", "Source port number or range (integer in [1, 65535] or range like 123:456)")
	firewall_group_ruleCmd.AddCommand(firewall_group_rule_setCmd)
}
