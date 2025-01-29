package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var network_connectCmd = &cobra.Command{
	Use:   "connect [options] NETWORK CONTAINER",
	Short: "Add container to a network",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(network_connectCmd).Standalone()

	network_connectCmd.Flags().StringSlice("alias", []string{}, "network scoped alias for container")
	network_connectCmd.Flags().String("ip", "", "set a static ipv4 address for this container network")
	network_connectCmd.Flags().String("ip6", "", "set a static ipv6 address for this container network")
	network_connectCmd.Flags().String("mac-address", "", "set a static mac address for this container network")
	networkCmd.AddCommand(network_connectCmd)
}
