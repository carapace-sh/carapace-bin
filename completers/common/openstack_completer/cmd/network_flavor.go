package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var network_flavorCmd = &cobra.Command{
	Use:   "flavor",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(network_flavorCmd).Standalone()

	networkCmd.AddCommand(network_flavorCmd)
}
