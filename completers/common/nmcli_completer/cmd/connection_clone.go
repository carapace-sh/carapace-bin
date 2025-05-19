package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net"
	"github.com/spf13/cobra"
)

var connection_cloneCmd = &cobra.Command{
	Use:   "clone",
	Short: "Clone an existing connection profile",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connection_cloneCmd).Standalone()

	connection_cloneCmd.Flags().Bool("temporary", false, "only temporary")
	connectionCmd.AddCommand(connection_cloneCmd)

	carapace.Gen(connection_cloneCmd).PositionalCompletion(
		net.ActionConnections(),
	)
}
