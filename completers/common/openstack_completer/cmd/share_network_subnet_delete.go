package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_network_subnet_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a share network subnet",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_network_subnet_deleteCmd).Standalone()

	share_network_subnetCmd.AddCommand(share_network_subnet_deleteCmd)
}
