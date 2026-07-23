package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var server_reloadAgentManifestsCmd = &cobra.Command{
	Use:   "reload-agent-manifests",
	Short: "Reload local agent detection manifest overrides",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(server_reloadAgentManifestsCmd).Standalone()

	serverCmd.AddCommand(server_reloadAgentManifestsCmd)
}
