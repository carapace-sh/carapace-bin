package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var router_unsetCmd = &cobra.Command{
	Use:   "unset",
	Short: "Unset router properties",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(router_unsetCmd).Standalone()

	router_unsetCmd.Flags().Bool("all-tag", false, "Clear all tags associated with the router")
	router_unsetCmd.Flags().Bool("external-gateway", false, "Remove external gateway information from the router")
	router_unsetCmd.Flags().String("extra-property", "", "Additional parameters can be passed using this property.")
	router_unsetCmd.Flags().Bool("qos-policy", false, "Remove QoS policy from router gateway IPs")
	router_unsetCmd.Flags().String("route", "", "Routes to be removed from the router.")
	router_unsetCmd.Flags().String("tag", "", "Tag to be removed from the router (repeat option to remove multiple tags)")
	routerCmd.AddCommand(router_unsetCmd)
}
