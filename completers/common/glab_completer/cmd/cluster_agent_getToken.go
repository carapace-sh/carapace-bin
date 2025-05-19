package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cluster_agent_getTokenCmd = &cobra.Command{
	Use:   "get-token [flags]",
	Short: "Create and return a k8s_proxy-scoped personal access token to authenticate with a GitLab Agents for Kubernetes.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cluster_agent_getTokenCmd).Standalone()

	cluster_agent_getTokenCmd.Flags().StringP("agent", "a", "", "The numerical Agent ID to connect to.")
	cluster_agent_getTokenCmd.MarkFlagRequired("agent")
	cluster_agentCmd.AddCommand(cluster_agent_getTokenCmd)

	// TODO flag completion
}
