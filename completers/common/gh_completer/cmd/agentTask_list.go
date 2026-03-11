package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gh"
	"github.com/spf13/cobra"
)

var agentTask_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List agent tasks (preview)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(agentTask_listCmd).Standalone()

	agentTask_listCmd.Flags().StringP("jq", "q", "", "Filter JSON output using a jq `expression`")
	agentTask_listCmd.Flags().StringSlice("json", nil, "Output JSON with the specified `fields`")
	agentTask_listCmd.Flags().StringP("limit", "L", "", "Maximum number of agent tasks to fetch")
	agentTask_listCmd.Flags().StringP("template", "t", "", "Format JSON output using a Go template; see \"gh help formatting\"")
	agentTask_listCmd.Flags().BoolP("web", "w", false, "Open agent tasks in the browser")
	agentTaskCmd.AddCommand(agentTask_listCmd)

	carapace.Gen(agentTask_listCmd).FlagCompletion(carapace.ActionMap{
		"json": gh.ActionSessionFields().UniqueList(","),
	})
}
