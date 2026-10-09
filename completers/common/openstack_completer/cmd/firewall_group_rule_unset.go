package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var firewall_group_rule_unsetCmd = &cobra.Command{
	Use:   "unset",
	Short: "Unset firewall rule properties",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(firewall_group_rule_unsetCmd).Standalone()

	firewall_group_rule_unsetCmd.Flags().Bool("destination-firewall-group", false, "Destination firewall group (name or ID)")
	firewall_group_rule_unsetCmd.Flags().Bool("destination-ip-address", false, "Destination IP address or subnet")
	firewall_group_rule_unsetCmd.Flags().Bool("destination-port", false, "Destination port number or range(integer in [1, 65535] or range like 123:456)")
	firewall_group_rule_unsetCmd.Flags().Bool("enable-rule", false, "(Deprecated) Use \"firewall rule set --disable-rule\" instead.")
	firewall_group_rule_unsetCmd.Flags().Bool("share", false, "(Deprecated) Use \"firewall rule set --no-share\" instead.")
	firewall_group_rule_unsetCmd.Flags().Bool("source-firewall-group", false, "Source firewall group (name or ID)")
	firewall_group_rule_unsetCmd.Flags().Bool("source-ip-address", false, "Source IP address or subnet")
	firewall_group_rule_unsetCmd.Flags().Bool("source-port", false, "Source port number or range(integer in [1, 65535] or range like 123:456)")
	firewall_group_ruleCmd.AddCommand(firewall_group_rule_unsetCmd)
}
