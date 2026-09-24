package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var vpn_ipsec_policy_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete IPsec policy(policies)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(vpn_ipsec_policy_deleteCmd).Standalone()

	vpn_ipsec_policyCmd.AddCommand(vpn_ipsec_policy_deleteCmd)
}
