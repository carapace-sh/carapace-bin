package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debug_peer_endpoint_changesCmd = &cobra.Command{
	Use:   "peer-endpoint-changes",
	Short: "Print debug information about a peer's endpoint changes",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debug_peer_endpoint_changesCmd).Standalone()

	debugCmd.AddCommand(debug_peer_endpoint_changesCmd)
}
