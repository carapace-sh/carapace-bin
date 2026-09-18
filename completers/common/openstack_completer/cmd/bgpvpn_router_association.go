package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bgpvpn_router_associationCmd = &cobra.Command{
	Use:   "association",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bgpvpn_router_associationCmd).Standalone()

	bgpvpn_routerCmd.AddCommand(bgpvpn_router_associationCmd)
}
