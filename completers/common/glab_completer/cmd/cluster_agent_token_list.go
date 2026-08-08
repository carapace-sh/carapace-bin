package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cluster_agent_token_listCmd = &cobra.Command{
	Use:   "list <agent-id> [flags]",
	Short: "List tokens for a GitLab Agent for Kubernetes.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cluster_agent_token_listCmd).Standalone()

	cluster_agent_token_listCmd.Flags().String("jq", "", "Filter JSON output with a jq expression.")
	cluster_agent_token_listCmd.Flags().StringP("output", "F", "", "Format output as: text, json.")
	cluster_agent_tokenCmd.AddCommand(cluster_agent_token_listCmd)

	carapace.Gen(cluster_agent_token_listCmd).FlagCompletion(carapace.ActionMap{
		"output": carapace.ActionValues("text", "json"),
	})
}
