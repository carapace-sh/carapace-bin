package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bgpvpn_routerCmd = &cobra.Command{
	Use:   "router",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bgpvpn_routerCmd).Standalone()

	bgpvpnCmd.AddCommand(bgpvpn_routerCmd)
}
