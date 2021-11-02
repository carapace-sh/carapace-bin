package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var compute_loadBalancer_removeDropletsCmd = &cobra.Command{
	Use:   "remove-droplets",
	Short: "Remove Droplets from a load balancer",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compute_loadBalancer_removeDropletsCmd).Standalone()
	compute_loadBalancer_removeDropletsCmd.Flags().StringSlice("droplet-ids", []string{}, "A comma-separated list of IDs of Droplets to remove from the load balancer, example value: `12,33`")
	compute_loadBalancerCmd.AddCommand(compute_loadBalancer_removeDropletsCmd)
}
