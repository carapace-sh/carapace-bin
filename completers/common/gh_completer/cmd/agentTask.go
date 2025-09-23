package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var agentTaskCmd = &cobra.Command{
	Use:     "agent-task <command>",
	Short:   "Work with agent tasks (preview)",
	Aliases: []string{"agent-tasks", "agent", "agents"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(agentTaskCmd).Standalone()

	rootCmd.AddCommand(agentTaskCmd)
}
