package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var compute_loadBalancerCmd = &cobra.Command{
	Use:   "load-balancer",
	Short: "Display commands to manage load balancers",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compute_loadBalancerCmd).Standalone()
	computeCmd.AddCommand(compute_loadBalancerCmd)
}
