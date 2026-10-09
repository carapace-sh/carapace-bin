package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var firewall_group_policy_unsetCmd = &cobra.Command{
	Use:   "unset",
	Short: "Unset firewall policy properties",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(firewall_group_policy_unsetCmd).Standalone()

	firewall_group_policy_unsetCmd.Flags().Bool("all-firewall-rule", false, "Remove all firewall rules from the firewall policy")
	firewall_group_policy_unsetCmd.Flags().Bool("audited", false, "Disable auditing for the policy")
	firewall_group_policy_unsetCmd.Flags().String("firewall-rule", "", "Remove firewall rule(s) from the firewall policy (name or ID)")
	firewall_group_policy_unsetCmd.Flags().Bool("share", false, "(Deprecated) Use \"firewall policy set --no-share\" instead.")
	firewall_group_policyCmd.AddCommand(firewall_group_policy_unsetCmd)
}
