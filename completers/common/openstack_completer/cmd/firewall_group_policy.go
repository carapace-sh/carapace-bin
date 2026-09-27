package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var firewall_group_policyCmd = &cobra.Command{
	Use:   "policy",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(firewall_group_policyCmd).Standalone()

	firewall_groupCmd.AddCommand(firewall_group_policyCmd)
}
