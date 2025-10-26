package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cluster_agent_tokenCmd = &cobra.Command{
	Use:   "token <command> [flags]",
	Short: "Manage GitLab Agents for Kubernetes tokens.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cluster_agent_tokenCmd).Standalone()

	cluster_agentCmd.AddCommand(cluster_agent_tokenCmd)
}
