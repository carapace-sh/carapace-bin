package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var network_flavor_addCmd = &cobra.Command{
	Use:   "add",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(network_flavor_addCmd).Standalone()

	network_flavorCmd.AddCommand(network_flavor_addCmd)
}
