package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/herdr_completer/cmd/action"
	"github.com/spf13/cobra"
)

var agent_attachCmd = &cobra.Command{
	Use:   "attach <target>",
	Short: "Attach directly to an agent terminal",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(agent_attachCmd).Standalone()

	agent_attachCmd.Flags().Bool("takeover", false, "")
	agentCmd.AddCommand(agent_attachCmd)

	carapace.Gen(agent_attachCmd).PositionalCompletion(action.ActionAgents())
}
