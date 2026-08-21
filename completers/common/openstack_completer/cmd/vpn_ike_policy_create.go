package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var vpn_ike_policy_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create an IKE policy",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(vpn_ike_policy_createCmd).Standalone()

	vpn_ike_policy_createCmd.Flags().String("auth-algorithm", "", "Authentication algorithm")
	vpn_ike_policy_createCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	vpn_ike_policy_createCmd.Flags().String("description", "", "Description of the IKE policy")
	vpn_ike_policy_createCmd.Flags().String("encryption-algorithm", "", "Encryption algorithm")
	vpn_ike_policy_createCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	vpn_ike_policy_createCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	vpn_ike_policy_createCmd.Flags().String("ike-version", "", "IKE version for the policy")
	vpn_ike_policy_createCmd.Flags().String("lifetime", "", "IKE lifetime attributes.")
	vpn_ike_policy_createCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	vpn_ike_policy_createCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	vpn_ike_policy_createCmd.Flags().String("pfs", "", "Perfect Forward Secrecy")
	vpn_ike_policy_createCmd.Flags().String("phase1-negotiation-mode", "", "IKE Phase1 negotiation mode")
	vpn_ike_policy_createCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	vpn_ike_policy_createCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	vpn_ike_policy_createCmd.Flags().String("project", "", "Owner's project (name or ID)")
	vpn_ike_policy_createCmd.Flags().String("project-domain", "", "Domain the project belongs to (name or ID).")
	vpn_ike_policy_createCmd.Flags().String("variable", "", "==SUPPRESS==")
	vpn_ike_policyCmd.AddCommand(vpn_ike_policy_createCmd)
}
