package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var router_remove_portCmd = &cobra.Command{
	Use:   "port",
	Short: "Remove a port from a router",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(router_remove_portCmd).Standalone()

	router_removeCmd.AddCommand(router_remove_portCmd)
}
