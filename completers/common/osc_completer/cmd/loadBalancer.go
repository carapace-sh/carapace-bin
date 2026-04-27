package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var loadBalancerCmd = &cobra.Command{
	Use:   "load-balancer",
	Short: "Load Balancer service commands",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(loadBalancerCmd).Standalone()
	rootCmd.AddCommand(loadBalancerCmd)

	subresources := []struct {
		use   string
		short string
	}{
		{"loadbalancer", "Load Balancer commands"},
		{"listener", "Listener commands"},
		{"pool", "Pool commands"},
		{"member", "Member commands"},
		{"healthmonitor", "Health Monitor commands"},
	}

	for _, r := range subresources {
		sub := &cobra.Command{Use: r.use, Short: r.short, Run: func(cmd *cobra.Command, args []string) {}}
		carapace.Gen(sub).Standalone()
		loadBalancerCmd.AddCommand(sub)
		addCrudSubcommands(sub)
	}
}
