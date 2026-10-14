package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var router_remove_subnetCmd = &cobra.Command{
	Use:   "subnet",
	Short: "Remove a subnet from a router",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(router_remove_subnetCmd).Standalone()

	router_removeCmd.AddCommand(router_remove_subnetCmd)
}
