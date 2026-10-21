package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var network_trunk_subport_removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove subport(s) from a given network trunk",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(network_trunk_subport_removeCmd).Standalone()

	network_trunk_subportCmd.AddCommand(network_trunk_subport_removeCmd)
}
