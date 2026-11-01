package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bgpvpn_network_associationCmd = &cobra.Command{
	Use:   "association",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bgpvpn_network_associationCmd).Standalone()

	bgpvpn_networkCmd.AddCommand(bgpvpn_network_associationCmd)
}
