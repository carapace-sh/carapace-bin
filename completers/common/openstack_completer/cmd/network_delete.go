package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var network_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete network(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(network_deleteCmd).Standalone()

	networkCmd.AddCommand(network_deleteCmd)
}
