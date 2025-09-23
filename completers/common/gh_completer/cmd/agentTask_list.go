package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var agentTask_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List agent tasks (preview)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(agentTask_listCmd).Standalone()

	agentTask_listCmd.Flags().StringP("limit", "L", "", "Maximum number of agent tasks to fetch (default 30)")
	agentTask_listCmd.Flags().BoolP("web", "w", false, "Open agent tasks in the browser")
	agentTaskCmd.AddCommand(agentTask_listCmd)
}
