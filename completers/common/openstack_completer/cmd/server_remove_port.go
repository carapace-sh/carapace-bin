package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/openstack_completer/cmd/action"
	"github.com/spf13/cobra"
)

var server_remove_portCmd = &cobra.Command{
	Use:   "port",
	Short: "Remove port from server",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(server_remove_portCmd).Standalone()

	server_removeCmd.AddCommand(server_remove_portCmd)

	carapace.Gen(server_remove_portCmd).PositionalCompletion(
		action.ActionServers(),
		action.ActionPorts(),
	)
}
