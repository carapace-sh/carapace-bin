package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cluster_agent_tokenCacheCmd = &cobra.Command{
	Use:   "token-cache <command> [flags]",
	Short: "Manage cached GitLab Agent tokens.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cluster_agent_tokenCacheCmd).Standalone()

	cluster_agentCmd.AddCommand(cluster_agent_tokenCacheCmd)
}
