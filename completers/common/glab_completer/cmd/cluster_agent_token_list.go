package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cluster_agent_token_listCmd = &cobra.Command{
	Use:   "list <agent-id> [flags]",
	Short: "List tokens of an agent.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cluster_agent_token_listCmd).Standalone()

	cluster_agent_tokenCmd.AddCommand(cluster_agent_token_listCmd)
}
