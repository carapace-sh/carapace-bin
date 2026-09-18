package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_network_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set share network properties",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_network_setCmd).Standalone()

	share_network_setCmd.Flags().Bool("check-only", false, "Run a dry-run of a security service replacement.")
	share_network_setCmd.Flags().String("current-security-service", "", "Name or ID of a security service that is currently associated with a share network that must be replaced.")
	share_network_setCmd.Flags().String("description", "", "Set a new description to the share network.")
	share_network_setCmd.Flags().String("name", "", "Set a new name to the share network.")
	share_network_setCmd.Flags().String("neutron-net-id", "", "Update the neutron network associated with the default share network subnet.")
	share_network_setCmd.Flags().String("neutron-subnet-id", "", "Update the neutron subnetwork associated with the default share network subnet.")
	share_network_setCmd.Flags().String("new-security-service", "", "Name or ID of a security service that must be associated with the share network.")
	share_network_setCmd.Flags().Bool("restart-check", false, "Restart a dry-run of a security service replacement.")
	share_network_setCmd.Flags().String("status", "", "Assign a status to the share network (Admin only).")
	share_networkCmd.AddCommand(share_network_setCmd)
}
