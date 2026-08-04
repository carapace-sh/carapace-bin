package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/herdr"
	"github.com/spf13/cobra"
)

var agent_promptCmd = &cobra.Command{
	Use:   "prompt <target>",
	Short: "Submit a prompt to an agent",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(agent_promptCmd).Standalone()

	agent_promptCmd.Flags().String("timeout", "", "Fail after this many milliseconds")
	agent_promptCmd.Flags().StringSlice("until", nil, "State to match after --wait; repeat for more than one state")
	agent_promptCmd.Flags().Bool("wait", false, "Wait for the first matching state observed after submission")
	agentCmd.AddCommand(agent_promptCmd)

	carapace.Gen(agent_promptCmd).PositionalCompletion(herdr.ActionAgents())

	carapace.Gen(agent_promptCmd).FlagCompletion(carapace.ActionMap{
		"until": carapace.ActionValues("idle", "working", "blocked", "done", "unknown"),
	})
}
