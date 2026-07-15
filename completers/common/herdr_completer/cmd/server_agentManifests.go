package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var server_agentManifestsCmd = &cobra.Command{
	Use:   "agent-manifests",
	Short: "Show active agent detection manifests",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(server_agentManifestsCmd).Standalone()

	server_agentManifestsCmd.Flags().Bool("json", false, "")
	serverCmd.AddCommand(server_agentManifestsCmd)
}
