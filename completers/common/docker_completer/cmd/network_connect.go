package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var network_connectCmd = &cobra.Command{
	Use:   "connect [OPTIONS] NETWORK CONTAINER",
	Short: "Connect a container to a network",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(network_connectCmd).Standalone()

	network_connectCmd.Flags().StringSlice("alias", nil, "Add network-scoped alias for the container")
	network_connectCmd.Flags().StringSlice("driver-opt", nil, "driver options for the network")
	network_connectCmd.Flags().String("ip", "", "IPv4 address (e.g., \"172.30.100.104\")")
	network_connectCmd.Flags().String("ip6", "", "IPv6 address (e.g., \"2001:db8::33\")")
	network_connectCmd.Flags().String("link", "", "Add link to another container")
	network_connectCmd.Flags().StringSlice("link-local-ip", nil, "Add a link-local address for the container")
	networkCmd.AddCommand(network_connectCmd)

	carapace.Gen(network_connectCmd).FlagCompletion(carapace.ActionMap{
		"link": docker.ActionContainers(),
	})

	carapace.Gen(network_connectCmd).PositionalCompletion(
		docker.ActionNetworks(),
		docker.ActionContainers(),
	)
}
