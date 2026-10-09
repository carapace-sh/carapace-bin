package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var vpn_ike_policy_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set IKE policy properties",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(vpn_ike_policy_setCmd).Standalone()

	vpn_ike_policy_setCmd.Flags().String("auth-algorithm", "", "Authentication algorithm")
	vpn_ike_policy_setCmd.Flags().String("description", "", "Description of the IKE policy")
	vpn_ike_policy_setCmd.Flags().String("encryption-algorithm", "", "Encryption algorithm")
	vpn_ike_policy_setCmd.Flags().String("ike-version", "", "IKE version for the policy")
	vpn_ike_policy_setCmd.Flags().String("lifetime", "", "IKE lifetime attributes.")
	vpn_ike_policy_setCmd.Flags().String("name", "", "Name of the IKE policy")
	vpn_ike_policy_setCmd.Flags().String("pfs", "", "Perfect Forward Secrecy")
	vpn_ike_policy_setCmd.Flags().String("phase1-negotiation-mode", "", "IKE Phase1 negotiation mode")
	vpn_ike_policyCmd.AddCommand(vpn_ike_policy_setCmd)
}
