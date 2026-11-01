package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/herdr"
	"github.com/spf13/cobra"
)

var agent_sendKeysCmd = &cobra.Command{
	Use:   "send-keys <target> <key> [key]",
	Short: "Send key presses to an agent",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(agent_sendKeysCmd).Standalone()

	agentCmd.AddCommand(agent_sendKeysCmd)

	carapace.Gen(agent_sendKeysCmd).PositionalCompletion(herdr.ActionAgents())
}
