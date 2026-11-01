package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var router_ndp_proxy_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set NDP proxy properties",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(router_ndp_proxy_setCmd).Standalone()

	router_ndp_proxy_setCmd.Flags().String("description", "", "Text to describe/contextualize the use of the NDP proxy configuration")
	router_ndp_proxy_setCmd.Flags().String("name", "", "Set NDP proxy name")
	router_ndp_proxyCmd.AddCommand(router_ndp_proxy_setCmd)
}
