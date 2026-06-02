package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net"
	"github.com/spf13/cobra"
)

var connection_upCmd = &cobra.Command{
	Use:   "up",
	Short: "Activate a connection",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connection_upCmd).Standalone()

	connectionCmd.AddCommand(connection_upCmd)

	carapace.Gen(connection_upCmd).PositionalCompletion(
		net.ActionConnections(),
	)
}
