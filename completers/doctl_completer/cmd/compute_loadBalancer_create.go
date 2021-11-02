package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/doctl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var compute_loadBalancer_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new load balancer",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compute_loadBalancer_createCmd).Standalone()
	compute_loadBalancer_createCmd.Flags().String("algorithm", "round_robin", "The algorithm to use when traffic is distributed across your Droplets; possible values: `round_robin` or `least_connections`")
	compute_loadBalancer_createCmd.Flags().Bool("disable-lets-encrypt-dns-records", false, "disable automatic DNS record creation for Let's Encrypt certificates that are added to the load balancer")
	compute_loadBalancer_createCmd.Flags().StringSlice("droplet-ids", []string{}, "A comma-separated list of Droplet IDs to add to the load balancer, e.g.: `12,33`")
	compute_loadBalancer_createCmd.Flags().Bool("enable-backend-keepalive", false, "enable keepalive connections to backend target droplets")
	compute_loadBalancer_createCmd.Flags().Bool("enable-proxy-protocol", false, "enable proxy protocol")
	compute_loadBalancer_createCmd.Flags().String("forwarding-rules", "", "A comma-separated list of key-value pairs representing forwarding rules, which define how traffic is routed, e.g.: `entry_protocol:tcp,entry_port:3306,target_protocol:tcp,target_port:3306`.")
	compute_loadBalancer_createCmd.Flags().String("health-check", "", "A comma-separated list of key-value pairs representing recent health check results, e.g.: `protocol:http,port:80,path:/index.html,check_interval_seconds:10,response_timeout_seconds:5,healthy_threshold:5,unhealthy_threshold:3`")
	compute_loadBalancer_createCmd.Flags().String("name", "", "The load balancer's name (required)")
	compute_loadBalancer_createCmd.Flags().Bool("redirect-http-to-https", false, "Redirects HTTP requests to the load balancer on port 80 to HTTPS on port 443")
	compute_loadBalancer_createCmd.Flags().String("region", "", "The load balancer's region, e.g.: `nyc1` (required)")
	compute_loadBalancer_createCmd.Flags().String("size", "", "The load balancer's size, e.g.: `lb-small`. Only one of size and size-unit should be used")
	compute_loadBalancer_createCmd.Flags().Int("size-unit", 0, "The load balancer's size, e.g.: 1. Only one of size-unit and size should be used")
	compute_loadBalancer_createCmd.Flags().String("sticky-sessions", "", "A comma-separated list of key-value pairs representing a list of active sessions, e.g.: `type:cookies, cookie_name:DO-LB, cookie_ttl_seconds:5`")
	compute_loadBalancer_createCmd.Flags().String("tag-name", "", "droplet tag name")
	compute_loadBalancer_createCmd.Flags().String("vpc-uuid", "", "The UUID of the VPC to create the load balancer in")
	compute_loadBalancerCmd.AddCommand(compute_loadBalancer_createCmd)

	// TODO flag completion
	carapace.Gen(compute_loadBalancer_createCmd).FlagCompletion(carapace.ActionMap{
		"algorithm": carapace.ActionValues("algorithm", "round_robin"),
		"region":    action.ActionRegions(),
	})
}
