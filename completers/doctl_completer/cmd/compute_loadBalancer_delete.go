package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var compute_loadBalancer_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Permanently delete a load balancer",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compute_loadBalancer_deleteCmd).Standalone()
	compute_loadBalancer_deleteCmd.Flags().BoolP("force", "f", false, "Delete the load balancer without a confirmation prompt")
	compute_loadBalancerCmd.AddCommand(compute_loadBalancer_deleteCmd)
}
