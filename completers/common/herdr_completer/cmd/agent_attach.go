package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/herdr"
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

	carapace.Gen(agent_attachCmd).PositionalCompletion(herdr.ActionAgents())
}
