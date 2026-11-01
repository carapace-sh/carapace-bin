package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var network_trunk_subportCmd = &cobra.Command{
	Use:   "subport",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(network_trunk_subportCmd).Standalone()

	network_trunkCmd.AddCommand(network_trunk_subportCmd)
}
