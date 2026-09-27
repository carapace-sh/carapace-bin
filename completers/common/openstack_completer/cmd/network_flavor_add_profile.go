package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var network_flavor_add_profileCmd = &cobra.Command{
	Use:   "profile",
	Short: "Add a service profile to a network flavor",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(network_flavor_add_profileCmd).Standalone()

	network_flavor_addCmd.AddCommand(network_flavor_add_profileCmd)
}
