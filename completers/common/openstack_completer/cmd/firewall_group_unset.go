package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var firewall_group_unsetCmd = &cobra.Command{
	Use:   "unset",
	Short: "Unset firewall group properties",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(firewall_group_unsetCmd).Standalone()

	firewall_group_unsetCmd.Flags().Bool("all-port", false, "Remove all ports for this firewall group")
	firewall_group_unsetCmd.Flags().Bool("egress-firewall-policy", false, "Egress firewall policy (name or ID) to delete")
	firewall_group_unsetCmd.Flags().Bool("enable", false, "(Deprecated) Use \"firewall group set --disable\" instead.")
	firewall_group_unsetCmd.Flags().Bool("ingress-firewall-policy", false, "Ingress firewall policy (name or ID) to delete")
	firewall_group_unsetCmd.Flags().String("port", "", "Port(s) (name or ID) to apply firewall group.")
	firewall_group_unsetCmd.Flags().Bool("share", false, "(Deprecated) Use \"firewall group set --no-share\" instead.")
	firewall_groupCmd.AddCommand(firewall_group_unsetCmd)
}
