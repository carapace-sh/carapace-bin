package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cluster_agent_listCmd = &cobra.Command{
	Use:     "list [flags]",
	Short:   "List GitLab Agents for Kubernetes in a project.",
	Aliases: []string{"ls"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cluster_agent_listCmd).Standalone()

	cluster_agent_listCmd.Flags().String("jq", "", "Filter JSON output with a jq expression.")
	cluster_agent_listCmd.Flags().StringP("output", "F", "", "Format output as: text, json.")
	cluster_agent_listCmd.Flags().StringP("page", "p", "", "Page number.")
	cluster_agent_listCmd.Flags().StringP("per-page", "P", "", "Number of items to list per page.")
	cluster_agentCmd.AddCommand(cluster_agent_listCmd)

	carapace.Gen(cluster_agent_listCmd).FlagCompletion(carapace.ActionMap{
		"output": carapace.ActionValues("text", "json"),
	})
}
