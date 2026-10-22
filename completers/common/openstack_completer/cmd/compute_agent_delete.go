package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/openstack_completer/cmd/action"
	"github.com/spf13/cobra"
)

var compute_agent_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete compute agent(s).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compute_agent_deleteCmd).Standalone()

	compute_agentCmd.AddCommand(compute_agent_deleteCmd)

	carapace.Gen(compute_agent_deleteCmd).PositionalAnyCompletion(action.ActionComputeAgents())
}
