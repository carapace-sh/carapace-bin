package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debug_peerRelayCmd = &cobra.Command{
	Use:   "peer-relay-sessions",
	Short: "Print the current set of active peer relay sessions relayed through this node",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debug_peerRelayCmd).Standalone()

	debugCmd.AddCommand(debug_peerRelayCmd)
}
