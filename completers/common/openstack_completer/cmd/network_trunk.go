package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var network_trunkCmd = &cobra.Command{
	Use:   "trunk",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(network_trunkCmd).Standalone()

	networkCmd.AddCommand(network_trunkCmd)
}
