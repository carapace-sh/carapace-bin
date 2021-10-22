package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var compute_loadBalancer_addForwardingRulesCmd = &cobra.Command{
	Use:   "add-forwarding-rules",
	Short: "Add forwarding rules to a load balancer",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compute_loadBalancer_addForwardingRulesCmd).Standalone()
	compute_loadBalancer_addForwardingRulesCmd.Flags().String("forwarding-rules", "", "A comma-separated list of key-value pairs representing forwarding rules, which define how traffic is routed, e.g.: `entry_protocol:tcp,entry_port:3306,target_protocol:tcp,target_port:3306`.")
	compute_loadBalancerCmd.AddCommand(compute_loadBalancer_addForwardingRulesCmd)
}
