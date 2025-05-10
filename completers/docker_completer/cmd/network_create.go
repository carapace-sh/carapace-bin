package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var network_createCmd = &cobra.Command{
	Use:   "create [OPTIONS] NETWORK",
	Short: "Create a network",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(network_createCmd).Standalone()

	network_createCmd.Flags().Bool("attachable", false, "Enable manual container attachment")
	network_createCmd.Flags().String("aux-address", "", "Auxiliary IPv4 or IPv6 addresses used by Network driver")
	network_createCmd.Flags().String("config-from", "", "The network from which to copy the configuration")
	network_createCmd.Flags().Bool("config-only", false, "Create a configuration only network")
	network_createCmd.Flags().StringP("driver", "d", "bridge", "Driver to manage the Network")
	network_createCmd.Flags().StringSlice("gateway", nil, "IPv4 or IPv6 Gateway for the master subnet")
	network_createCmd.Flags().Bool("ingress", false, "Create swarm routing-mesh network")
	network_createCmd.Flags().Bool("internal", false, "Restrict external access to the network")
	network_createCmd.Flags().StringSlice("ip-range", nil, "Allocate container ip from a sub-range")
	network_createCmd.Flags().String("ipam-driver", "default", "IP Address Management Driver")
	network_createCmd.Flags().String("ipam-opt", "", "Set IPAM driver specific options")
	network_createCmd.Flags().Bool("ipv6", false, "Enable IPv6 networking")
	network_createCmd.Flags().String("label", "", "Set metadata on a network")
	network_createCmd.Flags().StringP("opt", "o", "", "Set driver specific options")
	network_createCmd.Flags().String("scope", "", "Control the network's scope")
	network_createCmd.Flags().StringSlice("subnet", nil, "Subnet in CIDR format that represents a network segment")
	networkCmd.AddCommand(network_createCmd)

	carapace.Gen(network_createCmd).FlagCompletion(carapace.ActionMap{
		"config-from": docker.ActionNetworks(),
		"driver":      carapace.ActionValues("bridge", "host", "null", "overlay"),
	})

	carapace.Gen(network_createCmd).PositionalCompletion(
		docker.ActionNetworks(),
	)
}
