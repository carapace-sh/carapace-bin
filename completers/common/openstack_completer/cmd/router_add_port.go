package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var router_add_portCmd = &cobra.Command{
	Use:   "port",
	Short: "Add a port to a router",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(router_add_portCmd).Standalone()

	router_addCmd.AddCommand(router_add_portCmd)
}
