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

	carapace.Gen(vpn_ike_policy_setCmd).FlagCompletion(carapace.ActionMap{
		"auth-algorithm":          carapace.ActionValues("aes-cmac", "aes-xcbc", "sha1", "sha256", "sha384", "sha512"),
		"encryption-algorithm":    carapace.ActionValues("3des", "aes-128", "aes-128-ccm-12", "aes-128-ccm-16", "aes-128-ccm-8", "aes-128-ctr", "aes-128-gcm-12", "aes-128-gcm-16", "aes-128-gcm-8", "aes-192", "aes-192-ccm-12", "aes-192-ccm-16", "aes-192-ccm-8", "aes-192-ctr", "aes-192-gcm-12", "aes-192-gcm-16", "aes-192-gcm-8", "aes-256", "aes-256-ccm-12", "aes-256-ccm-16", "aes-256-ccm-8", "aes-256-ctr", "aes-256-gcm-12", "aes-256-gcm-16", "aes-256-gcm-8"),
		"ike-version":             carapace.ActionValues("v1", "v2"),
		"pfs":                     carapace.ActionValues("group14", "group15", "group16", "group17", "group18", "group19", "group2", "group20", "group21", "group22", "group23", "group24", "group25", "group26", "group27", "group28", "group29", "group30", "group31", "group5"),
		"phase1-negotiation-mode": carapace.ActionValues("aggressive", "main"),
	})
}
