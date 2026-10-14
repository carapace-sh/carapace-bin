package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var vpn_ipsec_policy_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set IPsec policy properties",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(vpn_ipsec_policy_setCmd).Standalone()

	vpn_ipsec_policy_setCmd.Flags().String("auth-algorithm", "", "Authentication algorithm for IPsec policy")
	vpn_ipsec_policy_setCmd.Flags().String("description", "", "Description of the IPsec policy")
	vpn_ipsec_policy_setCmd.Flags().String("encapsulation-mode", "", "Encapsulation mode for IPsec policy")
	vpn_ipsec_policy_setCmd.Flags().String("encryption-algorithm", "", "Encryption algorithm for IPsec policy")
	vpn_ipsec_policy_setCmd.Flags().String("lifetime", "", "IPsec lifetime attributes.")
	vpn_ipsec_policy_setCmd.Flags().String("name", "", "Name of the IPsec policy")
	vpn_ipsec_policy_setCmd.Flags().String("pfs", "", "Perfect Forward Secrecy for IPsec policy")
	vpn_ipsec_policy_setCmd.Flags().String("transform-protocol", "", "Transform protocol for IPsec policy")
	vpn_ipsec_policyCmd.AddCommand(vpn_ipsec_policy_setCmd)
}
