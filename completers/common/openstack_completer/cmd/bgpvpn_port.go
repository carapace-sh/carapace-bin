package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bgpvpn_portCmd = &cobra.Command{
	Use:   "port",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bgpvpn_portCmd).Standalone()

	bgpvpnCmd.AddCommand(bgpvpn_portCmd)
}
