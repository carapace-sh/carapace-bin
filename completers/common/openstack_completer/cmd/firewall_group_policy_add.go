package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var firewall_group_policy_addCmd = &cobra.Command{
	Use:   "add",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(firewall_group_policy_addCmd).Standalone()

	firewall_group_policyCmd.AddCommand(firewall_group_policy_addCmd)
}
