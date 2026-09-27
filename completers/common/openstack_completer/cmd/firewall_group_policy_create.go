package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var firewall_group_policy_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new firewall policy",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(firewall_group_policy_createCmd).Standalone()

	firewall_group_policy_createCmd.Flags().Bool("audited", false, "Enable auditing for the policy")
	firewall_group_policy_createCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	firewall_group_policy_createCmd.Flags().String("description", "", "Description of the firewall policy")
	firewall_group_policy_createCmd.Flags().String("firewall-rule", "", "Firewall rule(s) to apply (name or ID)")
	firewall_group_policy_createCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	firewall_group_policy_createCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	firewall_group_policy_createCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	firewall_group_policy_createCmd.Flags().Bool("no-audited", false, "Disable auditing for the policy")
	firewall_group_policy_createCmd.Flags().Bool("no-firewall-rule", false, "Unset all firewall rules from firewall policy")
	firewall_group_policy_createCmd.Flags().Bool("no-share", false, "Restrict use of the firewall policy to the current project")
	firewall_group_policy_createCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	firewall_group_policy_createCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	firewall_group_policy_createCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	firewall_group_policy_createCmd.Flags().String("project", "", "Owner's project (name or ID)")
	firewall_group_policy_createCmd.Flags().String("project-domain", "", "Domain the project belongs to (name or ID).")
	firewall_group_policy_createCmd.Flags().Bool("share", false, "Share the firewall policy to be used in all projects (by default, it is restricted to be used by the current project).")
	firewall_group_policy_createCmd.Flags().String("variable", "", "==SUPPRESS==")
	firewall_group_policyCmd.AddCommand(firewall_group_policy_createCmd)
}
