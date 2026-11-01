package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var network_flavor_profile_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete network flavor profile",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(network_flavor_profile_deleteCmd).Standalone()

	network_flavor_profileCmd.AddCommand(network_flavor_profile_deleteCmd)
}
