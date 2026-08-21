package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bgpvpn_port_associationCmd = &cobra.Command{
	Use:   "association",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bgpvpn_port_associationCmd).Standalone()

	bgpvpn_portCmd.AddCommand(bgpvpn_port_associationCmd)
}
