package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_network_subnet_unsetCmd = &cobra.Command{
	Use:   "unset",
	Short: "Unset a share network subnet property",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_network_subnet_unsetCmd).Standalone()

	share_network_subnet_unsetCmd.Flags().String("property", "", "Remove a property from share network subnet (repeat option to remove multiple properties).")
	share_network_subnetCmd.AddCommand(share_network_subnet_unsetCmd)
}
