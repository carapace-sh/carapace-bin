package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var network_existsCmd = &cobra.Command{
	Use:   "exists NETWORK",
	Short: "Check if network exists",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(network_existsCmd).Standalone()

	networkCmd.AddCommand(network_existsCmd)
}
