package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var router_ndp_proxyCmd = &cobra.Command{
	Use:   "proxy",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(router_ndp_proxyCmd).Standalone()

	router_ndpCmd.AddCommand(router_ndp_proxyCmd)
}
