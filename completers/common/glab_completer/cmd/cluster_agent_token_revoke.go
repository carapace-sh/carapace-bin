package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cluster_agent_token_revokeCmd = &cobra.Command{
	Use:   "revoke <agent-id> <token-id>",
	Short: "Revoke a token of an agent.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cluster_agent_token_revokeCmd).Standalone()

	cluster_agent_tokenCmd.AddCommand(cluster_agent_token_revokeCmd)
}
