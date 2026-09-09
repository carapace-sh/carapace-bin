package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var vpn_ike_policy_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete IKE policy (policies)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(vpn_ike_policy_deleteCmd).Standalone()

	vpn_ike_policyCmd.AddCommand(vpn_ike_policy_deleteCmd)
}
