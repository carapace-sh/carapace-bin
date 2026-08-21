package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var network_flavor_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete network flavors",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(network_flavor_deleteCmd).Standalone()

	network_flavorCmd.AddCommand(network_flavor_deleteCmd)
}
