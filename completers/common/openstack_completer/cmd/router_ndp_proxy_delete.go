package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var router_ndp_proxy_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete NDP proxy",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(router_ndp_proxy_deleteCmd).Standalone()

	router_ndp_proxyCmd.AddCommand(router_ndp_proxy_deleteCmd)
}
