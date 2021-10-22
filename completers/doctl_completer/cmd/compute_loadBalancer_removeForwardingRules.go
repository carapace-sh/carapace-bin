package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var compute_loadBalancer_removeForwardingRulesCmd = &cobra.Command{
	Use:   "remove-forwarding-rules",
	Short: "Remove forwarding rules from a load balancer",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compute_loadBalancer_removeForwardingRulesCmd).Standalone()
	compute_loadBalancer_removeForwardingRulesCmd.Flags().String("forwarding-rules", "", "A comma-separated list of key-value pairs representing forwarding rules, which define how traffic is routed, e.g.: `entry_protocol:tcp,entry_port:3306,target_protocol:tcp,target_port:3306`.")
	compute_loadBalancerCmd.AddCommand(compute_loadBalancer_removeForwardingRulesCmd)
}
