package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debug_peerRelayServersCmd = &cobra.Command{
	Use:   "peer-relay-servers",
	Short: "Print the current set of candidate peer relay servers",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debug_peerRelayServersCmd).Standalone()

	debugCmd.AddCommand(debug_peerRelayServersCmd)
}
