package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var network_trunk_unsetCmd = &cobra.Command{
	Use:   "unset",
	Short: "Unset subports from a given network trunk",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(network_trunk_unsetCmd).Standalone()

	network_trunk_unsetCmd.Flags().String("subport", "", "Subport to unset (name or ID of the port) (repeat option to unset multiple subports)")
	network_trunk_unsetCmd.MarkFlagRequired("subport")
	network_trunkCmd.AddCommand(network_trunk_unsetCmd)
}
