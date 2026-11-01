package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var network_flavor_profileCmd = &cobra.Command{
	Use:   "profile",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(network_flavor_profileCmd).Standalone()

	network_flavorCmd.AddCommand(network_flavor_profileCmd)
}
