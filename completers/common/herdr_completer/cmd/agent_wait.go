package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/herdr"
	"github.com/spf13/cobra"
)

var agent_waitCmd = &cobra.Command{
	Use:   "wait <target>",
	Short: "Wait until an agent reaches one of the requested states",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(agent_waitCmd).Standalone()

	agent_waitCmd.Flags().String("timeout", "", "Fail after this many milliseconds")
	agent_waitCmd.Flags().StringSlice("until", nil, "State to match; repeat for more than one state")
	agentCmd.AddCommand(agent_waitCmd)

	carapace.Gen(agent_waitCmd).PositionalCompletion(herdr.ActionAgents())

	carapace.Gen(agent_waitCmd).FlagCompletion(carapace.ActionMap{
		"until": carapace.ActionValues("idle", "working", "blocked", "done", "unknown"),
	})
}
