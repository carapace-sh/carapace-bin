package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/herdr"
	"github.com/spf13/cobra"
)

var agent_getCmd = &cobra.Command{
	Use:   "get <target>",
	Short: "Show an agent",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(agent_getCmd).Standalone()

	agentCmd.AddCommand(agent_getCmd)

	carapace.Gen(agent_getCmd).PositionalCompletion(herdr.ActionAgents())
}
