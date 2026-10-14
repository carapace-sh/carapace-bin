package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var firewall_group_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set firewall group properties",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(firewall_group_setCmd).Standalone()

	firewall_group_setCmd.Flags().String("description", "", "Description of the firewall group")
	firewall_group_setCmd.Flags().Bool("disable", false, "Disable firewall group")
	firewall_group_setCmd.Flags().String("egress-firewall-policy", "", "Egress firewall policy (name or ID)")
	firewall_group_setCmd.Flags().Bool("enable", false, "Enable firewall group")
	firewall_group_setCmd.Flags().String("ingress-firewall-policy", "", "Ingress firewall policy (name or ID)")
	firewall_group_setCmd.Flags().String("name", "", "Name for the firewall group")
	firewall_group_setCmd.Flags().Bool("no-egress-firewall-policy", false, "Detach egress firewall policy from the firewall group")
	firewall_group_setCmd.Flags().Bool("no-ingress-firewall-policy", false, "Detach ingress firewall policy from the firewall group")
	firewall_group_setCmd.Flags().Bool("no-port", false, "Detach all port from the firewall group")
	firewall_group_setCmd.Flags().Bool("no-share", false, "Restrict use of the firewall group to the current project")
	firewall_group_setCmd.Flags().String("port", "", "Port(s) (name or ID) to apply firewall group.")
	firewall_group_setCmd.Flags().Bool("share", false, "Share the firewall group to be used in all projects (by default, it is restricted to be used by the current project).")
	firewall_groupCmd.AddCommand(firewall_group_setCmd)
}
