package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var network_flavor_remove_profileCmd = &cobra.Command{
	Use:   "profile",
	Short: "Remove service profile from network flavor",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(network_flavor_remove_profileCmd).Standalone()

	network_flavor_removeCmd.AddCommand(network_flavor_remove_profileCmd)
}
