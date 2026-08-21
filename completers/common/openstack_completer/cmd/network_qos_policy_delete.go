package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var network_qos_policy_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete Qos Policy(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(network_qos_policy_deleteCmd).Standalone()

	network_qos_policyCmd.AddCommand(network_qos_policy_deleteCmd)
}
