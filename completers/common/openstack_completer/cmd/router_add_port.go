package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/openstack_completer/cmd/action"
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

	carapace.Gen(router_add_portCmd).PositionalCompletion(
		action.ActionRouters(router_add_portCmd),
		action.ActionPorts(router_add_portCmd),
	)
}
