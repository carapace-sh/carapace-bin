package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/herdr_completer/cmd/action"
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

	carapace.Gen(agent_focusCmd).PositionalCompletion(action.ActionAgents())
}
