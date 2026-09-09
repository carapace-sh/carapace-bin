package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var network_flavor_removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(network_flavor_removeCmd).Standalone()

	network_flavorCmd.AddCommand(network_flavor_removeCmd)
}
