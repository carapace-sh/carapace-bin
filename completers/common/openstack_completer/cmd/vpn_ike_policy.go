package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var vpn_ike_policyCmd = &cobra.Command{
	Use:   "policy",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(vpn_ike_policyCmd).Standalone()

	vpn_ikeCmd.AddCommand(vpn_ike_policyCmd)
}
