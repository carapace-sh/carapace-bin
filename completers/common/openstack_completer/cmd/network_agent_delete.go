package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/openstack_completer/cmd/action"
	"github.com/spf13/cobra"
)

var network_agent_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete network agent(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(network_agent_deleteCmd).Standalone()

	network_agentCmd.AddCommand(network_agent_deleteCmd)

	carapace.Gen(network_agent_deleteCmd).PositionalAnyCompletion(action.ActionNetworkAgents())
}
