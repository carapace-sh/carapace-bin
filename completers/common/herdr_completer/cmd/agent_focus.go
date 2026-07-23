package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/herdr"
	"github.com/spf13/cobra"
)

var agent_focusCmd = &cobra.Command{
	Use:   "focus <target>",
	Short: "Focus an agent",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(agent_focusCmd).Standalone()

	agentCmd.AddCommand(agent_focusCmd)

	carapace.Gen(agent_focusCmd).PositionalCompletion(herdr.ActionAgents())
}
