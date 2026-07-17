package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debug_peer_relay_serversCmd = &cobra.Command{
	Use:   "peer-relay-servers",
	Short: "Print the current set of candidate peer relay servers",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debug_peer_relay_serversCmd).Standalone()

	debugCmd.AddCommand(debug_peer_relay_serversCmd)
}
