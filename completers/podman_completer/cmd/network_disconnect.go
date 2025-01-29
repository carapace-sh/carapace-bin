package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var network_disconnectCmd = &cobra.Command{
	Use:   "disconnect [options] NETWORK CONTAINER",
	Short: "Disconnect a container from a network",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(network_disconnectCmd).Standalone()

	network_disconnectCmd.Flags().BoolP("force", "f", false, "force removal of container from network")
	networkCmd.AddCommand(network_disconnectCmd)
}
