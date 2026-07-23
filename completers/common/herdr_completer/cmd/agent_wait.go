package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/herdr"
	"github.com/spf13/cobra"
)

var agent_waitCmd = &cobra.Command{
	Use:   "wait <target>",
	Short: "Wait for an agent status",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(agent_waitCmd).Standalone()

	agent_waitCmd.Flags().String("status", "", "")
	agent_waitCmd.Flags().String("timeout", "", "")
	agent_waitCmd.MarkFlagRequired("status")
	agentCmd.AddCommand(agent_waitCmd)

	carapace.Gen(agent_waitCmd).PositionalCompletion(herdr.ActionAgents())

	carapace.Gen(agent_waitCmd).FlagCompletion(carapace.ActionMap{
		"status": carapace.ActionValues("idle", "working", "blocked", "unknown"),
	})
}
