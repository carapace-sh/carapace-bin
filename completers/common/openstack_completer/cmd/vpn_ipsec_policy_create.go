package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var vpn_ipsec_policy_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create an IPsec policy",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(vpn_ipsec_policy_createCmd).Standalone()

	vpn_ipsec_policy_createCmd.Flags().String("auth-algorithm", "", "Authentication algorithm for IPsec policy")
	vpn_ipsec_policy_createCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	vpn_ipsec_policy_createCmd.Flags().String("description", "", "Description of the IPsec policy")
	vpn_ipsec_policy_createCmd.Flags().String("encapsulation-mode", "", "Encapsulation mode for IPsec policy")
	vpn_ipsec_policy_createCmd.Flags().String("encryption-algorithm", "", "Encryption algorithm for IPsec policy")
	vpn_ipsec_policy_createCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	vpn_ipsec_policy_createCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	vpn_ipsec_policy_createCmd.Flags().String("lifetime", "", "IPsec lifetime attributes.")
	vpn_ipsec_policy_createCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	vpn_ipsec_policy_createCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	vpn_ipsec_policy_createCmd.Flags().String("pfs", "", "Perfect Forward Secrecy for IPsec policy")
	vpn_ipsec_policy_createCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	vpn_ipsec_policy_createCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	vpn_ipsec_policy_createCmd.Flags().String("project", "", "Owner's project (name or ID)")
	vpn_ipsec_policy_createCmd.Flags().String("project-domain", "", "Domain the project belongs to (name or ID).")
	vpn_ipsec_policy_createCmd.Flags().String("transform-protocol", "", "Transform protocol for IPsec policy")
	vpn_ipsec_policy_createCmd.Flags().String("variable", "", "==SUPPRESS==")
	vpn_ipsec_policyCmd.AddCommand(vpn_ipsec_policy_createCmd)
}
