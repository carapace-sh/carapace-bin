package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_network_subnet_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set share network subnet properties",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_network_subnet_setCmd).Standalone()

	share_network_subnet_setCmd.Flags().String("property", "", "Set a property to this share network subnet (repeat option to set multiple properties).")
	share_network_subnetCmd.AddCommand(share_network_subnet_setCmd)
}
