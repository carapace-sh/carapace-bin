package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var firewall_group_policy_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set firewall policy properties",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(firewall_group_policy_setCmd).Standalone()

	firewall_group_policy_setCmd.Flags().Bool("audited", false, "Enable auditing for the policy")
	firewall_group_policy_setCmd.Flags().String("description", "", "Description of the firewall policy")
	firewall_group_policy_setCmd.Flags().String("firewall-rule", "", "Firewall rule(s) to apply (name or ID)")
	firewall_group_policy_setCmd.Flags().String("name", "", "Name for the firewall policy")
	firewall_group_policy_setCmd.Flags().Bool("no-audited", false, "Disable auditing for the policy")
	firewall_group_policy_setCmd.Flags().Bool("no-firewall-rule", false, "Remove all firewall rules from firewall policy")
	firewall_group_policy_setCmd.Flags().Bool("no-share", false, "Restrict use of the firewall policy to the current project")
	firewall_group_policy_setCmd.Flags().Bool("share", false, "Share the firewall policy to be used in all projects (by default, it is restricted to be used by the current project).")
	firewall_group_policyCmd.AddCommand(firewall_group_policy_setCmd)
}
