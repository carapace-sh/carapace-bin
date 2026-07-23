package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var server_updateAgentManifestsCmd = &cobra.Command{
	Use:   "update-agent-manifests",
	Short: "Fetch and reload agent detection manifests",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(server_updateAgentManifestsCmd).Standalone()

	server_updateAgentManifestsCmd.Flags().Bool("json", false, "")
	serverCmd.AddCommand(server_updateAgentManifestsCmd)
}
