package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var compute_loadBalancer_addDropletsCmd = &cobra.Command{
	Use:   "add-droplets",
	Short: "Add Droplets to a load balancer",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compute_loadBalancer_addDropletsCmd).Standalone()
	compute_loadBalancer_addDropletsCmd.Flags().StringSlice("droplet-ids", []string{}, "A comma-separated list of IDs of Droplet to add to the load balancer, example value: `12,33`")
	compute_loadBalancerCmd.AddCommand(compute_loadBalancer_addDropletsCmd)
}
