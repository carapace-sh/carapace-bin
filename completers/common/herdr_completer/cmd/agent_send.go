package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/herdr_completer/cmd/action"
	"github.com/spf13/cobra"
)

var agent_sendCmd = &cobra.Command{
	Use:   "send <target> <text>",
	Short: "Send text to an agent",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(agent_sendCmd).Standalone()

	agentCmd.AddCommand(agent_sendCmd)

	carapace.Gen(agent_sendCmd).PositionalCompletion(action.ActionAgents())
}
