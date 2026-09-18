package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var router_add_subnetCmd = &cobra.Command{
	Use:   "subnet",
	Short: "Add a subnet to a router",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(router_add_subnetCmd).Standalone()

	router_add_subnetCmd.Flags().Bool("advertise-host", false, "Mark the subnet's prefixes to be advertised as host routes within the router's EVPN VNI.")
	router_addCmd.AddCommand(router_add_subnetCmd)
}
