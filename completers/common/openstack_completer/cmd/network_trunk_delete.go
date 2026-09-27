package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var network_trunk_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a given network trunk",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(network_trunk_deleteCmd).Standalone()

	network_trunkCmd.AddCommand(network_trunk_deleteCmd)
}
